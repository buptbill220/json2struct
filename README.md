![](j2s.png)

`json2struct` generates strongly-typed models and serializers from JSON, making it a breeze to work with JSON type-safely in any other languages.

### Supported Inputs

| JSON | JSON FILE |
| ---- | ------------- |


| TypeScript | GraphQL queries |
| ---------- | --------------- |


### Target Models

| [Golang]() | [Protobuf]() |
| ---------------------------------------- | -------------------------------------------- |

### Features
* written by golang
* comparing with ![json-to-go](https://github.com/mholt/json-to-go), it can distinguish golang `map` model.
```javascript
json data for example:
{
        "23": 343434,
        "3232": 23
}
the result translated by `json-to-go`
type AutoGenerated struct {
        Num23   int `json:"23"`
        Num3232 int `json:"3232"`
}

we expect result in follow
type AutoGenerated map[int]int
```
* comparing with ![quicktype](https://github.com/quicktype/quicktype), it identifys int/int64/float model more precise. Apart from this, `json2struct` generates
single type name for embedding struct model if there are 2 or more json input.
```
json data
{
  "aid": 6858121093648023822,
  "bid": 6858073458614731015,
  "show_more": true,
  "coupon": {
    "6858121093648023822": {
      "max_coupon": {
        "credit": 8
      },
      "total_stock": 2
    }
  }
}
`quicktype` result
type Welcome struct {
        AID      *float64        `json:"aid,omitempty"`
        BID      *float64        `json:"bid,omitempty"`
        ShowMore *bool     `json:"show_more,omitempty"`
        MaxCouponDefault   map[string] Coupon `json:"coupon,omitempty"`
}
type Coupon struct {
        MaxCoupon  *MaxCoupon `json:"max_coupon,omitempty"`
        TotalStock *int64     `json:"total_stock,omitempty"`
}
type MaxCoupon struct {
        Credit *int64 `json:"credit,omitempty"`
}

we expect
type Welcome struct {
        ShareComboID        *int64                    `json:"share_combo_id,omitempty"`
        AddressComboID      *int64                    `json:"address_combo_id,omitempty"`
        ShowMoreTalentAreas *bool                       `json:"show_more_talent_areas,omitempty"`
        ShowMoreBrandAreas  *bool                       `json:"show_more_brand_areas,omitempty"`
        MaxCouponDefault    map[int64]MaxCouponDefault `json:"max_coupon_default,omitempty"`
}
for more, if we translate the following json to the same go file, the model `Coupon、MaxCoupon` will be double definition.
{
    "coupon": {
    "6858121093648023822": {
      "max_coupon": {
        "credit": 818
      },
      "total_stock": 2000
    }
}
```
* make it protobuf defination

### Installation

```
go get github.com/buptbill220/json2struct/json2struct
```

### Using `json2struct`
* for sdk lib
```
import "github.com/buptbill220/json2struct"

ret, err := Json2Struct("auto_name", `[1,2,3,44444444444]`, GoStructType)
fmt.Printf("ret:\n %s, %v\n", ret, err)
```
* for installed cmd
json2struct gen --json `[1,2,3]` [--file a.json] [--type go]