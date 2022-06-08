package task

import (
	"example.com/golang-mysql-docker-template/helpers/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	gorm.Model
	Name   string
	Status string
}

//DB初期化
func Init() {
	db := mysql.ConnectDB(100)
	db.AutoMigrate(&Task{})
	defer db.Close()
}

//DB追加
func Add(name string, status string) {
	db := mysql.ConnectDB(100)
	db.Create(&Task{Name: name, Status: status})
	defer db.Close()
}

//DB更新
func Update(id int, name string, status string) {
	db := mysql.ConnectDB(100)
	var task Task
	db.First(&task, id)
	task.Name = name
	task.Status = status
	db.Save(&task)
	db.Close()
}

//DB削除
func Delete(id int) {
	db := mysql.ConnectDB(100)
	var task Task
	db.First(&task, id)
	db.Delete(&task)
	db.Close()
}

//DB取得(ALL)
func GetAll() []Task {
	db := mysql.ConnectDB(100)
	var tasks []Task
	db.Order("created_at asc").Find(&tasks)
	db.Close()
	return tasks
}

//DB取得(単一)
func GetOne(id int) Task {
	db := mysql.ConnectDB(100)
	var task Task
	db.First(&task, id)
	db.Close()
	return task
}
