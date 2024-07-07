package controllers

import (
	"fmt"
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
		fmt.Println(err)
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
