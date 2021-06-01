package fakemodel

import (
	MyType "product-api-demo/type"
)

// GetTypes : 取得 Types
func (r *Repo) GetCategories() (cats []MyType.Category, err error) {
	cats = []MyType.Category{
		{
			CategoryID:  "T0001",
			Description: "最新商品",
		},
		{
			CategoryID:  "T0002",
			Description: "推薦商品",
		},
	}
	return
}
