// Task 2:- Develop a system to manage employees and their departments.
// Each employee has a name, age, and salary.
// Each department has a name, a list of employees, and a method to calculate the average salary of its employees.
// Additionally, create methods to add
// and remove employees from departments and to give a raise to an employee.

// DEADLINE: eod 29th jan 2026
package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary float32
}

type Department struct {
	name string
	emps []Employee
}

func (d Department) calcAvgSal() {
	var count int
	var totalSal float32

	count = 0
	totalSal = 0

	arr := d.emps
	for _, u := range arr { //here i want to take and add the salary value from the employee list of the department
		//*edit* - done!
		totalSal += u.salary
		count++
	}
	totalAvgSal := totalSal / float32(count)
	fmt.Printf("total employees in department %s are %d and average salary of employees of department %s is %v\n", d.name, count, d.name, totalAvgSal)

	fmt.Println("")
}

func (d *Department) addEmpToDept(newEmp Employee) {
	// var newEmp Employee

	// fmt.Println("this method adds an employee to a department")

	// fmt.Println("enter employee details to be added: ")
	// fmt.Scan(&newEmp)
	// fmt.Scanf("na/me: %s, age: %d, salary: %v", newEmp.name, newEmp.age, newEmp.salary)

	d.emps = append(d.emps, newEmp)

	fmt.Println("\ndepartment: ", d.name)
	fmt.Println("added employee successfully!")
	fmt.Printf("%+v\n", d.emps)
	fmt.Println("")
}

func (d *Department) remEmpFromDept() {
	fmt.Println("removing employee from..", d.name)
	fmt.Printf("%+v\n", d.emps)
	d.emps = d.emps[:len(d.emps)-1]
	fmt.Println("removed employee from..", d.name, "successfully!")
	fmt.Printf("%+v\n", d.emps)
	fmt.Println("")

}

func (e Employee) raiseEmpSal() {
	var raiseAmt float32
	fmt.Printf("enter amount to be raised in %s's salary: ", e.name)
	fmt.Scan(&raiseAmt)

	e.salary += raiseAmt
	fmt.Printf("\nupdated salary of %s is %.2f\n", e.name, e.salary)
	fmt.Println("")
}

func main() {

	nakul := Employee{
		name:   "nakul",
		age:    22,
		salary: 11000,
	}

	ramesh := Employee{
		name:   "ramesh",
		age:    45,
		salary: 78000,
	}

	suresh := Employee{
		name:   "suresh",
		age:    47,
		salary: 83000,
	}

	paul := Employee{
		name:   "paul",
		age:    33,
		salary: 60000,
	}

	john := Employee{
		name:   "john",
		age:    28,
		salary: 68000,
	}

	java := Department{
		name: "java development",
		emps: []Employee{},
	}

	goLang := Department{
		name: "go development",
		emps: []Employee{},
	}

	java.addEmpToDept(suresh)
	java.addEmpToDept(ramesh)
	java.addEmpToDept(john)

	goLang.addEmpToDept(nakul)
	goLang.addEmpToDept(paul)

	java.calcAvgSal()
	goLang.calcAvgSal()

	java.remEmpFromDept()
	java.calcAvgSal()

	nakul.raiseEmpSal()
	goLang.calcAvgSal()
}
