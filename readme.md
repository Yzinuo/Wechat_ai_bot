# 基于deepseek支持的微信ai机器人

# 研发初衷
openWechat框架十分强大和易扩展。正好deepseek免费送了大量token，就想着自己写写玩一玩，实现起来也非常容易。当然这个机器人还可以实现非常多的功能，比如管理群聊，定时发送早安晚安给你女朋友等等，需要您自己拓展了。

----

# 部署说明
clone 项目到本地，然后进入项目目录，执行下面的指令
在/configfile/prod.yaml中输入你的api

```
go mod tidy //下载依赖

go run main.go // 运行程序
```

# 功能介绍
初步实现了艾特机器人，自动调用deepseek回话功能。
同时你也可以在POST请求的message中进行初步调教，比如：你现在是.....，你的主人是....。

**使用截图**：

![alt text](image.png)

**虽然这是一个很简单的项目，但是如果能帮到你，或者能要你感到好玩，请你留下你的Star,感谢您的支持 :) :) :)**