作为一次Gin框架的练手小项目。前端采用Jquery和Bootstrap5，数据库使用Mysql。来实现一个简单的TodoList网页。用户可以**登录，注册**。并在网页中查看自己的待办事项并可以**添加，删除，修改**任务。

## 文件结构

```
gin
├─ .idea
│  ├─ gin.iml
│  ├─ modules.xml
│  └─ workspace.xml
├─ dao        //数据库操作和数据定义
│  ├─ data.go
│  └─ mysql.go
├─ go.mod
├─ go.sum
├─ handler   //路由处理器
│  ├─ handler.go
│  └─ middleware.go
├─ main.go   //主函数
├─ README.md
├─ router    //路由
│  └─ router.go
├─ static    //静态资源
│  ├─ css
│  │  ├─ bootstrap.min.css
│  │  ├─ bootstrap.min.css.map
│  │  └─ login.css
│  └─ js
│     ├─ bootstrap.bundle.min.js
│     ├─ bootstrap.bundle.min.js.map
│     ├─ bootstrap.esm.min.js
│     ├─ bootstrap.esm.min.js.map
│     ├─ bootstrap.min.js
│     ├─ bootstrap.min.js.map
│     ├─ list.js
│     └─ login.js
└─ templates  //HTML文件
   ├─ checklist.html
   └─ login.html

```



## 数据定义

本项目使用结构体来表示用户数据和清单中的任务，并根据数据库和前端json交换，基于go的反射机制来处理数据之间的传递。

```go
type User struct {
	Id       int    `db:"user_id" json:"user_id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

type Todo struct {
	UserId int    `db:"user_id" json:"user_id"`
	Id     int    `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Status int    `db:"status" json:"status"`
}
```

