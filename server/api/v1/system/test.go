package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

func TestT(c *gin.Context) {

	response.Ok(c)
}
