package utils

import (
	"fmt"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtAccessKey = []byte("my_secret_key")  // change later
var jwtRefreshKey = []byte("my_secret_key") // change later
var accessTokenExperation = 1               // time in hours for expiration of token
var refreshTokenExperation = 30 * 24        // time in hours for expiration refresh token

func GenerateJWTAccessToken(id int, username string) (string, error) {

	claims := jwt.MapClaims{
		"id":       int(id),
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(accessTokenExperation)).Unix(), //
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtAccessKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyAccessToken(tokenString string) (models.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtRefreshKey), nil
	})

	if err != nil {
		return models.Claims{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return models.Claims{}, fmt.Errorf("invalid token")
	}

	return *claims, nil

}

func VerifyRefreshToken(tokenString string) (models.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtRefreshKey), nil
	})

	if err != nil {
		return models.Claims{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return models.Claims{}, fmt.Errorf("invalid token")
	}

	return *claims, nil
}

func GenerateRefreshToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       int(id),
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(refreshTokenExperation)).Unix(), //
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtRefreshKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
