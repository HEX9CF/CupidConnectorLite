# Cupid Connector Lite

## 项目简介

一款用于自动登录汕头大学校园网的轻量级命令行工具，是 Cupid Connector 的轻量版，通过发送登录请求并处理响应，实现自动化登录功能，方便汕大师生登录校园网。

该工具支持自动登录、重新认证、注销账号等功能，同时支持初始化配置文件，简化用户的配置过程。不需要用户手动打开浏览器，输入账号密码，只需双击运行程序即可完成登录操作。将程序添加到开机启动项，即可实现开机自动登录校园网。不同于 Python 脚本，该工具是一个独立的可执行文件，开箱即用，无需安装 Python 环境，方便用户使用。

A lightweight tool for automatically logging into the Shantou University campus network. It automates the login process by sending login requests and handling responses, making it convenient for Shantou University students and staff to log into the campus network.

## 软件版本

### Cupid Connector

完整版，带 GUI 程序

仓库地址：[https://github.com/HEX9CF/CupidConnector](https://github.com/HEX9CF/CupidConnector)

### Cupid Connector Lite

轻量版分支，命令行工具

仓库地址：[https://github.com/HEX9CF/CupidConnectorLite](https://github.com/HEX9CF/CupidConnectorLite)

## 功能简介

- **自动登录**：通过发送登录请求，并处理登录响应，实现自动化登录。
- **重新认证**：用户可以在需要时重新进行登录认证。
- **注销账号**：用户可以注销当前登录的账号，退出校园网。
- **初始化配置**：支持初始化配置文件，简化用户的配置过程。

## 使用方法

1. **运行程序**：在 Release 页面下载最新的可执行文件，双击运行程序。
2. **初始化配置**：首次运行程序时，会自动生成配置文件。用户可以根据提示填写校园网用户名和密码。同时，程序会询问是否在认证成功后自动退出程序，用户可以根据需求选择是否开启该功能。
3. **自动登录**：程序会自动发送认证请求，登录校园网，登录成功后会显示登录成功的提示信息。
4. **自动退出**：如果用户开启了认证成功自动退出功能，程序会在认证成功后立即自动退出，不会进入命令模式。如果没有认证成功，仍会进入命令模式。
5. **命令模式**: 如果用户没有开启认证成功自动退出功能，在自动登录完成后，程序会进入命令模式，用户可以输入命令来执行相应的操作。
6. **开机自启**：用户可以将程序添加到开机启动项，实现开机自动登录校园网。具体方法请参考下文。

## 命令列表

- `login`：登录账号。
- `logout`：注销账号。
- `init`: 重新初始化配置文件。
- `exit`：退出程序。

![image](https://github.com/user-attachments/assets/22ad0a14-e7ec-4490-932c-89d906c2a3a0)

## 配置文件

配置文件为 `stu.env`，位于程序运行目录下。 配置文件内容如下：

```env
# 校园网认证接口
STU_URL=http://a.stu.edu.cn/ac_portal/login.php

# 校园网用户名
STU_USERNAME=username

# 校园网密码
STU_PASSWORD=password

# 认证成功后自动退出程序
AUTO_EXIT=FALSE
```

## 开机自启

用户可以将程序添加到开机启动项，实现开机自动登录校园网。步骤如下：

1. 将可执行文件移动到一个固定的目录，如 `C:\Program Files\Cupid Connector`。
2. 右键点击可执行文件 `cupid-connector.exe`，选择“创建快捷方式”。
3. 按 `Win + R` 打开运行对话框。
4. 输入 `shell:startup` 并按回车，打开启动文件夹。 
5. 将刚才创建的快捷方式移动到这个文件夹中。

如果需要取消开机自启，只需将快捷方式移出启动文件夹即可。

## 趣闻

项目名称 "Cupid Connector" 灵感来源于游戏《Death Stranding》中的 "丘比特连接器"（Q-pid）。在游戏中，丘比特连接器是一种设备，用于连接和扩展开罗尔网络，使得玩家能够与更广阔的世界进行通信和互动。这个设备在游戏中扮演着连接和桥梁的角色，象征着连接和沟通的重要性。Cupid Connector 作为一个自动登录校园网的工具，同样扮演着连接的角色。它将学生和教职工与校园网络连接起来，使他们能够方便地访问网络资源，寓意它能够像游戏中的丘比特连接器一样，为用户架起一座连接的桥梁。

## 参与贡献

- 如果你觉得这个项目对你有所帮助，欢迎给个 Star⭐️。
- 如果你有任何问题或建议，欢迎提交 Issue。
- 如果你有兴趣参与贡献代码，欢迎提交 Pull Request。

## 许可与声明

本项目采用 GPL-3.0 license 进行许可，详情请参阅 [LICENSE](LICENSE) 文件。

本项目仅供学习交流使用，不得用于任何商业用途。使用本项目所造成的任何后果，由使用者自行承担，与项目作者无关。

本项目是开源项目，用户可以自行审查源代码，确认程序的安全性。用户的账号密码等信息只会保存在用户的计算机上，仅用于校园网登录认证。
