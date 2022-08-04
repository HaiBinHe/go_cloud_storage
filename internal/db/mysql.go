package db

import (
	"fmt"
	"go-cloud/cmd"
	"go-cloud/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	Db *gorm.DB
)

func InitMySQLConn() error{
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
	cmd.DataBaseSetting.UserName,
	cmd.DataBaseSetting.Password,
	cmd.DataBaseSetting.Host,
	cmd.DataBaseSetting.DBName,
	cmd.DataBaseSetting.Charset,
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

	sqlDb.SetMaxIdleConns(cmd.DataBaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cmd.DataBaseSetting.MaxOpenConns)
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
func autoMigrate() error{
	return Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		&model.User{},
		&model.UserInfo{},
		&model.UserFile{},
		&model.FileStore{},
		&model.FileFolder{},
		&model.FileShare{},
		)
}
