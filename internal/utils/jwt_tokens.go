package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtAccessKey = []byte("my_secret_key")  // change later
var jwtRefreshKey = []byte("my_secret_key") // change later
var accessTokenExperation = 1               // time in hours for expiration of token
var refreshTokenExperation = 30 * 24        // time in hours for expiration refresh token

func GenerateJWTAccessToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(accessTokenExperation)).Unix(), //
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtAccessKey)

	fmt.Println("err", err)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, accessKeyFunction)

	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return jwt.MapClaims{}, fmt.Errorf("invalid token")

}

func VerifyRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, refreshKeyFunction)

	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return jwt.MapClaims{}, nil
}

func accessKeyFunction(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	}
	return jwtAccessKey, nil
}

func refreshKeyFunction(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	}
	return jwtRefreshKey, nil
}

func GenerateRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
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
