package router

import (
	"BitginHomework/database"
	"BitginHomework/model"
	"BitginHomework/service/trade"
	"BitginHomework/service/trade/event"

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

	// trade context to apply
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
	// get context user database
	ctx := getContext(c)
	user := getUser(c)
	if user == nil {
		internalFaild(c, "user get problem")
		return
	}
	db := database.GetDB()
	if db == nil {
		panic("db not set")
	}

	// parse param
	var PayBalanceJSON struct {
		Balance float64 `json:"balance"`
		Point   int     `json:"point"`
	}
	if err := c.BindJSON(&PayBalanceJSON); err != nil {
		internalFaild(c, err.Error())
		return
	}

	// start tx
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

	// trade context to apply
	tc := trade.TradeContext{
		EventList: []trade.EventFunc{
			event.RoleDiscount,
			event.PointDiscount,
			event.CostPoint,
			event.CostBalance,
		},
	}
	err = tc.Apply(*ctx, *user, *userbalance, PayBalanceJSON.Balance, PayBalanceJSON.Point)
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
