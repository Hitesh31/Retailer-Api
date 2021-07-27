package Models
type Product struct {
	Id   string   `json:"id"`
	ProductName string `json:"product_name"`
	Quantity    uint `json:"quantity"`
	Price       uint `json:"price"`
}
type OrderTrack struct {
	CustomerId string `json:"customer_id",gorm:"foreignKey"`
	ProductId string `json:"product_id"`
	Id string `json:"id"`
	Quantity uint `json:"quantity"`
	Bill uint `json:"bill"`
}

type Customer struct {
	CustomerName string `json:"customer_name"`
	Id string `json:"id"`
	Phone string `json:"phone"`
}
func (b *Product) TableName() string {
	return "product"
}
func (b *OrderTrack) TableName() string {
	return "order_track"
}
func (b *Customer) TableName() string {
	return "customer"
}
