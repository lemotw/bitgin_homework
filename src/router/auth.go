package router

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/model"
	"BitginHomework/service/auth"
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
		internalFaild(c, err.Error())
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
		Role     string `json:"role"`
	}

	// parse param
	if err := c.BindJSON(&signUpJSON); err != nil {
		internalFaild(c, err.Error())
		return
	}

	// check role legal or not
	if !config.RoleLegal(signUpJSON.Role) {
		internalFaild(c, "role not legal")
		return
	}

	// generate password hash
	password_hash, err := bcrypt.GenerateFromPassword([]byte(signUpJSON.Password), 12)
	if err != nil {
		internalFaild(c, err.Error())
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
		internalFaild(c, err.Error())
		return
	}

	// init user balance
	userBalance := model.UserBalance{
		UserID:  user.ID,
		Balance: 0,
		Point:   0,
	}
	userBalance.Insert(*ctx, database.GetDB())

	c.JSON(200, gin.H{
		"status": 200,
		"user":   user,
	})
}
