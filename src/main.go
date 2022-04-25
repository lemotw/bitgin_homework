package main

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/middleware"
	"BitginHomework/router"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	db, err := sqlx.Open("mysql", "root:root@tcp(db:3306)/bitgin?parseTime=true")
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

	ginRouter.POST("/login", middleware.WithContext, router.Login)
	ginRouter.POST("/signup", middleware.WithContext, router.SignUp)
	ginRouter.POST("/deposite/in", middleware.WithContext, middleware.WithUser, router.DepositeInBalance)
	ginRouter.POST("/pay", middleware.WithContext, middleware.WithUser, router.PayByBalance)

	ginRouter.Run()
}
