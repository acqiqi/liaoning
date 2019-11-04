package platform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vgateway/common"
)

func GetRuleLog(c *gin.Context) {
	m := struct {
		DriveKey string `json:"drive_key"`
		RuleTag  string `json:"rule_tag"`
	}{}
	if err := c.BindJSON(&m); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ApiJsonError("格式不正确"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, common.ApiJsonSuccess("获取成功", "123"))
}
