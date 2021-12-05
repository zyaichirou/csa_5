//@Title		jwt.go
//@Description	jwt鉴权
//@Author		zy
//@Update		2021.12.5

package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 结构体
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}



// JWT过期时间
const TokenExpireDuration = time.Hour * 2

// Secret签名
var MySecret = []byte("zyaichirou")

//GenToken
//@title		GenToken()
//@description	生成JWT
//@author		zy
//@param		username string
//@return		string error
func GenToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),		//过期时间
			Issuer: "zy",												//签发人
			Subject: "userinformation token",							//签发主题
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名
	tokenString, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

//ParseToken
//@title		ParseToken()
//@description	解析JWT
//@author		zy
//@param		tokenString string
//@return		*MyClaims error
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}	//校验token
	return nil, errors.New("invalid token")
}
