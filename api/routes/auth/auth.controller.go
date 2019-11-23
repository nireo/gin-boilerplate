package auth

import (
	"io/ioutil"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nireo/gin-boilerplate/database/models"
	"github.com/nireo/gin-boilerplate/lib/common"
	"golang.org/x/crypto/bcrypt"
)

// User alias for model
type User = models.User

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(data common.JSON) (string, error) {
	// token is valid for 7 days
	date := time.Now().Add(time.Hour * 24 * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp": data.Unix(),
	})

	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key.pub"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func login(c *gin.Context) {
	// get gin from context
	db := c.MustGet("db").(*gorm.DB)

	// define request body interface for validation
	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		// if the body is invalid
		c.AbortWithStatus(400)
		return
	}

	// check if the user even exists
	var user User
	if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
		// no user found
		c.AbortWithStatus(404)
		return
	}

	if !checkHash(body.Password, user.PasswordHash) {
		// invalid credentials
		c.AbortWithStatus(401)
		return
	}

	serialized := user.Serialize()
	token, err := generateToken(serialized)

	if err != nil {
		// something went wrong with token creation
		// return a internal server error
		c.AbortWithStatus(500)
		return
	}

	c.SetCookie("token", token, 60*60*26*7, "/", "", false, true)
	c.JSON(200, common.JSON{
		"user":  user.Serialize(),
		"token": token,
	})
}