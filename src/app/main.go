package main

import (
	"strconv"

	"example.com/golang-mysql-docker-template/helpers/mysql"
	"example.com/golang-mysql-docker-template/models/task"
	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.ConnectDB(100)
	db.AutoMigrate(&task.Task{})

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	// Init
	task.Init()

	//Index
	router.GET("/", func(ctx *gin.Context) {
		tasks := task.GetAll()
		ctx.HTML(200, "index.html", gin.H{"tasks": tasks})
	})

	//Create
	router.GET("/create/", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		status := ctx.PostForm("status")
		task.Add(name, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		task := task.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"task": task})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		name := ctx.PostForm("name")
		status := ctx.PostForm("status")
		task.Update(id, name, status)
		ctx.Redirect(302, "/")
	})

	//Delete_check
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		task := task.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"task": task})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		task.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
