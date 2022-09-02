package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/model"
)

var (
    db  *gorm.DB
    err error
)

// server立ち上げ時に接続
func Init() {
    db, err = gorm.Open("postgres", "host=db port=5432 user=root dbname=test_db password=root sslmode=disable")
    if err != nil {
        panic(err)
    }

	autoMigration()
}


// psql操作時にdb呼び出しする用
func GetDB() *gorm.DB {
    return db
}

// db閉じる
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

//自動でフィールド作成
func autoMigration() {
    db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Plan{})
    //db.AutoMigrate(&model.Calendar{})
}
