package muip

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) State(c *gin.Context) {
	allService := a.muip.getAllService()
	c.IndentedJSON(200, allService)
}
