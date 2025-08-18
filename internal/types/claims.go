package types

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID         float64
	Username   string
	Experation float64
	IAT        float64
}

func CreateClaims(claims jwt.MapClaims) (*Claims, error) {
	newClaims := Claims{}
	var ok bool

	if newClaims.ID, ok = claims["id"].(float64); !ok {
		return nil, fmt.Errorf("error for id parse")
	}
	if newClaims.Username, ok = claims["username"].(string); !ok {
		return nil, fmt.Errorf("error for username parse")
	}
	if newClaims.Experation, ok = claims["exp"].(float64); !ok {
		return nil, fmt.Errorf("error for experation parse")
	}
	if newClaims.IAT, ok = claims["iat"].(float64); !ok {
		return nil, fmt.Errorf("error for iat parse")
	}

	return &newClaims, nil
}
