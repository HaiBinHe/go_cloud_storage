package cache

import (
	"fmt"
	"go-cloud/conf"
	"go-cloud/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	Db *gorm.DB
)

func InitMySQLConn() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.DataBaseSetting.UserName,
		conf.DataBaseSetting.Password,
		conf.DataBaseSetting.Host,
		conf.DataBaseSetting.DBName,
		conf.DataBaseSetting.Charset,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "",
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDb, err := Db.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(conf.DataBaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(conf.DataBaseSetting.MaxOpenConns)
	err = sqlDb.Ping()
	if err != nil {
		log.Println("mysql ping err:", err)
		return err
	}
	log.Println("mysql server start")
	err = autoMigrate()
	if err != nil {
		log.Println("gorm autoMigrate err: ", err)
	}
	log.Println("AutoMigrate Success")
	return nil
}
func autoMigrate() error {
	return Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		&model.User{},
		&model.UserFile{},
		&model.FileStore{},
		&model.FileFolder{},
		&model.FileShare{},
		&model.FileCenter{},
	)
}
