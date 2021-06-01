package fakemodel

import (
	MyType "product-api-demo/type"
)

var PriceDB map[string]int
var AmountDB map[string]int

// init : Fake DB
func init() {
	PriceDB = map[string]int{
		"000001": 100,
		"000002": 200,
	}

	AmountDB = map[string]int{
		"000001": 10,
		"000002": 20,
	}
}

// GetOrder : 取得訂單明細
func (r *Repo) GetOrder(id string) (order MyType.Order, err error) {
	order = MyType.Order{
		OrderID:    "20210601000001",
		Name:       "王小名",
		Address:    "台北市中山區",
		Mobile:     "0910000000",
		TotalPrice: 5000,
		Currency:   "NTD",
		Products: []MyType.ProductResult{
			{
				ProductID: "000001",
				Amount:    10,
			},
			{
				ProductID: "000002",
				Amount:    20,
			},
		},
	}
	return
}

// NewOrder : 產生新訂單
func (r *Repo) NewOrder(order MyType.ReqOrderNew) (result MyType.Order, err error) {
	result = MyType.Order{
		OrderID:  "20210601000002",
		Name:     order.Name,
		Address:  order.Address,
		Mobile:   order.Mobile,
		Currency: "NTD",
		Products: []MyType.ProductResult{},
	}

	reProd := []MyType.ProductResult{}
	totalPrice := 0

	for _, p := range order.Products {
		if price, ok := PriceDB[p.ProductID]; ok == true {
			if maxAmount, ok2 := AmountDB[p.ProductID]; ok2 == true {
				// 計算可下單數量
				amount := p.Amount
				if maxAmount < p.Amount {
					amount = maxAmount
				}
				reProd = append(reProd, MyType.ProductResult{
					ProductID: p.ProductID,
					Amount:    amount,
				})
				// 計算總價
				totalPrice += amount * price
			}
		}
	}
	result.Products = reProd
	result.TotalPrice = totalPrice

	return
}
