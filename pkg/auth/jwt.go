package auth

import (
	"github.com/dgrijalva/jwt-go"
	"go-cloud/conf"
	"go-cloud/internal/model"
	"log"
	"time"
)

type Claim struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(conf.JWTSetting.Secret)
}
func GenerateToken(user model.User, Exp time.Time) (string, error) {
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim{
		Id:   user.ID,
		Name: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Exp.Unix(),
			Issuer:    conf.JWTSetting.Issuer,
		},
	})
	signed := GetJWTSecret()
	token, err := tokenClaim.SignedString(signed)
	if err != nil {
		log.Println("generate token failed:", err)
		return "", err
	}
	return token, nil
}
func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid：验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），
		//需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的。
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
