package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/libgo/gin"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/utils"
	"net/http"
)

func Book(c *gin.Context) {
	limit, offset := ginerr.DefaultOffsetLimitPaginator.ParsePagination(c)
	fmt.Println(limit, offset)
	c.Status(204)
}

// Test pagination for /p
// @Summary Test pagination
// @Description
// @Accept  json
// @Produce  json
// @Param   default limit query int false "limit" default(100)
// @Param   default offset query int false "offset" default(0)
// @Success 200 {array} model.User
// @Failure 500 "StatusInternalServerError"
// @Router /p [get]
func Test(c *gin.Context) {
	limit, offset := ginerr.DefaultOffsetLimitPaginator.ParsePagination(c)

	q := mysql.DB.Model(new(model.User)).Where("1=1")
	var data []model.User
	count, err := utils.PaginationQuery(q, &data, limit, offset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	utils.ResponsePagination(c, count, data)
}
