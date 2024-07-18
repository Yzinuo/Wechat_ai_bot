package msg

import (
	"strings"

	"Robit_Deepseek/tools/deepseek"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"
)

func HandleMsg(msg *openwechat.Message){
	if msg.IsSendBySelf(){
		return
	}

	var(
		contentText = ""
	)

	if msg.IsText(){
		contentText = trimMsgContent(msg.Content) // 去除空格

		// 获取自己的用户名或昵称
		self, err := msg.Bot().GetCurrentUser()
		if err != nil {
			logrus.Errorf("获取当前用户失败: %v", err)
			return
		}

		// 检查消息是否艾特了自己
		if strings.Contains(msg.Content, "@"+self.NickName) {
			handleTextReplyBypass(msg, contentText)
		}
	}
	
}

func handleTextReplyBypass(msg *openwechat.Message,contentText string){
	var(
		response = ""
		err error
	)
	response,err = deepseek.Getreply(contentText)
	if err != nil{
		logrus.Errorf("获取deepseek响应失败")
		return
	}
	_,error := msg.ReplyText(response)
	if error != nil{
		logrus.Errorf("error of replying")
		return
	}
	return
}


func trimMsgContent(content string) string {
	content = strings.TrimLeft(content, " ")
	content = strings.TrimRight(content, " ")
	return content
}