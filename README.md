# order_demo
Golang 訂餐系統API 練習

## Docker快速佈署API環境

```
docker-compose up -d
```

## Swagger API Demo

```
http://localhost:8080/swagger/index.html
```

## schema

**account (會員帳號)**

    id, account, password, status, balance

**transfer (充值紀錄只能存入、扣款 balance會跟著變動)**

    id, accountId, [+-amount], transferTime

**goods (上架餐點)**

    id, goods, amount, status, orderTime

**report (每日訂餐報表)**

    id, date, memberid, goods, createAt

**order (訂單紀錄)**

    id, accountId, goodsId, status, orderTime

**announcement (公告內容)**

    id, comment, status, date 唯一

**id(PK, AI), amount(decimal(12, 2)), status(enum('0', '1'))**

## handler 功能以下

### admin router

設一個 Group admin router

- 新增公告 (addAnno) 

- 修改公告 (editAnno) 可以修改參數 comment、status

- 新增、刪除、更新、查看餐點 (goods)

- 查看今天所有訂的便當，明細要有加總 (list)

```
{
	user1: {goods1: 20, goods2: 30},
	user2: {goods1: 20, goods2: 40},
	summary: {count: 4, total: 110}
}
```

用一個middleware的驗證 role=1 才能做以上的動作，其他一律導向到 401 你沒有權限

執行產生今天的訂餐報表(只能做一次，做完後使用者不能再取消訂錯的便當，也不能再追加訂便當) 使用announcement


### user router

以下使用者功能 必需先登入 使用 jwt 派給 bearer token 才能操作

設一個 Group api router

- 充值金額  deposit

- 訂便當  order

- 查看自己訂的便當 GetTodayOrder

- 取消訂錯的便當 DeleteOrder
