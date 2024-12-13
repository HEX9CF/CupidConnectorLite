package conf

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"stu-campus-network-auto-login/model"
	"stu-campus-network-auto-login/utils"
)

func createEnv() error {
	reader := bufio.NewReader(os.Stdin)
	for reader.Buffered() > 0 {
		_, err := reader.Discard(reader.Buffered())
		if err != nil {
			log.Println(err)
			break
		}
	}

	u := model.User{}
	fmt.Println("请输入校园网认证信息：")
	fmt.Print("用户名：")
	u.Username, _ = reader.ReadString('\n')
	u.Username = strings.TrimSpace(u.Username)
	fmt.Print("密码：")
	u.Password, _ = reader.ReadString('\n')
	u.Password = strings.TrimSpace(u.Password)

	autoExitVal := defaultAutoExit
	fmt.Print("认证成功后是否自动退出程序(Y/N)？")
	autoExit, _ := reader.ReadString('\n')
	autoExit = strings.TrimSpace(autoExit)
	autoExit = strings.ToLower(autoExit)
	if autoExit == "y" {
		autoExitVal = "TRUE"
	}

	log.Println("正在创建配置文件...")
	v, err := utils.IsFileExists(envPath)
	if err != nil {
		log.Println("创建配置文件失败！")
		return err
	}
	if v {
		err := os.Remove(envPath)
		if err != nil {
			log.Println("删除旧配置文件失败！")
			return err
		}
	}
	file, err := os.Create(envPath)
	if err != nil {
		log.Println("创建配置文件失败！")
		return err
	}
	defer file.Close()

	content := "STU_URL=" + defaultUrl + "\n" +
		"STU_USERNAME=" + u.Username + "\n" +
		"STU_PASSWORD=" + u.Password + "\n" +
		"AUTO_EXIT=" + autoExitVal + "\n"

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("写入配置文件失败！")
		return err
	}

	log.Println("创建配置文件成功！")
	return nil
}
