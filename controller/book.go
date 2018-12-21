package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/utils"
	"net/http"
)

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
	limit, offset := ginhelper.DefaultOffsetLimitPaginator.ParsePagination(c)

	q := mysql.DB.Model(new(model.User)).Where("1=1")
	var data []model.User
	count, err := utils.PaginationQuery(q, &data, limit, offset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	utils.ResponsePagination(c, count, data)
}

type BookView struct {
	ginhelper.Rest
	ginhelper.RestView
}

func (b *BookView) GetQuerySet() *gorm.DB {
	return mysql.DB.Model(new(model.Book))
}

func (b *BookView) GetSerializers() interface{} {
	var book []model.Book
	return &book
}

func (b *BookView) GetSerializer() interface{} {
	return &model.Book{}
}

// Books query all for /books
// @Summary Books pagination query all
// @Description
// @Accept  json
// @Produce  json
// @Param   default limit query int false "limit" default(100)
// @Param   default offset query int false "offset" default(0)
// @Success 200 {array} model.Book
// @Failure 500 "StatusInternalServerError"
// @Router /books [get]
func DocsBookList() {

}

// Books query by id for /books/:id
// @Summary Books query by id
// @Description
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "Book id"
// @Success 200 {object} model.Book
// @Failure 500 "StatusInternalServerError"
// @Router /books/{id} [get]
func DocsBookRetrieve() {

}
