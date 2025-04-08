package controllers

import (
	"github.com/arbeeorlar/jwt-example/initializers"
	"github.com/arbeeorlar/jwt-example/models"
	"github.com/arbeeorlar/jwt-example/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func SignUpController(c *gin.Context) {

	var user models.JWTUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.JWTUser
	initializers.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GeneratePasswordhash(user.Password)
	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errHash.Error()})
		return
	}
	initializers.DB.Create(&user)
	response := models.Response{
		Status:  "success",
		Message: "User Created",
		Data:    user,
	}
	c.JSON(http.StatusOK, response)
}

func HomeController(c *gin.Context) {
	var header = c.GetHeader("Authorization")
	if header == "" {
		log.Println("Authorization Header is empty")
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	parts := strings.Split(header, "Bearer ")
	if len(parts) != 2 {
		log.Println("Authorization Header is invalid")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	tokenString := strings.TrimSpace(parts[1]) // Clean whitespace

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if claims.User.ID == 0 {
		log.Println("User ID is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	JWTUser := claims.User
	existingUser := models.JWTUser{}
	initializers.DB.Where("email = ?", JWTUser.Email).First(&existingUser)
	if existingUser.ID == 0 {
		log.Println("User ID is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	expiresAt := time.Unix(claims.ExpiresAt, 0)
	if time.Now().After(expiresAt) {
		log.Println("Token has expired")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	if claims.Role != "user" && claims.Role != "manager" {
		log.Println("Role is invalid")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	var users []models.JWTUser
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	response := models.Response{
		Status:  "success",
		Message: "User Found",
		Data:    users,
	}
	c.JSON(http.StatusOK, response)

}

func PremiumController(c *gin.Context) {

}
func SignOutController(c *gin.Context) {

}

func LoginController(c *gin.Context) {
	var User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password"  binding:"required"`
	}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.JWTUser
	initializers.DB.Where("email = ?", User.Email).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var errHash error
	errHash = utils.ComparePassword(existingUser.Password, User.Password)
	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}
	existingUser.Password = ""
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &models.Claims{
		User: existingUser,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not sign JWT"})
		return
	}
	response := models.Response{
		Status:  "success",
		Message: "Ok",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}
