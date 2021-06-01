package fakemodel

import (
	Model "product-api-demo/model"
)

type Repo struct{}

// 檢查 Model 是否符合 IModel contract
// 若不符合，在 build time 就會提出警告
var _ Model.IModel = &Repo{}

// New : Open DB Connections
func New() *Repo {
	return &Repo{}
}
