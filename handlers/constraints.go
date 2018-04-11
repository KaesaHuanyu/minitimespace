package handlers

const (
	_ = iota
	RequestErr
	BindErr
	HTTPGetErr
	WxApiErr
	JSONErr
	SessionErr
	DatabaseErr

	MiniProgramAppID                     = "wx54708c54c1c17b21"
	MiniProgramAppSecret                 = "d0088d19e9cfba2e91fb41ecda5f2cea"
	WeChatGetAccessToken                 = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	WeChatSendMiniProgramTemplateMessage = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
	WeChatGetUserInfo                    = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	WeChatLoginCredentialsCheck          = "https://api.weixin.qq.com/sns/jscode2session?appid=wx54708c54c1c17b21&secret=d0088d19e9cfba2e91fb41ecda5f2cea&js_code=%s&grant_type=authorization_code"
)
