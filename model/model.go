package model

import (
	"github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/rest-go/mysql"
	utils2 "github.com/zcong1993/rest-go/utils"
	"time"
)

var (
	// 3 days
	TOKEN_EXPIRE = time.Hour * 24 * 3
)

// User is user model
type User struct {
	Model
	Username string `json:"username" gorm:"type:varchar(150);unique_index;not null"`
	Password string `json:"-" gorm:"type:varchar(150);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index;not null"`
}

func (u *User) Save() error {
	u.Password = utils.HashPassword(u.Password)
	return mysql.DB.Create(u).Error
}

type Token struct {
	Model
	User   *User  `json:"-"`
	UserID uint   `json:"user_id"`
	Token  string `json:"token" gorm:"type:varchar(100);index;not null"`
}

func (t *Token) IsExpired() bool {
	return time.Now().Sub(t.UpdatedAt) > TOKEN_EXPIRE
}

func (t *Token) Refresh() error {
	return mysql.DB.Model(t).Update("token", utils2.GenerateToken()).Error
}

type Book struct {
	Model
	Title    string  `json:"title" gorm:"type:varchar(150);unique_index;not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Author   *User   `json:"-" gorm:"foreignkey:AuthorID"`
	AuthorID uint    `json:"author_id"`
}
