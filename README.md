# product-api-demo

需求：

>實作內容：假設團隊現在要實作一個簡易的服飾電子商務網站，身為後端工程師的您，請設計並實作包含以下功能的後端 API
>
>1. Product Listing - 使用者可以取得產品清單，用於陳列目前所有產品
>    Required fields for UI display: 產品名稱、價位、圖片
>
>2. Product Details - 使用者可以取得產品明細
>    Required fields for UI display: 產品名稱、價位、詳細說明、圖片
>
>3. Place Order - 使用者可以下訂單並取得是否成功下訂 (請自訂下訂所需欄位、無須考慮金流)
>    Required fields for UI display: 下訂數量、訂單狀態 (成功與否)、訂單編號

## System Requirement

Go 1.16 up

## How to debug

```bash
go run main.go
```

## How to build

```bash
# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o app.exe
```

## How to test

Postman import *\_\_test\_\_/postman/product-api-demo.postman_collection.json*

## How to use

```bash
# linux
app --debugmode -f "./_config/config.yaml"

# windows
app.exe --debugmode -f "./_config/config.yaml"
```

## API

| Method | API                     | Description              |
| :----- | :---------------------- | :----------------------- |
| GET    | /api/cats               | 取得分類清單             |
| GET    | /api/products/:catid    | 取得指定分類的商品清單   |
| GET    | /api/product/:productid | 取得指定商品的明細       |
| POST   | /api/order/new          | 建立新訂單               |
| GET    | /api/order/:orderid     | 根據訂單編號取得訂單明細 |

### [GET] /api/cats

HTTP/1.1  
application/json  

```text
[GET] /api/cats
```

* Response Body

| Key   | Must | Type     | Description          |
| :---- | :--- | :------- | :------------------- |
| cats  | Y    | Object[] | 分類清單             |
| catid | Y    | String   | category unique id   |
| desc  | Y    | String   | category description |

```js
{
    "cats":[
        {
            "catid": "T0001",
            "desc": "最新商品"
        },
        {
            "catid": "T0002",
            "desc": "推薦商品"
        }
    ]
}

```

### [GET] /api/products/:catid

HTTP/1.1  
application/json  

* Request Route Variables

| Key   | Description                     |
| :---- | :------------------------------ |
| catid | 由 */api/cats* 取得的分類 catid |

```text
[GET] /api/products/T0001
```

* Response Body

| Key                | Must | Type     | Description          |
| :----------------- | :--- | :------- | :------------------- |
| products           | Y    | Object[] | 商品清單             |
| products.productid | Y    | String   | 商品 unique id       |
| products.name      | Y    | String   | 商品名稱             |
| products.desc      | Y    | String   | 商品 description     |
| products.price     | Y    | Number   | 價格 (參考幣別)      |
| products.currency  | Y    | String   | 幣別                 |
| products.thumb     | Y    | String   | 商品縮圖(中小圖) URL |

```js
{
    "products":[
        {
            "productid": "000001",
            "name": "商品1",
            "desc": "商品1描述",
            "price": 100,
            "currency": "NTD",
            "thumb": "https://img.product.demo/product/000001/thumb/001.jpg"
        },
        {
            "productid": "000002",
            "name": "商品2",
            "desc": "商品2描述",
            "price": 200,
            "currency": "NTD",
            "thumb": "https://img.product.demo/product/000002/thumb/001.jpg"
        }
    ]
}

```

### [GET] /api/product/:productid

HTTP/1.1  
application/json  

* Request Route Variables

| Key       | Description              |
| :-------- | :----------------------- |
| productid | 商品清單取得的 productid |

```text
[GET] /api/product/000001
```

* Response Body

| Key       | Must | Type     | Description               |
| :-------- | :--- | :------- | :------------------------ |
| productid | Y    | String   | 商品 unique id            |
| name      | Y    | String   | 商品名稱                  |
| desc      | Y    | String   | 商品 description          |
| html      | Y    | String   | 商品詳細說明 HTML         |
| price     | Y    | Number   | 價格 (參考幣別)           |
| currency  | Y    | String   | 幣別                      |
| thumb     | Y    | String   | 商品縮圖(中小圖) URL      |
| pics      | Y    | String[] | 商品展示圖(大圖) URL 清單 |

