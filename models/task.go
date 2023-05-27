package models

type Task struct{
	Id int              `json:"id" xml:"id" form:"taskId"             orm:"auto"` 
	Subject string      `json:"subject" xml:"Subject" form:"subject"  orm:"size(100)"`
	IsCompleted bool    `json:"isCompleted" xml:"isCompleted" form:"isCompleted"`

}