package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// func ConnectXorm(host string, port string, database string, user string, pass string, options string) (db *xorm.Engine, err error) {
// 	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&"+options)
// }

func Connect() (db *gorm.DB) {

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbdatabase := os.Getenv("DB_DATABASE")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbdatabase+"?charset=utf8")
	// gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	return db
}

// func ConnectSql(database string, user string, pass string) (db *gorm.DB, err error) {
// 	return gorm.Open("mysql", user+":"+pass+"@/"+database)
// }
