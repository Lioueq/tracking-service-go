package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
	"tracking-service-go/internal/models"
	"tracking-service-go/internal/repositories"
)

var userRepo *repositories.UserRepository
var jwtSecret string

func InitUserRepository(db *gorm.DB, secret string) {
	userRepo = repositories.NewUserRepository(db)
	jwtSecret = secret
}

// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegister true "New User"
// @Success 201 {object} models.User
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/register [post]
func Register(c echo.Context) error {
	var userCreds models.UserRegister
	if err := c.Bind(&userCreds); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	var user models.User
	user.Name = userCreds.Name
	user.Email = userCreds.Email
	user.Password = userCreds.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Couldn't hash password"})
	}
	user.Password = string(hashedPassword)

	createdUser, err := userRepo.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create user"})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.Credentials true "User Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/login [post]
func Login(c echo.Context) error {
	var credentials models.Credentials
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}

	user, err := userRepo.FindByEmail(credentials.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid password"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"expiredAt": time.Now().Add(time.Hour * 72),
	})

	token, err := claims.SignedString([]byte(jwtSecret))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// @Summary Get all users
// @Description Get details of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.Error
// @Router /users [get]
func GetUsers(c echo.Context) error {
	users, err := userRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}
