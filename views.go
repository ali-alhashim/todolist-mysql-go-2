package main

import (
	"log"

	"github.com/ali-alhashim/todolist-mysql-go-2/models"
	"github.com/astaxie/beego/orm"
	"github.com/gofiber/fiber/v2"
)



func Home(c *fiber.Ctx) error {
	
		 return c.Render("index", fiber.Map{
			                                  "Title": "Home Page | All Tasks",
		                                   })
	
}

func AddTask(c *fiber.Ctx) error {
	
	return c.Render("addTask", fiber.Map{
		"Title": "Add New Task Page",
	 })

}

func PostTask(c *fiber.Ctx) error{
    //Save the new task

	//object of struct 
	newTask := new(models.Task)

	if err := c.BodyParser(newTask); err != nil {
		log.Println("Error in body parser")
		return err
	}

    log.Println("Your Task ID is :" , newTask.Id)
	log.Println("Your Task subject is :" , newTask.Subject)
	log.Println("Your Task Status :" , newTask.IsCompleted)

	dbORM := orm.NewOrm()
    
	id, err := dbORM.Insert(&newTask)

	if err != nil{
		log.Println("Error in db insert")
		return err
	}
	log.Println("the new Task Insert with ID", id)

	//back to home page
	return Home(c)

}