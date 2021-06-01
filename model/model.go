package imodel

import (
	MyType "product-api-demo/type"
)

type IModel interface {
	GetOrder(id string) (order MyType.Order, err error)
	NewOrder(order MyType.ReqOrderNew) (result MyType.Order, err error)
	GetProductsByCatID(catid string) (products []MyType.Product, err error)
	GetProduct(id string) (product MyType.Product, err error)
	GetCategories() (cats []MyType.Category, err error)
}
