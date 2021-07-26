package Models
type Product struct {
	ProductId          uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}
type OrderTrack struct {
	CustomerId uint `json:"customer_id"`
	ProductId uint `json:"product_id"`
	Quantity uint `json:"quantity"`
	Bill uint `json:"bill"`
}
func (b *Product) TableName() string {
	return "product"
}
func (b *OrderTrack) TableName() string {
	return "order_track"
}
