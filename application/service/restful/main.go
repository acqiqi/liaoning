package restful

import (
	"github.com/gin-gonic/gin"
	"vgateway/common"
)

var Router = gin.Default()

func InitRestfulServer() {
	Router.Run(common.GetConfig("restful_port"))
}
