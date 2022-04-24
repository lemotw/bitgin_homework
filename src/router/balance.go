package router

import (
	"BitginHomework/database"
	"BitginHomework/model"
	"BitginHomework/service/trade"

	"github.com/gin-gonic/gin"
)

// deposite in user balance
func DepositeInBalance(c *gin.Context) {
	ctx := getContext(c)
	user := getUser(c)
	if user == nil {
		internalFaild(c, "user get problem")
		return
	}

	// parse param
	var depositeInJSON struct {
		Balance float64 `json:"balance"`
	}

	// parse param
	if err := c.BindJSON(&depositeInJSON); err != nil {
		internalFaild(c, err.Error())
		return
	}

	// get database
	db := database.GetDB()
	if db == nil {
		panic("db not set")
	}

	// get tx
	tx, err := db.BeginTxx(*ctx, nil)
	if err != nil {
		internalFaild(c, err.Error())
		return
	}

	// get balance
	userbalance, err := model.UserBalanceByUserID(*ctx, tx, user.ID)
	if err != nil {
		internalFaild(c, err.Error())
		return
	}

	tc := trade.TradeContext{}
	err = tc.Apply(*ctx, *user, *userbalance, depositeInJSON.Balance, 0)
	if err != nil {
		internalFaild(c, err.Error())
		return
	}

	// update to db
	err = tc.UpdateBalance(*ctx, *userbalance, *tx)
	if err != nil {
		internalFaild(c, err.Error())
	}

	if err = tx.Commit(); err != nil {
		internalFaild(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
	})
}

// pay api
func PayByBalance(c *gin.Context) {
}
