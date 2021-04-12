package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	// 检测用户是否存在
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUser(c *gin.Context) {
	var username string
	_ = c.ShouldBindJSON(username)
	data := model.GetUser(username)
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	//
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	fmt.Println("username,role", data.Username, data.Role)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		fmt.Println(data)
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	//
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
