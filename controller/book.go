package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/reflect"
	utils2 "github.com/zcong1993/libgo/utils"
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

func (b *BookView) GetModel(isMany bool) interface{} {
	if isMany {
		var book []model.Book
		return &book
	}
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

func BookBatch(ctx *gin.Context) {
	batchIds := ctx.Query("ids")
	if batchIds == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ids, ok := ginhelper.NormalizeBatchQuery(batchIds)
	if !ok {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var books []model.Book
	err := mysql.DB.Where("id in (?)", ids).Find(&books).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	dataMap := reflecthelper.Slice2Map(books, "ID")
	res := make(map[string]interface{}, len(ids))

	for _, id := range ids {
		b, ok := dataMap[uint(utils2.Str2uint64(id))]
		if !ok {
			res[id] = nil
			continue
		}
		res[id] = b
	}

	ctx.JSON(http.StatusOK, res)
}
