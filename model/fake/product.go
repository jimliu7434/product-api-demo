package fakemodel

import (
	MyType "product-api-demo/type"
)

// GetProductsByCatID : 根據 CategoryID 取得 商品 清單
func (r *Repo) GetProductsByCatID(catid string) (products []MyType.Product, err error) {
	products = []MyType.Product{
		{
			ProductID:   "000001",
			Name:        "商品1",
			Description: "商品1描述",
			Price:       100,
			Currency:    "NTD",
			ThumbURL:    "https://img.product.demo/product/000001/thumb/001.jpg",
		},
		{
			ProductID:   "000002",
			Name:        "商品2",
			Description: "商品2描述",
			Price:       200,
			Currency:    "NTD",
			ThumbURL:    "https://img.product.demo/product/000002/thumb/001.jpg",
		},
	}
	return
}

// GetProduct : 根據 productid 取得 商品明細
func (r *Repo) GetProduct(id string) (product MyType.Product, err error) {
	product = MyType.Product{
		ProductID:   "000002",
		Name:        "商品2",
		Description: "商品2描述",
		HTML:        "<div>商品1詳細描述<b>加強描述</b></div>",
		Price:       200,
		Currency:    "NTD",
		ThumbURL:    "https://img.product.demo/product/000002/thumb/001.jpg",
		PictureURLs: []string{
			"https://img.product.demo/product/000001/pic/001.jpg",
			"https://img.product.demo/product/000001/pic/002.jpg",
		},
	}
	return
}
