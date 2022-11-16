package global

import (
	"github.com/apache/rocketmq-client-go/v2"
	"go_study/dstributed_demo/order/config"
	"go_study/dstributed_demo/order/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	ServerConfig       config.ServerConfig
	OrderSrvClient     proto.OrderClient
	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	MqDelay            rocketmq.Producer
	MqInTimeOut        rocketmq.Producer
	MqTrans            rocketmq.TransactionProducer
)

func init() {
	dsn := "root:123456@tcp(localhost:3306)/mxshop_order_srv?charset=utf8mb4&parseTime=True&loc=Local"

	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 禁用彩色打印
	//	},
	//)

	// 全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

}
