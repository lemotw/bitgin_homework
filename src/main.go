package main

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/middleware"
	"BitginHomework/model"
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

	ginRouter.POST("/test", middleware.WithContext, middleware.WithUser, func(c *gin.Context) {
		user, userExist := c.Get("user")
		if !userExist {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "user get problem",
			})
		}

		c.JSON(200, gin.H{
			"status": 200,
			"user":   user.(*model.User),
		})
	})
	ginRouter.POST("/login", middleware.WithContext, router.Login)
	ginRouter.POST("/signup", middleware.WithContext, router.SignUp)
	ginRouter.GET("/discount", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"dis": model.USERROLE_DISCOUNT,
		})
	})

	ginRouter.Run()
}
