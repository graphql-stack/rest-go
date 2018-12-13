package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/service"
	"net/http"
	"strings"
)

func Auth(c *gin.Context) {
	a := c.GetHeader("Authorization")
	token := strings.Replace(a, "Bearer ", "", 1)

	u, err := service.GetUserByToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.TOKEN_ERROR)
		return
	}

	c.Set(common.AUTH_CONTEXT_KEY, u)
	c.Next()
}
