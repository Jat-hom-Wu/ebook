package dao

import (
	"fmt"

	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() error {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.MysqlUser,
		setting.MysqlPassword,
		setting.Addr,
		setting.MysqlPort,
		setting.MysqlDB)
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.Error("open mysql failed,", err)
		return err
	}
	model.DB = db
	return nil
}

//数据淘汰：设置redis,通过lru淘汰数据
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddr,
		Password: "",
		DB:       setting.RedisDB,
	})
	model.RDB = rdb

	rdb2 := redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddr,
		Password: "",
		DB:       setting.RedisDB2,
	})
	model.RDB2 = rdb2

	model.RCUser = model.NewRedisClient(model.RDB)
	model.RCToken = model.NewRedisClient(model.RDB2)
}


