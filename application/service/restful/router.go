package restful

import (
	"vgateway/application/service/restful/controller/platform"
)

//路由初始化
func init() {

	router_platform := Router.Group("/platform")
	{
		//router_platform.Use(middleware.ValidateAuthMiddleware([]string{
		//	"/Public/WechatSmallLogin",
		//}))
		router_platform.POST("??", platform.GetRuleLog)

	}

	//Router.POST("/Main/Index", manager.NewsList)

}
