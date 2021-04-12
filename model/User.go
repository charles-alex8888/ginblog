package model

import (
	"encoding/base64"
	"fmt"
	"ginblog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20) not null" json:"username"`
	Password string `gorm:"type:varchar(20) not null" json:"password"`
	Role     int    `gorm:"type:int " json:"Role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 查询单个用户
func GetUser(name string) User {
	var user User
	db.Where("username = ?", name).Find(&user)
	fmt.Println(user)
	return user
}

// 新增用户
func CreateUser(data *User) int {
	// data.Password = ScrypPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func (u *User) BeforeSave() {
	u.Password = ScrypPw(u.Password)
}
func ScrypPw(password string) string {
	const KeyLen = 10
	// salt := make([]byte, 8)
	salt := []byte{12, 34, 68, 65, 79, 93, 81, 44}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}
