package main

import (
	"log"

	"github.com/ali-alhashim/todolist-mysql-go-2/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func DBMigrations() {
	
	
	
	log.Println("Database init start ...")

	// register model
	orm.RegisterModel(new(models.Task))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Akte@12345678@/todolist?charset=utf8", 30)
	
	// create table
	orm.RunSyncdb("default", false, true)

	
   

	
}	