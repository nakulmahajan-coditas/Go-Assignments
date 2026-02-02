package services

import (
	models "EmployeesManagementSystem/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Emp = models.Employee
type Dept = models.Department

type EmpDb struct {
	// empMap map[string][]Emp
	empMap map[string]*Dept
}

func (db *EmpDb) Initialize() {
	db.empMap = make(map[string]*Dept)
}

func (db *EmpDb) AddEmpToDept() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter department name: ")
	deptName, _ := reader.ReadString('\n')
	deptName = strings.TrimSpace(deptName)

	dept, ok := db.empMap[deptName]
	if !ok {
		fmt.Printf("department %s doesn't exist!\n", deptName)
		return
	}

	fmt.Print("enter employee's name: ")
	empName, _ := reader.ReadString('\n')
	empName = strings.TrimSpace(empName)

	fmt.Print("enter employee's age: ")
	var empAge uint8
	fmt.Scanln(&empAge)

	fmt.Print("enter emloyee's salary: ")
	var empSal float32
	fmt.Scanln(&empSal)

	// empList := db.empMap[deptName]

	emp := Emp{
		Id:     len(dept.Employees) + 1,
		Name:   empName,
		Age:    empAge,
		Salary: empSal,
	}

	dept.Employees = append(dept.Employees, &emp)
	fmt.Println("employee added successfully!")
}

func (db *EmpDb) ShowDeptDetails() {
	for _, dept := range db.empMap {
		fmt.Printf("department: %s\n", dept.Name)
		for _, emp := range dept.Employees {
			fmt.Printf("Id: %d, Name: %s, Age: %d, Salary: %.2f\n\n", emp.Id, emp.Name, emp.Age, emp.Salary)
		}
	}
}

func (db *EmpDb) RemEmpFromDept() {

	// var deptName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter department: ")
	deptName, _ := reader.ReadString('\n')
	deptName = strings.TrimSpace(deptName)

	var delEmpID int
	fmt.Print("enter id of the employee who has to be deleted: ")
	fmt.Scanln(&delEmpID)

	dept, ok := db.empMap[deptName]
	if !ok {
		fmt.Print("department not founf!\n\n")
		fmt.Println("do you want to add a new department? (0: yes, 1: no)")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			db.AddNewDept()
		case 1:
			return
		default:
			fmt.Print("invalid choice!\n\n")
		}
	}

	for i, emp := range dept.Employees {
		if emp.Id == delEmpID {
			dept.Employees = append(dept.Employees[:i], dept.Employees[i+1:]...)
			fmt.Printf("employee with id: %d and name: %s removed successfully from department %s\n\n", emp.Id, emp.Name, deptName)
			return
		}
	}
	fmt.Println("employee ID not found!")
}

func (db *EmpDb) CalcAvgSal() {
	var deptName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter name of the department to calculate it's average salary: ")
	deptName, _ = reader.ReadString('\n')
	deptName = strings.TrimSpace(deptName)

	dept, ok := db.empMap[deptName]
	if !ok {
		fmt.Printf("departent %s doesn't exist in th database\n\n", deptName)
		return
	}

	if len(dept.Employees) == 0 {
		fmt.Printf("department %s doesn't ave any employees!\n\n", dept.Name)
		return
	}

	var sum float32
	for _, emp := range dept.Employees {
		sum += emp.Salary
	}

	avgSal := sum / float32(len(dept.Employees))
	fmt.Printf("dept %s has %d employees and average salary of employees of %s department is: %.2f\n\n", deptName, len(dept.Employees), deptName, avgSal)
}

func (db *EmpDb) RaiseEmpSal() {
	fmt.Println("enter name of the department in which the employee is present: ")
	reader := bufio.NewReader(os.Stdin)
	deptName, _ := reader.ReadString('\n')
	deptName = strings.TrimSpace(deptName)

	dept, ok := db.empMap[deptName]
	if !ok {
		fmt.Print("enter correct department name!\n\n")
		return
	}

	fmt.Print("enter employee id: ")
	var empID int
	fmt.Scanln(&empID)

	for i := range dept.Employees {
		// for _, emp := range empList {
		if dept.Employees[i].Id == empID {
			// if emp.Id == empID {
			fmt.Println("employee found!")
			fmt.Printf("enter amount to be raised in employee's(id: %d, name: %s) salary: ", dept.Employees[i].Id, dept.Employees[i].Name)
			// fmt.Printf("enter amount to be raised in employee's(id: %d, name: %s) salary: ", emp.Id, emp.Name)
			var amt float32
			fmt.Scanln(&amt)
			dept.Employees[i].Salary += amt
			// emp.Salary += amt
			fmt.Printf("updated salary of employee (id: %d, name: %s) is %.2f\n\n", dept.Employees[i].Id, dept.Employees[i].Name, dept.Employees[i].Salary)
			// fmt.Printf("updated salary of employee (id: %d, name: %s) is %.2f\n\n", emp.Id, emp.Name, emp.Salary)
			return
		}
	}
	fmt.Println("employee nt found!")

}

func (db *EmpDb) AddNewDept() {

	fmt.Println("enter name of the department to be added: ")
	reader := bufio.NewReader(os.Stdin)
	deptName, _ := reader.ReadString('\n')
	deptName = strings.TrimSpace(deptName)

	_, exists := db.empMap[deptName]
	if exists {
		fmt.Printf("department %s already exists!\n\n", deptName)
		return
	}

	// db.empMap[deptName] = []Emp{}
	db.empMap[deptName] = &Dept{
		Name:      deptName,
		Employees: []*Emp{},
	}
	fmt.Printf("department %s added successfully!\n\n", deptName)
}
