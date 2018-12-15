package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/rest-go/common"
	"net/http"
)

var DefaultPaginator = &ginhelper.OffsetLimitPaginator{DefaultNumPerPage: 100}

func ResponsePagination(ctx *gin.Context, count int, data interface{}) {
	ctx.Header(common.HEADER_TOTAL_COUNT, utils.Num2String(count))
	ctx.JSON(http.StatusOK, data)
}

func PaginationQuery(db *gorm.DB, t interface{}, limit, offset int) (int, error) {
	var count int
	count = 0
	err := db.Count(&count).Error

	if err != nil {
		return count, err
	}

	err = db.Limit(limit).Offset(offset).Find(t).Error

	if err != nil {
		return count, err
	}

	return count, nil
}
