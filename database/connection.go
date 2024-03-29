package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	godotenv.Load()
	dbPassword2 := os.Getenv("DB_PASSWORD")

	// 데이터베이스 연결 정보 설정
	dbUser := "root"          // MySQL 사용자 이름
	dbPassword := dbPassword2 // MySQL 암호
	dbHost := "localhost"     // MySQL 호스트
	dbPort := "3306"          // MySQL 포트
	dbName := "sys"           // 데이터베이스 이름

	// MySQL 데이터베이스에 연결하는 데이터베이스 연결 문자열 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

}
