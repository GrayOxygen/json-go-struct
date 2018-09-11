# json-go-struct
convert json to go struct(one nest struct and multi normal separate structs),you can check my another project https://github.com/GrayOxygen/json-go-struct-app/ just a app version of this project
>
1. 将json转为golang的struct，目前开源的项目、工具中，只找到转成嵌套的struct形式，于是写了个转成非嵌套的struct
2. 思路：用已有库生成内嵌struct，再构建树，从顶部到底部依次入栈(左大括号)出栈(右大括号)即得到树模型，只需注意每层是否有struct或者struct数组的类型的儿子节点(不看孙子及以下节点)，替换为*某struct和[]*某struct这种形式即可，我的替换方式是，对于某层，将下边的所有struct/struct数组的类型用“树节点ID”替换，最后输出时，再替换即可，替换时就处理了同名的情况，用父节点名字+当前名字

## Example
input
## JSON 
```
{
  "items": [
    {
      "order_number": "614318762004012957",
      "order_source": "po",
      "sender_name": "jd",
      "order_status": "processing",
      "currency_code": "CNY",
      "total_qty_ordered": 1.0000,
      "total_item_count": "1",
      "total_weight": 100,
      "discount_amount": "0.00",
      "shipping_amount": "0.00",
      "subtotal": "79.9900",
      "grand_total": "79.9900",
      "created_at": "2015-12-11 22:51:53",
      "updated_at": "2015-12-22 20:14:18",
      "paid_at": "2015-12-22 20:14:18",
      "payment_type": "alipay_payment",
      "shipping_address_info": {
        "email": "416757228@qq.com",
        "name": "常璐",
        "telephone": "18687079066",
        "province": "北京市",
        "city": "北京市",
        "county": "海淀区",
        "street": "西土城路25号中国政法大学研究生院",
        "postcode": "100088",
        "id_card": "431102199603656899",
        "order_items": [
          {
            "sku": "LANCOSC73978802",
            "weight": "100.0000",
            "price": "79.9900",
            "qty_ordered": "1.0000",
            "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
            "item_discount_amount": "0.00"
          }
        ]
      },
      "order_items": [
        {
          "sku": "LANCOSC73978802",
          "weight": "100.0000",
          "price": "79.9900",
          "qty_ordered": "1",
          "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
          "item_discount_amount": "0.00"
        }
      ]
    },
    {
      "order_number": "614318762004012951",
      "order_source": "po",
      "sender_name": "jd",
      "order_status": "processing",
      "currency_code": "CNY",
      "total_qty_ordered": 1,
      "total_item_count": "1",
      "total_weight": 100,
      "discount_amount": "0.00",
      "shipping_amount": "0.00",
      "subtotal": "79.9900",
      "grand_total": "79.9900",
      "created_at": "2015-12-11 22:51:53",
      "updated_at": "2015-12-22 20:14:18",
      "paid_at": "2015-12-22 20:14:18",
      "payment_type": "vt_payment",
      "shipping_address_info": {
        "email": "416757228@qq.com",
        "name": "常璐",
        "telephone": "18687079066",
        "province": "北京市",
        "city": "北京市",
        "county": "海淀区",
        "street": "西土城路25号中国政法大学研究生院",
        "postcode": "100088",
        "id_card": "431102199603656899"
      },
      "order_items": [
        {
          "sku": "LANCOSC73978802",
          "weight": "100.0000",
          "price": "79.9900",
          "qty_ordered": "1.0000",
          "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
          "item_discount_amount": "0.00"
        }
      ]
    }
  ]
}
```
output
## Multi Separate Struct
```
type StructName struct {
	Items []*Items `json:"items"`
}
type Items struct {
	TotalItemCount      string               `json:"total_item_count"`
	GrandTotal          string               `json:"grand_total"`
	CreatedAt           string               `json:"created_at"`
	PaidAt              string               `json:"paid_at"`
	PaymentType         string               `json:"payment_type"`
	ShippingAddressInfo *ShippingAddressInfo `json:"shipping_address_info"`
	ItemsOrderItems []*ItemsOrderItems `json:"items_order_items"`
	OrderSource     string             `json:"order_source"`
	TotalWeight     int                `json:"total_weight"`
	Subtotal        string             `json:"subtotal"`
	SenderName      string             `json:"sender_name"`
	OrderStatus     string             `json:"order_status"`
	UpdatedAt       string             `json:"updated_at"`
	DiscountAmount  string             `json:"discount_amount"`
	ShippingAmount  string             `json:"shipping_amount"`
	OrderNumber     string             `json:"order_number"`
	CurrencyCode    string             `json:"currency_code"`
	TotalQtyOrdered float64            `json:"total_qty_ordered"`
}
type ShippingAddressInfo struct {
	City       string        `json:"city"`
	County     string        `json:"county"`
	IDCard     string        `json:"id_card"`
	OrderItems []*OrderItems `json:"order_items"`
	Email      string        `json:"email"`
	Name       string        `json:"name"`
	Telephone  string        `json:"telephone"`
	Province   string        `json:"province"`
	Street     string        `json:"street"`
	Postcode   string        `json:"postcode"`
}
type OrderItems struct {
	QtyOrdered         string `json:"qty_ordered"`
	Name               string `json:"name"`
	ItemDiscountAmount string `json:"item_discount_amount"`
	Sku                string `json:"sku"`
	Weight             string `json:"weight"`
	Price              string `json:"price"`
}
type ItemsOrderItems struct {
	Price              string `json:"price"`
	QtyOrdered         string `json:"qty_ordered"`
	Name               string `json:"name"`
	ItemDiscountAmount string `json:"item_discount_amount"`
	Sku                string `json:"sku"`
	Weight             string `json:"weight"`
}

```
and
## Nested Struct
```
type StructName struct {
	Items []struct {
		TotalItemCount string `json:"total_item_count"`
		GrandTotal string `json:"grand_total"`
		CreatedAt string `json:"created_at"`
		PaidAt string `json:"paid_at"`
		PaymentType string `json:"payment_type"`
		ShippingAddressInfo struct {
			City string `json:"city"`
			County string `json:"county"`
			IDCard string `json:"id_card"`
			OrderItems []struct {
				QtyOrdered string `json:"qty_ordered"`
				Name string `json:"name"`
				ItemDiscountAmount string `json:"item_discount_amount"`
				Sku string `json:"sku"`
				Weight string `json:"weight"`
				Price string `json:"price"`
			} `json:"order_items"`
			Email string `json:"email"`
			Name string `json:"name"`
			Telephone string `json:"telephone"`
			Province string `json:"province"`
			Street string `json:"street"`
			Postcode string `json:"postcode"`
		} `json:"shipping_address_info"`
		OrderItems []struct {
			Price string `json:"price"`
			QtyOrdered string `json:"qty_ordered"`
			Name string `json:"name"`
			ItemDiscountAmount string `json:"item_discount_amount"`
			Sku string `json:"sku"`
			Weight string `json:"weight"`
		} `json:"order_items"`
		OrderSource string `json:"order_source"`
		TotalWeight int `json:"total_weight"`
		Subtotal string `json:"subtotal"`
		SenderName string `json:"sender_name"`
		OrderStatus string `json:"order_status"`
		UpdatedAt string `json:"updated_at"`
		DiscountAmount string `json:"discount_amount"`
		ShippingAmount string `json:"shipping_amount"`
		OrderNumber string `json:"order_number"`
		CurrencyCode string `json:"currency_code"`
		TotalQtyOrdered float64 `json:"total_qty_ordered"`
	} `json:"items"`
}
```


> I used https://github.com/mholt/json-to-go to get nested golang struct , nice project ,  much appreciated!!!
