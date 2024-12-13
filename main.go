package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	url := "http://a.stu.edu.cn/ac_portal/login.php"
	username := ""
	password := ""
	opr := "pwdLogin"
	rememberPwd := "0"
	data := "opr=" + opr + "&userName=" + username + "&pwd=" + password + "&rememberPwd=" + rememberPwd
	contentType := "application/x-www-form-urlencoded"

	resp, err := http.Post(url, contentType, bytes.NewReader([]byte(data)))
	if err != nil {
		log.Println("请求发送失败！")
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("请求发送成功")
	log.Println("登录成功，用户：" + username + "，您已通过上网认证！")
}
