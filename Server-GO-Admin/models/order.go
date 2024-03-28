package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id         uint        `json:"id"`
	Firstanme  string      `json:"-"`
	Lastname   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Total      float64     `json:"total" gorm:"-"`
	Email      string      `json:"email"`
	OrderItems []OrderItem `json:"order_time" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float64 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {

	var total int64

	db.Model(order).Count(&total)

	return total

}
func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {

	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {

		var total float64 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float64(orderItem.Quantity)
		}
		orders[i].Name = orders[i].Firstanme + " " + orders[i].Lastname
		orders[i].Total = total

	}

	return orders

}
