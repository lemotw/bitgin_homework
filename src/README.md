# API 
### Auth User Notice
如果需要權限的 API 需要在 Header 設置一個欄位 token:"token_str"，token_str 用 /login 取得。

## /login
登入並獲得 tokenstr
- email(string)
- password(string)
## /signup
註冊帳戶
- email(string) Email unique
- password(string)
- role(string) 該帳戶為正常用戶還是 VIP
## /deposite/in
儲值 API，直接將錢儲入帳戶，後面可以寫活動家在 trade cycle 裡面
- balance (float64): 要儲值多少 
## /pay
付費 API ，如果餘額不足會回傳錯誤，而點數如果不夠也不會被使用，只會用他有的但最高就是指定的 point 數量
- balance (float64): 要付多少
- point (int): 使用多少點數


# Dir tour

## config
放 auth 相關的 hash secret，還有折扣設定。

## database
讓大家可以抓到同個 db pool 

## middleware
middleware function 會放這邊，而處理玩的資料通常放 `gin.Context` 裡面。

## model
資料庫對應 struct，由 `xo` 自動產生

## router
放對應的 API `HandleFunc`

## service
auth 部份實做簽 jwt、trade 部份實做 event cycle 下面會細講這部份


# Trade Cycle
這塊設計是給各個活動設計的，接著細講。service 下的 trade 有 TradeContext 這個 context 在每次金流活動的時候創建，創的時候定義你所想的 EventList ，他就會依據你排定的活動依序執行，而活動的 func 只要符合下面定義就行。

```go
type EventFunc func(u model.User, trade *model.TradeRecord, val float64) error
```

## Apply (TradeContext)
```go
func (tc *TradeContext) Apply(ctx context.Context, u model.User, balance model.UserBalance, cost float64, point int) error
```
u 代表哪個使用者參與這次金流活動， balance　是該使用者的餘額狀態，cost 代表這次金流會變動多少 balance，<b>切記如果是付費要用負的，或是使用 CostBalance 在最後將他變為負的</b> ，point 同 cost 代表這次要用多少點數。

## UpdateBalance (TradeContext)
``` go
func (tc *TradeContext) UpdateBalance(ctx context.Context, balance model.UserBalance, tx sqlx.Tx) error
```
balance 指的是現在 user balance，而該 func 會將 TradeContext Apply 過後的結果儲存到資料庫。


## Example:
```go
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
```