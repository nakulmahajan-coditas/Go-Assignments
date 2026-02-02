package models

type Employee struct {
	Id     int
	Name   string
	Age    uint8
	Salary float32
}

type Department struct {
	Name      string
	Employees []*Employee
}
