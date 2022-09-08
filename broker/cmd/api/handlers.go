package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]int{"hello": 1})
}

func pong(c *gin.Context) {
	var requestPayload RequestPayload

	err := c.ShouldBindJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unknown action 29")
		return
	}
	switch requestPayload.Action {
	case "auth":
		authenticate(c, AuthPayload{Email: "admin@example.com", Password: "verysecret!"})
	default:
		c.JSON(http.StatusBadRequest, "Unknown action 36")
		return
	}
}

type jsonUserData struct {
	Error bool `json:"error"`
	Data  Data `json:"data"`
}

type Data struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
}

func authenticate(c *gin.Context, payload AuthPayload) {
	jsonData, _ := json.MarshalIndent(payload, "", "\t")

	request, err := http.NewRequest(
		"POST", "http://authentication-service/authenticate",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unknown action 60")
		return
	}
	// somecomment
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Line 67: %s", err.Error()))
		return
	}

	defer response.Body.Close()

	var jsonDataOfUser jsonUserData

	err = json.NewDecoder(response.Body).Decode(&jsonDataOfUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Line 89: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, jsonDataOfUser)

}
