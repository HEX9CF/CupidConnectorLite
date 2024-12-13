package main

import (
	"log"
	"stu-campus-network-auto-login/api"
)

const url = "http://a.stu.edu.cn/ac_portal/login.php"

func main() {
	err := api.Login(url, username, password, "0")
	if err != nil {
		log.Println(err)
	}
}
