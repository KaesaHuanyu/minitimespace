package handlers

const (
	MiniProgramAppID                     = "wx54708c54c1c17b21"
	MiniProgramAppSecret                 = "d0088d19e9cfba2e91fb41ecda5f2cea"
	WeChatGetAccessToken                 = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	WeChatSendMiniProgramTemplateMessage = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
	WeChatGetUserInfo                    = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"

	_ = iota
	RequestErr
)
