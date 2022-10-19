package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	data := map[string]interface{}{
		"age":1,
		"name":"xx",
		"email":"12",
		"password":"xx1",
		"re_password":"xx2",
	}
	body, err := json.Marshal(data)
	fmt.Println(body)
	if err != nil {
		return
	}
	resp, err := http.Post("http://localhost:8090/signup", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("response Body:", string(body1))
}
