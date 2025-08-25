package main

import (
	"GO-18-08/simplecalc"
	"fmt"
)
// Structure
type Person struct {
	Name string
	Age  int
}
// Structure
type Address struct {
	Addr1     string
	Apartment string
	City      string
	State     string
	ZipCode   string
	zone      string
}
// Structure
type Account struct {
	Name    string
	Balance int
	Address Address
}
// Pointers and function
func (a *Account) Deposit(amt int) {
	a.Balance += amt
}

func (a Account) BalanceUSD() int {
	return a.Balance
}

func Transfer(from, to *Account, amt int) bool {
	if from.Balance < amt {
		return false
	}
	from.Balance -= amt
	to.Balance += amt
	return true
}
func main() {

	//Map
	accounts := make(map[string]*Account)  

	accounts["Srikar"] = &Account{Name: "Srikar", Balance: 150}
	accounts["Rana"] = &Account{Name: "Rana", Balance: 50}

	a, ok := accounts["Srikar"]
	if ok {
		a.Deposit(40)
	}

	moved := Transfer(accounts["Srikar"], accounts["Rana"], 30)
	fmt.Println("Transfer succes?", moved)

	fmt.Println("Srikar Balance", accounts["Srikar"].BalanceUSD())
	fmt.Println("Rana Balance", accounts["Rana"].BalanceUSD())

	person := Person{
		Name: "Srikar",
		Age:  20888888,
	}
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	home := Address{
		Addr1:     "82 Huntington Rd",
		Apartment: "",
		City:      "Bridgeport",
		State:     "CT",
		ZipCode:   "06608-23",
	}

	office := Address{}
	office.Addr1 = "333 Park Ave"
	office.Apartment = "B12"
	office.City = "Norwalk"
	office.State = "CT"
	office.ZipCode = "06609-23"
	home.zone = "East"

	fmt.Println("Home Address:", home)
	fmt.Println("Office Address:", office)
	fmt.Println("Home City:", home.City)
	fmt.Println("Office City:", office.City)

	strArray := [3]string{"Name", "is", "Srikar"}
	for i, v := range strArray {
		if i < 3 {
			fmt.Println(i, v)
		}
	}
	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Array value")
	for i, v := range arr {
		fmt.Println("Index", i, "Value", v)
	}
	sum := 20
	for _, v := range arr {
		sum = sum + v
	}

	Sub := 20
	for _, v := range arr {
		Sub = Sub - v
	}
	Multi := 20
	for _, v := range arr {
		Multi = Multi * v
	}
	divide := 20
	for _, v := range arr {
		divide = divide / v
	}
	fmt.Println("Sum of array", sum)
	fmt.Println("Sub of array", Sub)
	fmt.Println("Multi of array", Multi)
	fmt.Println("Divide of array", divide)

	slice := arr[1:4]
	fmt.Println("Slice value", slice)

	slice = append(slice, 60)
	fmt.Println("Slice value after append", slice)

	slice = append(slice, 70, 80, 90)

	fmt.Println("Slice value after append", slice)

	fmt.Println("Hello world!")

	num1, num2 := 6, 4
	fmt.Println(simplecalc.Add(num1, num2))

	num3, num4 := 15, 4
	fmt.Println(simplecalc.Sub(num3, num4))

	e, f := 14.0, 8.8
	fmt.Println(simplecalc.Div(e, f))

	g, h := 24, 8
	fmt.Println(simplecalc.Mult(g, h))

	num := 40
	if num > 40 {
		fmt.Println("Number is greater than 40")
	} else if num == 40 {
		fmt.Println("Number is exactly 40")
	} else {
		fmt.Println("Number is less than 40")
	}

	day := "Wednesday"
	switch day {
	case "Wednesday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("It's another day")
	}

	fmt.Println("Counting from 1 to 9:")
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}

	j := 8
	fmt.Println("Before:", j)
	increase(&j)
	fmt.Println("After:", j)

	defer fmt.Println("Second")
	defer fmt.Println("Middle")
	defer fmt.Println("First")

}

func increase(num *int) {
	*num = *num + 1

}
