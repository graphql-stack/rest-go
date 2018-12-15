package common

import "time"

const (
	AUTH_CONTEXT_KEY   = "AUTH_CONTEXT_KEY"
	HEADER_TOTAL_COUNT = "X-TOTAL-COUNT"
)

var (
	LONG_CACHE_DURATION  = time.Hour * 24 * 3
	SHORT_CACHE_DURATION = time.Minute * 1
)
