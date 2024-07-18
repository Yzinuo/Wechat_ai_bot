package Global

import(
	"github.com/eatmoreapple/openwechat"
	"Robit_Deepseek/tools/config"
)

var(
	Conf *config.Conf
	WxSelf *openwechat.Self
	WxFriends openwechat.Friends
	WxGroups openwechat.Groups
)