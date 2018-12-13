package utils

import utils2 "github.com/zcong1993/libgo/utils"

const TOKEN_LEN = 32

func GenerateToken() string {
	tk, err := utils2.GenerateRandomStringURLSafe(TOKEN_LEN)
	if err != nil {
		panic(err)
	}
	return tk
}
