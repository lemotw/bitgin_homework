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

