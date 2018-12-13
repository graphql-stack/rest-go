package utils

import "github.com/zcong1993/libgo/gin"

var DefaultPaginator = &ginerr.OffsetLimitPaginator{DefaultNumPerPage: 100}
