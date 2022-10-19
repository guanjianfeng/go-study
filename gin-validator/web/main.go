package main
// https://pkg.go.dev/github.com/go-playground/validator
// https://github.com/go-playground/validator
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)
var trans ut.Translator

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //跨字段
}

func init()  {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
}
// removeTopStruct 删掉返回消息键值的struct名
func removeTopStruct(fields map[string]string)map[string]string  {
	rsp := map[string]string{}
	for field, err := range fields{
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}

	return
}
func main()  {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.POST("/signup", func(c *gin.Context) {
		var sign SignUpForm
		if err := c.ShouldBind(&sign); err !=nil{
			errs,ok := err.(validator.ValidationErrors) // 将err转为validator
			if !ok{  // 转换失败的情况
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":removeTopStruct(errs.Translate(trans)), // 翻译错误信息
			})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	_ = router.Run(":8082")

}