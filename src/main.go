package main

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/middleware"
	"BitginHomework/model"
	"BitginHomework/router"
	"context"
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
	ginRouter.POST("/test", func(c *gin.Context) {
		ctx := context.Background()

		// parse param
		var depositeInJSON struct {
			ID int `json:"id"`
		}

		// parse param
		if err := c.BindJSON(&depositeInJSON); err != nil {
			log.Println(err.Error())
			return
		}

		ub := model.UserBalance{
			UserID:  depositeInJSON.ID,
			Balance: 0,
			Point:   0,
		}
		err := ub.Insert(ctx, database.GetDB())
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(ub.ID)
	})

	ginRouter.Run()
}
