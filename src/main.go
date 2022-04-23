package main

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/router"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/bitgin?parseTime=true")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	database.SetDB(db)
	log.Println(config.HASH_SECRET)
}

func main() {
	defer database.GetDB().Close()

	// set router
	ginRouter := gin.Default()

	ginRouter.POST("/login", router.Login)
	ginRouter.POST("/signup", router.SignUp)

	ginRouter.Run()
}
