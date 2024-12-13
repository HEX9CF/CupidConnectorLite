/*
@author Howard Zheng <hex9cf@aliyun.com>
*/
package main

import (
	"bufio"
	"cupid-connector/api"
	"cupid-connector/conf"
	"fmt"
	"log"
	"os"
	"strings"
)

const rememberPwd = "0"

func main() {
	log.SetOutput(os.Stdout)
	err := conf.InitEnv()
	if err != nil {
		panic(err)
	}

	err = api.Login(conf.Url, conf.Username, conf.Password, rememberPwd)
	if err != nil {
		log.Println(err)
	}
	if err == nil && conf.AutoExit == "TRUE" {
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(" -------- 命令列表 -------- ")
		fmt.Println(" 退出程序：exit")
		fmt.Println(" 重新认证：login")
		fmt.Println(" 注销账号：logout")
		fmt.Println(" 初始化配置：init")
		fmt.Println(" ------------------------- ")
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
			err := api.Logout(conf.Url)
			if err != nil {
				log.Println(err)
				break
			}
		case "login":
			err := api.Login(conf.Url, conf.Username, conf.Password, rememberPwd)
			if err != nil {
				log.Println(err)
				break
			}
		case "init":
			err := conf.CreateEnv()
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
