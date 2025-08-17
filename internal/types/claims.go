package types

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID         int
	Username   string
	Experation time.Time
	IAT        time.Time
}

func CreateClaims(claims jwt.MapClaims) (*Claims, error) {
	newClaims := Claims{}
	var ok bool

	if newClaims.ID, ok = claims["id"].(int); !ok {
		return nil, fmt.Errorf("error for id parse")
	}
	if newClaims.Username, ok = claims["username"].(string); !ok {
		return nil, fmt.Errorf("error for username parse")
	}
	if newClaims.Experation, ok = claims["exp"].(time.Time); !ok {
		return nil, fmt.Errorf("error for experation parse")
	}
	if newClaims.IAT, ok = claims["iat"].(time.Time); !ok {
		return nil, fmt.Errorf("error for iat parse")
	}

	return &newClaims, nil
}
