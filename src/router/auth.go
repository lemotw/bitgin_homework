package router

import (
	"BitginHomework/database"
	"BitginHomework/model"
	"BitginHomework/service/auth"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// user login handler
func Login(c *gin.Context) {
	var loginJSON struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// parse param
	if err := c.BindJSON(&loginJSON); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	ctx := getContext(c)
	signStr, err := auth.SignUser(*ctx, loginJSON.Email, loginJSON.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"signstr": signStr,
	})
}

// signup user handler
func SignUp(c *gin.Context) {
	var signUpJSON struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// parse param
	if err := c.BindJSON(&signUpJSON); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	// generate password hash
	password_hash, err := bcrypt.GenerateFromPassword([]byte(signUpJSON.Password), 12)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	user := model.User{
		Email:     signUpJSON.Email,
		Password:  string(password_hash),
		Role:      "N",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	ctx := getContext(c)
	err = user.Insert(*ctx, database.GetDB())
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"user":   user,
	})
}
