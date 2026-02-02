package main

import (
	services "EmployeesManagementSystem/services"
	"fmt"
)

func main() {

	db := services.EmpDb{}
	db.Initialize()

	for {
		fmt.Println("\nwhat do you want to do?")
		fmt.Println("1: add a new department to the database")
		fmt.Println("2: add employee to an existing department")
		fmt.Println("3: remove employee from a department")
		fmt.Println("4: calculate average salary of employees of a department")
		fmt.Println("5: give raise to an emoloyee")
		fmt.Println("6: show department details")
		fmt.Println("7: exit the system")
		var choice int
		fmt.Print("enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			db.AddNewDept()
		case 2:
			db.AddEmpToDept()
		case 3:
			db.RemEmpFromDept()
		case 4:
			db.CalcAvgSal()
		case 5:
			db.RaiseEmpSal()
		case 6:
			db.ShowDeptDetails()
		case 7:
			fmt.Print("exiting the system..\n\n")
			return
		default:
			fmt.Print("invalid choice!\n\n")
		}
	}

}
