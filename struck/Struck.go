package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	Position   string
	StartDate  string
	EndDate    string
}

func main() {
	// emi := Employee{
	// 	ID:         1,
	// 	Name:       "John Doe",
	// 	Department: "IT",
	// 	Salary:     50000,
	// 	Position:   "Software Engineer",
	// 	StartDate:  "2020-01-01",
	// 	EndDate:    "",
	// }
	// fmt.Println(emi)
	
	emp := map[any]Employee{
		1: {ID: 1, Name: "John Doe", Department: "IT", Salary: 50000, Position: "Software Engineer", StartDate: "2020-01-01", EndDate: ""},
		2: {ID: 2, Name: "Jane Smith", Department: "HR", Salary: 45000, Position: "HR Manager", StartDate: "2019-02-15", EndDate: ""},
		3: {ID: 3, Name: "Alice Johnson", Department: "Marketing", Salary: 48000, Position: "Marketing Specialist", StartDate: "2021-06-01", EndDate: ""},
		4: {ID: 4, Name: "Bob Brown", Department: "Finance", Salary: 55000, Position: "Financial Analyst", StartDate: "2018-10-10", EndDate: ""},
		5: {ID: 5, Name: "Charlie Davis", Department: "IT", Salary: 52000, Position: "System Administrator", StartDate: "2022-03-20", EndDate: ""},
		6: {ID: 6, Name: "Dana White", Department: "Sales", Salary: 47000, Position: "Sales Executive", StartDate: "2020-08-15", EndDate: ""},
		7: {ID: 7, Name: "Evan Green", Department: "Operations", Salary: 53000, Position: "Operations Manager", StartDate: "2017-09-05", EndDate: ""},
		8: {ID: 8, Name: "Fiona Black", Department: "Legal", Salary: 60000, Position: "Legal Advisor", StartDate: "2016-11-01", EndDate: ""},
		9: {ID: 9, Name: "George White", Department: "IT", Salary: 58000, Position: "DevOps Engineer", StartDate: "2019-04-30", EndDate: ""},
		10: {ID: 10, Name: "Hannah Blue", Department: "HR", Salary: 46000, Position: "Recruiter", StartDate: "2021-01-10", EndDate: ""},
	}
	fmt.Print(emp)
}