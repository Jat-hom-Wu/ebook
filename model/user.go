package model

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RDB *redis.Client	//unused
	RDB2 *redis.Client 	//unused
	RCUser *RedisClient		//store user info
	RCToken *RedisClient	//store token
)

type RedisClient struct{
	db *redis.Client
}

func NewRedisClient(rdb *redis.Client) *RedisClient{
	return &RedisClient{db:rdb}
}


type UserTable struct{
	Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func AddUser(name,password string) error{
	user := UserTable{name,password}
	return DB.Create(&user).Error
}

func UpdateUser(name,password string) error{
	result := DB.Model(&UserTable{}).Where("name = ?", name).Update("password", password)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func DeleteUser(){

}

func FindUser(name string)(UserTable, error){
	user := UserTable{}
	result := DB.Find(&user, "name = ?", name)
	if result.Error != nil{
		return user,result.Error
	}
	return user,nil
}

func (r *RedisClient)RedisFindUser(name string) (string,error){
	ctx := context.Background()
	val,err := r.db.Get(ctx, name).Result()
	if err != nil{
		if err == redis.Nil{
			return "",nil
		}else{
			return "",err
		}
	}else{
		return val,nil
	}
}

func (r *RedisClient)RedisUpdateUser(name,password string,time time.Duration) error{
	ctx := context.Background()
	val := r.db.Set(ctx, name, password, time)
	
	if val.Err() != nil{
		return val.Err()
	}
	return nil
}