package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

type User struct {
	UserId   int `json:"id"`
	UserName string `json:"user_name"`
	Phone	int	`json:"phone"`
}

type CustomerClaim struct {
	User
	*jwt.StandardClaims
}

func GetToken(r *http.Request) string {
	tokens := r.Header.Get("Authorization")

	token := strings.Split(tokens, " ")
	fmt.Println(tokens)
	if len(token)<=1 {
		return ""
	}
	return token[1]
}

func (user *User) GetParams(userID int,userName string,phone int) {
	user.UserId=userID
	user.UserName=userName
	user.Phone=phone
}

func Encode(user User) (string, error) {
	//privateKey, _ := base64.URLEncoding.DecodeString(key)

	//设置超时时间
	expTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	fmt.Println("jwt编码的数据user：", user)
	//设置Claim
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: expTime}}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString([]byte("secret"))

}

//解析jwt
func Decode(tokenString string) (*CustomerClaim, error) {
	//tokenString:=GetToken(ctx.Request)
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println("token解码出错: ", err, "接收到的token为：", tokenString)
		return nil, err
	}
	if user, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return user, nil
	} else {
		return nil, err
	}
}

