package controllers

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/web3santa/JWT-Authentication-Golang/database"
	"github.com/web3santa/JWT-Authentication-Golang/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user model.User
	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	// Get the expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Construct a *NumericDate from the expiration time
	expiresAt := jwt.NewNumericDate(expirationTime)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: expiresAt,
	})

	godotenv.Load()
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := claims.SignedString([]byte(SECRET_KEY))

	if err != nil {
		// JWT 토큰 생성에 실패한 경우
		c.Status(fiber.StatusInternalServerError) // 500 상태 코드 반환
		return c.JSON(fiber.Map{
			"message": "could not login", // 오류 메시지를 JSON 응답에 포함
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// JWT 토큰 생성에 성공한 경우
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	godotenv.Load()
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized) // 500 상태 코드 반환
		return c.JSON(fiber.Map{
			"message": "unauthenticated", // 오류 메시지를 JSON 응답에 포함
		})
	}

	claims := token.Claims

	str, err := claims.GetIssuer()
	if err != nil {
		c.Status(fiber.StatusBadRequest) // 500 상태 코드 반환
		return c.JSON(fiber.Map{
			"message": "GetIssuer error", // 오류 메시지를 JSON 응답에 포함
		})
	}

	var user model.User
	database.DB.Where("id = ?", str).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	log.Println("Logout!")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