```js
{
    "productid": "000001",
    "name": "商品1",
    "desc": "商品1描述",
    "html": "<div>商品1詳細描述<b>加強描述</b></div>",
    "price": 100,
    "currency": "NTD",
    "thumb": "https://img.product.demo/product/000001/thumb/001.jpg",
    "pics": [
        "https://img.product.demo/product/000001/pic/001.jpg",
        "https://img.product.demo/product/000001/pic/002.jpg"
    ]
}
```

### [POST] /api/order/new

HTTP/1.1  
application/json  

* Request Body

| Key                | Must | Type     | Description    |
| :----------------- | :--- | :------- | :------------- |
| name               | Y    | String   | 購買人姓名     |
| addr               | Y    | String   | 購買人寄送地址 |
| mobile             | Y    | String   | 購買人手機     |
| products           | Y    | Object[] | 訂單內容       |
| products.productid | Y    | String   | 商品 unique id |
| products.amount    | Y    | Number   | 欲購買數量     |

```js
// [POST] /api/order/new
{
    "name": "王小名",
    "addr": "台北市中山區",
    "mobile": "0910000000",
    "products": [
        {
            "productid": "000001",
            "amount": 11
        },
        {
            "productid": "000002",
            "amount": 30
        }
    ]
}
```

* Response Body

| Key                | Must | Type     | Description                              |
| :----------------- | :--- | :------- | :--------------------------------------- |
| result             | Y    | Boolean  | 訂單建立是否成功? true=成功 / false=失敗 |
| orderid            | N    | String   | 當 result=true 時，回傳建立的訂單編號    |
| name               | N    | String   | 當 result=true 時，回傳購買人姓名        |
| addr               | N    | String   | 當 result=true 時，回傳購買人寄送地址    |
| mobile             | N    | String   | 當 result=true 時，回傳購買人手機        |
| totalprice         | N    | Number   | 當 result=true 時，回傳訂單總價          |
| currency           | N    | String   | 當 result=true 時，回傳幣別              |
| products           | N    | Object[] | 當 result=true 時，回傳訂購內容          |
| products.productid | Y    | String   | 訂購商品 ID                              |
| products.amount    | Y    | Number   | Approved amount (不一定與要求的數量相同) |

```js
{
    "result": true,
    "orderid": "20210601000001",
    "name": "王小名",
    "addr": "台北市中山區",
    "mobile": "0910000000",
    "totalprice": 5000,
    "currency": "NTD",
    "products": [
        {
            "productid": "000001",
            "amount": 10
        },
        {
            "productid": "000002",
            "amount": 20
        }
    ]
}
// or
{
    "result": false
}
```

### [GET] /api/order/:orderid

HTTP/1.1  
application/json  

* Request Route Variables

| Key     | Description |
| :------ | :---------- |
| orderid | 訂單編號    |

```text
[GET] /api/order/20210601000001
```

* Response Body

| Key                | Must | Type     | Description    |
| :----------------- | :--- | :------- | :------------- |
| orderid            | Y    | String   | 訂單編號       |
| name               | Y    | String   | 購買人姓名     |
| addr               | Y    | String   | 購買人寄送地址 |
| mobile             | Y    | String   | 購買人手機     |
| totalprice         | Y    | Number   | 訂單總價       |
| currency           | Y    | String   | 幣別           |
| products           | Y    | Object[] | 訂單細目       |
| products.productid | Y    | String   | 商品代碼       |
| products.amount    | Y    | Number   | 訂購數量       |

```js
{
    "orderid": "20210601000001",
    "name": "王小名",
    "addr": "台北市中山區",
    "mobile": "0910000000",
    "totalprice": 5000,
    "currency": "NTD",
    "products": [
        {
            "productid": "000001",
            "amount": 10
        },
        {
            "productid": "000002",
            "amount": 20
        }
    ]
}
```
