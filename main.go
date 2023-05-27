package main

import (
    "log" 
)

func main() {
    
    DBMigrations()
    app:= Server()
    MyURLS(&app) 
   
    log.Fatal(app.Listen(":8000"))
}