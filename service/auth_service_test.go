package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)


func TestGenerateToken(t *testing.T){
	mySigningKey := []byte("AllYourBase")
	claims := UserClaims{
		User{"xiaoming","123"},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
			Issuer:    "test",
		},
	}
	token,err := GenerateToken(jwt.SigningMethodHS256, claims, mySigningKey)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("token:",token)
	fmt.Println("length:", len(token))
	raw,err := ParseToken(token, &UserClaims{}, mySigningKey)
	if err != nil{
		t.Error("parse token failed,",err)
	}
	c,ok := raw.Claims.(*UserClaims)
	if ok && raw.Valid{
		fmt.Println(*c)
		fmt.Println(c.Name)
	}else{
		t.Error("token invalid")
	}
}