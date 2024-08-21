package main

import (
	"fmt"
	"time"

	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	sqle "github.com/dolthub/go-mysql-server"
)

var (
	dbName = "mydb"
	port   = 3309
)

type FakeMySQL struct {
	DB      string
	Address string
	Server  *memory.DbProvider
}

type CertificationInfo struct {
	ID       int `gorm:"index;primaryKey;autoIncrement"`
	Username string
	Password string
}

func (receiver CertificationInfo) TableName() string {
	return "certification_info"
}

func main() {
	s := createMySQLServer(dbName, port)
	go func() {
		time.Sleep(10 * time.Second)
		_ = s.Close()
	}()

	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root@tcp(127.0.0.1:%d)/"+dbName, port)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&CertificationInfo{}); err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		err = db.Create(&CertificationInfo{Username: "root", Password: uuid.NewString()}).Error
		if err != nil {
			panic(err)
		}
	}

	var result []CertificationInfo
	err = db.Select("id", "username", "password").Find(&result).Error
	if err != nil {
		panic(err)
	}
	for _, i := range result {
		fmt.Println(i)
	}

	time.Sleep(11 * time.Second)
}

func createMySQLServer(dbName string, port int) *server.Server {
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	pro := memory.NewDBProvider(db)
	engine := sqle.NewDefault(pro)

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("0.0.0.0:%d", port),
	}

	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		panic(err)
	}
	return s
}
