# json-go-struct
convert json to go struct(one nest struct and multi normal separate structs),you can check my another project https://github.com/GrayOxygen/json-go-struct-app/ just a app version of this project
> 将json转为golang的struct，目前开源的项目、工具中，只找到转成嵌套的struct形式，于是写了个转成非嵌套的struct，you will get result like this
``` 
type StructName struct {
	Items []struct {
		PaymentType string `json:"payment_type"`
		SenderName string `json:"sender_name"`
		CurrencyCode string `json:"currency_code"`
		TotalQtyOrdered float64 `json:"total_qty_ordered"`
		UpdatedAt string `json:"updated_at"`
		PaidAt string `json:"paid_at"`
		OrderItems []struct {
			ItemDiscountAmount string `json:"item_discount_amount"`
			Sku string `json:"sku"`
			Weight string `json:"weight"`
			Price string `json:"price"`
			QtyOrdered string `json:"qty_ordered"`
			Name string `json:"name"`
		} `json:"order_items"`
		OrderSource string `json:"order_source"`
		TotalItemCount string `json:"total_item_count"`
		DiscountAmount string `json:"discount_amount"`
		Subtotal string `json:"subtotal"`
		GrandTotal string `json:"grand_total"`
		OrderNumber string `json:"order_number"`
		OrderStatus string `json:"order_status"`
		TotalWeight int `json:"total_weight"`
		ShippingAmount string `json:"shipping_amount"`
		CreatedAt string `json:"created_at"`
		ShippingAddressInfo struct {
			Email string `json:"email"`
			Telephone string `json:"telephone"`
			County string `json:"county"`
			Postcode string `json:"postcode"`
			OrderItems []struct {
				Name string `json:"name"`
				ItemDiscountAmount string `json:"item_discount_amount"`
				Sku string `json:"sku"`
				Weight string `json:"weight"`
				Price string `json:"price"`
				QtyOrdered string `json:"qty_ordered"`
			} `json:"order_items"`
			Name string `json:"name"`
			Province string `json:"province"`
			City string `json:"city"`
			Street string `json:"street"`
			IDCard string `json:"id_card"`
		} `json:"shipping_address_info"`
	} `json:"items"`
}
``` 
和
``` 
type StructName struct {
	Items []*Items `json:"items"`
}
type Items struct {
	PaymentType         string               `json:"payment_type"`
	SenderName          string               `json:"sender_name"`
	CurrencyCode        string               `json:"currency_code"`
	TotalQtyOrdered     float64              `json:"total_qty_ordered"`
	UpdatedAt           string               `json:"updated_at"`
	PaidAt              string               `json:"paid_at"`
	OrderItems          []*OrderItems        `json:"order_items"`
	OrderSource         string               `json:"order_source"`
	TotalItemCount      string               `json:"total_item_count"`
	DiscountAmount      string               `json:"discount_amount"`
	Subtotal            string               `json:"subtotal"`
	GrandTotal          string               `json:"grand_total"`
	OrderNumber         string               `json:"order_number"`
	OrderStatus         string               `json:"order_status"`
	TotalWeight         int                  `json:"total_weight"`
	ShippingAmount      string               `json:"shipping_amount"`
	CreatedAt           string               `json:"created_at"`
	ShippingAddressInfo *ShippingAddressInfo `json:"shipping_address_info"`
}
type OrderItems struct {
	ItemDiscountAmount string `json:"item_discount_amount"`
	Sku                string `json:"sku"`
	Weight             string `json:"weight"`
	Price              string `json:"price"`
	QtyOrdered         string `json:"qty_ordered"`
	Name               string `json:"name"`
}
type ShippingAddressInfo struct {
	Email                         string                           `json:"email"`
	Telephone                     string                           `json:"telephone"`
	County                        string                           `json:"county"`
	Postcode                      string                           `json:"postcode"`
	OrderItemsShippingAddressInfo []*OrderItemsShippingAddressInfo `json:"order_items_shipping_address_info"`
	Name                          string                           `json:"name"`
	Province                      string                           `json:"province"`
	City                          string                           `json:"city"`
	Street                        string                           `json:"street"`
	IDCard                        string                           `json:"id_card"`
}
type OrderItemsShippingAddressInfo struct {
	Name               string `json:"name"`
	ItemDiscountAmount string `json:"item_discount_amount"`
	Sku                string `json:"sku"`
	Weight             string `json:"weight"`
	Price              string `json:"price"`
	QtyOrdered         string `json:"qty_ordered"`
}
``` 
