// Task 1 :- create the Person class using a struct in Go to represent individuals with attributes like Name, age,
// and methods to introduce themselves, update their age, and check if they are eligible to vote.

//to do:

//create a struct Named Person
//give fields - Name, age
//create reciever methods for - intro, update age on bday, and for checking vote elligibility

// DEDLINE: eod 29th jan 2026
package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) userIntro() {
	fmt.Printf("user's Name is %s and %s is %d years old!\n", p.Name, p.Name, p.age)
}

func (p *Person) updateUserAge() {
	var newAge int
	fmt.Printf("age of %s is %d years\n", p.Name, p.age)
	fmt.Print("enter new age to be updated: ")
	fmt.Scan(&newAge)
	p.age = newAge
	fmt.Printf("updated age of %s is %d", p.Name, p.age)
}

func (p Person) checkVoteEligi() {
	var choice int
	if p.age <= 18 {
		fmt.Printf("age of the user Named %s is %d, since they're below 18, they cant vote\n", p.Name, p.age)
		fmt.Printf("do you want to update %s's age? (0 for YES, 1 for NO)", p.Name)
		fmt.Scan(&choice)
		switch choice {
		case 0:
			p.updateUserAge()
		case 1:
			fmt.Printf("not updating %s's age\n", p.Name)
		default:
			fmt.Println("please enter valid choice!")
		}

	} else {
		fmt.Printf("age of %s is %d, they can vote!\n", p.Name, p.age)
	}
}

func main() {

	p := Person{
		Name: "nakul",
		age:  3,
	}

	p.userIntro()
	p.checkVoteEligi()

}

//WORKING...

//userIntro:
// when called the userIntro function, it just gives the users intro in a line
//-------------------------------
//checkVoteEligi:
//when called this function, it checks for user's age if they're above 18 or not,
//if they're above 18, it print - user can vote...
//if they're below 18, it gives an option to update their age, if choose to update,..>

//it then calls th updateUserAge function, which asks for new age, stores that age in a variable and then assigns the newAge to the original variable
//and after assigning/ successfully updating the user's age, it prints the new age
//-------------------------------
//updateUserAge:
//this fcuntion first prints the old age of the user
//then asks for new age on the terminal
//it scans the new age in the newAge variable
//assigns that value to orignial variable and prints a statement stating the user's updated/ new age
//-------------------------------
