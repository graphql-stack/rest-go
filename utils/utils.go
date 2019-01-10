package utils

import (
	"fmt"
	"github.com/jinzhu/copier"
	"math/rand"
	"regexp"
	"strings"
)

var (
	uUID4RegexString = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID4Regex       = regexp.MustCompile(uUID4RegexString)
)

func MustCopy(dest, source interface{}) {
	err := copier.Copy(dest, source)
	if err != nil {
		panic(err)
	}
}

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func RandomColor() string {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)

	return getHex(r) + getHex(g) + getHex(b)
}

func GenerateAvatar(name string) string {
	return fmt.Sprintf("https://via.placeholder.com/125/%s/%s?text=%s", RandomColor(), RandomColor(), strings.Replace(name, " ", "%20", -1))
}

func IsUuidError(err error) bool {
	return strings.Contains(err.Error(), "invalid input syntax for type uuid")
}

func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "pq: duplicate key value violates unique")
}

func IsUuid4(str string) bool {
	return uUID4Regex.MatchString(str)
}
