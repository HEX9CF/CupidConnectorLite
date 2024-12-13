package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"stu-campus-network-auto-login/api"
)

const url = "http://a.stu.edu.cn/ac_portal/login.php"

func main() {
	rememberPwd := "0"
	err := api.Login(url, username, password, rememberPwd)
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(" -------- 命令行模式 -------- ")
		fmt.Println("注销：logout")
		fmt.Println("退出程序：exit")
		fmt.Println("重新认证：login")
		for reader.Buffered() > 0 {
			_, err := reader.Discard(reader.Buffered())
			if err != nil {
				log.Println(err)
				break
			}
		}
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		switch input {
		case "logout":
			err := api.Logout(url)
			if err != nil {
				log.Println(err)
				break
			}
		case "login":
			err := api.Login(url, username, password, rememberPwd)
			if err != nil {
				log.Println(err)
				break
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("未知命令！")
		}
	}
}
