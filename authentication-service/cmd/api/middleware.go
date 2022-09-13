package main

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	authorizationHeader = "Authorization"
)

//type ErrorJSonResponse struct {
//	Message string `json:"message"`
//}

//func userIdentity(c *gin.Context) {
//	header := c.GetHeader(authorizationHeader)
//	if header == "" {
//		c.JSON(401, ErrorJSonResponse{
//			"empty auth header",
//		})
//		return
//	}
//	headerParts := string.Split(header, " ")
//	if len(headerParts) != 2 {
//		c.JSON(401, ErrorJSonResponse{"invalid auth header"})
//		return
//	}
//
//}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

//func ParseToken(token string) (int, error) {
//	token, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token, *jwt.Token) (interface{}, error){
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
//			return nil, errors.New("invalid signed method")
//		}
//
//		return []byte(signingKey), nil
//
//	}
//}
