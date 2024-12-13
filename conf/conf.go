package conf

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("初始化环境变量失败！")
		return err
	}

	log.Println("初始化环境变量成功")
	return nil
}
