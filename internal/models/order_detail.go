package models

type OrderDetail struct {
	Order
	Items []OrderItemList `json:"items"`
}
