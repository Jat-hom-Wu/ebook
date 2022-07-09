package service

import (
	"github.com/dgrijalva/jwt-go"
)

const SingingKey string = "ebook"

type User struct{
	Name string `json:"name"`
	PassWord string `json:"password"`
}

type UserClaims struct{
	User
	jwt.StandardClaims
}

func GenerateToken(method jwt.SigningMethod, claims jwt.Claims, mySignKey interface{}) (string,error){
	token := jwt.NewWithClaims(method, claims)
	ss,err := token.SignedString(mySignKey)	//签名
	if err != nil{
		return ss,err
	}
	return ss,nil
}

func ParseToken(tokenString string, claims jwt.Claims, mySignKey interface{}) (*jwt.Token,error){
	token,err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
		return mySignKey,nil
	})
	return token,err
}

