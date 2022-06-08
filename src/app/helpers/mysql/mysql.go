package mysql

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB(count uint8) *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("DB open error:", err)

		// リトライ回数満了
		if count <= 0 {
			log.Fatal("time out")
		}

		// リトライ(再帰)
		log.Println("retry... count:", count)
		time.Sleep(time.Second * 2)
		count--
		return ConnectDB(count)
	}

	log.Println("db connected!")
	return db
}
