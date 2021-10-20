package handler

import (
	"fmt"
	data "gin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	data.InitDB()
	fmt.Println("use /")
	_, err := c.Cookie("login")
	if err != nil {
		fmt.Println("no cookie")
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/list")
	}
}

func GetListHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "checklist.html", nil)
}

func AddNewList(c *gin.Context) {
	var todo data.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error bind", err)
		return
	}
	//todo.UserId=
	fmt.Println(todo)
	ok, err := data.CreatNewTodo(todo)
	if err != nil {
		fmt.Println("error creat", err)
	}
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加成功!",
		})
	}
}

func DeleteList(c *gin.Context) {
	title := c.Param("title")
	fmt.Println("title" + title)
	ok, err := data.DeleteTodo(title)
	if err != nil {
		fmt.Println(err.Error())
	}
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功!",
		})
	}
}

func UpdateList(c *gin.Context) {
	title := c.PostForm("title")
	status := c.PostForm("status")
	ok, err := data.UpdateTodo(title, status)
	if err != nil {
		fmt.Println(err.Error())
	}
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "修改成功!",
		})
	}
}

func GetAllList(c *gin.Context) {
	name := c.Query("name")
	fmt.Println("get:", name)
	list, id := data.GetList(name)
	fmt.Println("Id：", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": list,
		"id":  id,
	})
}

func SignCheck(c *gin.Context) {
	user := new(data.User)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	ok := data.IsUserExsit(user)
	if ok {
		ok, id, err := data.CreatNewUser(user)
		if ok {
			c.SetCookie("login", user.Name, 3600, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{
				"type": 1,
				"id":   id,
				"msg":  "注册成功!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"type": 3,
				"id":   id,
				"msg":  "注册失败" + err.Error(),
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"type": 3,
			"msg":  "注册失败,该用户名已被注册!",
		})
	}
}

func LoginCheck(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Println("name:" + name)
	fmt.Println("user:" + name)
	user := &data.User{
		Name:     name,
		Password: password,
	}
	fmt.Println(user.Name)
	fmt.Println(user.Password)
	ok, _ := data.CheckUser(user)
	if ok {
		c.SetCookie("login", user.Name, 3600, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{
			"type": 1,
			"msg":  "登录成功!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"type": 3,
			"msg":  "登录失败，请重新输入用户名和密码!",
		})
	}
}
