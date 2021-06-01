package mytype

// IReq : interface of HTTP Req
type IReq interface {
	CheckInputArgs() bool
}

// 檢查 ReqOrderNew 是否符合 IReq 合約
var _ IReq = &ReqOrderNew{}

type ReqOrderNew struct {
	Name     string          `json:"name"`
	Address  string          `json:"addr"`
	Mobile   string          `json:"mobile"`
	Products []ProductResult `json:"products"`
}

func (r *ReqOrderNew) CheckInputArgs() bool {
	if r.Name == "" || r.Address == "" || r.Mobile == "" {
		return false
	}

	if len(r.Products) <= 0 {
		return false
	}

	return true
}

type RespCategories struct {
	Categories []Category `json:"cats"`
}

type RespProducts struct {
	Products []Product `json:"products"`
}

type RespProductDetail Product

type RespOrderDetail Order

type RespOrderNew struct {
	Result bool `json:"result"`
	Order
}

type Order struct {
	OrderID    string          `json:"orderid"`
	Name       string          `json:"name"`
	Address    string          `json:"addr"`
	Mobile     string          `json:"mobile"`
	TotalPrice int             `json:"totalprice"`
	Currency   string          `json:"currency"`
	Products   []ProductResult `json:"products"`
}

type Product struct {
	ProductID   string   `json:"productid"`
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	HTML        string   `json:"html"`
	Price       int      `json:"price"`
	Currency    string   `json:"currency"`
	ThumbURL    string   `json:"thumb"`
	PictureURLs []string `json:"pics,omitempty"`
}

type ProductResult struct {
	ProductID string `json:"productid"`
	Amount    int    `json:"amount"`
}

type Category struct {
	CategoryID  string `json:"catid"`
	Description string `json:"desc"`
}
