package repositories

import (
	"gorm.io/gorm"
	"tracking-service-go/internal/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(order models.Order) (models.Order, error) {
	result := r.db.Create(&order)
	return order, result.Error
}

func (r *OrderRepository) FindAllOrders(id int) ([]models.Order, error) {
	var orders []models.Order
	result := r.db.Where("user_id = ?", id).Find(&orders)
	return orders, result.Error
}