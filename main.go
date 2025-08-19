package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Addr1     string
	Apartment string
	City      string
	State     string
	ZipCode   string
	zone      string
}

func main() {

	p := Person{
		Name: "Srikar",
		Age:  20,
	}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	home := Address{
		Addr1:     "82 Huntington Rd",
		Apartment: "",
		City:      "Bridgeport",
		State:     "CT",
		ZipCode:   "06608-08",
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

	a := [3]string{"Name", "is", "Srikar"}
	for i, v := range a {
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
}
