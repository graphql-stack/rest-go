package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/libgo/gin"
)

func Book(c *gin.Context) {
	limit, offset := ginerr.DefaultOffsetLimitPaginator.ParsePagination(c)
	fmt.Println(limit, offset)
	c.Status(204)
}
