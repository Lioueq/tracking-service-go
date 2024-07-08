package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"tracking-service-go/internal/models"
	"tracking-service-go/internal/repositories"
)

var orderRepo *repositories.OrderRepository

func InitOrderRepository(db *gorm.DB) {
	orderRepo = repositories.NewOrderRepository(db)
}

// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} []models.Order
// @Failure 500 {object} models.Error
// @Security BearerAuth
// @Router /orders [get]
func GetOrders(c echo.Context) error {
	user := c.Get("id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	orders, err := orderRepo.FindAllOrders(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch orders"})
	}
	return c.JSON(http.StatusOK, orders)
}

// @Summary Create order
// @Description Create order
// @Tags orders
// @Produce json
// @Param order body models.OrderCreate true "New Order"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Security BearerAuth
// @Router /orders/create [post]
func CreateOrder(c echo.Context) error {
	var orderData models.OrderCreate
	if err := c.Bind(&orderData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	user := c.Get("id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var order models.Order
	order.Name = orderData.Name
	order.Status = "created"
	order.UserId = int(claims["id"].(float64))
	createdOrder, err := orderRepo.Create(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create order"})
	}
	return c.JSON(http.StatusCreated, createdOrder)
}

// @Summary Update order
// @Description Update order
// @Tags orders
// @Produce json
// @Param order body models.OrderUpdate true "Update Order"
// @Success 200 {object} models.Error
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Security BearerAuth
// @Router /orders/update [put]
func UpdateOrder(c echo.Context) error {
	var orderData models.OrderUpdate
	if err := c.Bind(&orderData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	user := c.Get("id").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	order, err := orderRepo.FindById(orderData.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Order not found"})
	}
	if order.UserId != int(claims["id"].(float64)) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Not authorized"})
	}
	order.Status = orderData.Status
	err = orderRepo.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update order"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Ok"})
}
