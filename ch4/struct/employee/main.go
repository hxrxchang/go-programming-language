package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	var dilbert Employee
	dilbert.ID = 1
	dilbert.Name = "dilbert"
	dilbert.Address = "Tokyo"
	dilbert.DoB = time.Now()
	dilbert.Position = "Manager"
	dilbert.Salary = 1000
	dilbert.ManagerID = 1
	AwardAnnualRaise(&dilbert)
	fmt.Println(dilbert)
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
