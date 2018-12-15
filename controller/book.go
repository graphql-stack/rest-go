package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/mysql"
	"github.com/zcong1993/rest-go/utils"
	"net/http"
)

type BookRest struct {
	ginhelper.Rest
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
func (br *BookRest) List(c *gin.Context) {
	limit, offset := ginhelper.DefaultOffsetLimitPaginator.ParsePagination(c)
	q := mysql.DB.Model(new(model.Book)).Order("created_at DESC")

	var data []model.Book
	count, err := utils.PaginationQuery(q, &data, limit, offset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	ginhelper.WithCacheControl(c, common.SHORT_CACHE_DURATION)

	utils.ResponsePagination(c, count, data)
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
func (br *BookRest) Retrieve(c *gin.Context, id string) {
	var book model.Book
	err := mysql.DB.First(&book, "id = ?", id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	ginhelper.WithCacheControl(c, common.SHORT_CACHE_DURATION)

	c.JSON(http.StatusOK, book)
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
func BooksAll(c *gin.Context) {
	limit, offset := ginhelper.DefaultOffsetLimitPaginator.ParsePagination(c)
	q := mysql.DB.Model(new(model.Book)).Order("created_at DESC")

	var data []model.Book
	count, err := utils.PaginationQuery(q, &data, limit, offset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	ginhelper.WithCacheControl(c, common.SHORT_CACHE_DURATION)

	utils.ResponsePagination(c, count, data)
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
func BooksGet(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	err := mysql.DB.First(&book, "id = ?", id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	ginhelper.WithCacheControl(c, common.SHORT_CACHE_DURATION)

	c.JSON(http.StatusOK, book)
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
