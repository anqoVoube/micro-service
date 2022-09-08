package main

import (
	"authentication/data"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	jsonData := jsonUserData{false, Data{user.ID, user.FirstName, body.Password}}
	c.JSON(http.StatusOK, jsonData)
}
