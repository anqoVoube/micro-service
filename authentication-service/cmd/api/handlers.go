package main

import (
	"authentication/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type jsonUserData struct {
	Error bool `json:"error"`
	Data  Data `json:"data"`
}

type Data struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
}

func ping(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		err = c.AbortWithError(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	userStruct := data.User{}
	user, err := userStruct.GetByEmail(body.Email)
	if err != nil {
		err = c.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	isAuthenticated, err := user.PasswordMatches(body.Password)
	if isAuthenticated == false {
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, "Passwords didn't match")
			return
		}
	}
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	jsonData := Data{user.ID, user.FirstName, body.Password}
	instance, err := generateTokenPair(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Login/Password is not valid."})
	}
	c.JSON(http.StatusOK, instance)
}

func generateTokenPair(data *Data) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = data.Id
	claims["name"] = data.FirstName
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["sub"] = data.Id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
