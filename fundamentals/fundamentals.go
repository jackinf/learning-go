package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

type Admin struct {
	Person Person
	Roles []string
}

func (person Person) PrintDetails() {
	fmt.Printf("FirstName: %s, LastName: %s", person.FirstName, person.LastName)
	fmt.Println()
}

func (admin Admin) PrintDetails() {
	fmt.Print("Admin: ");
	admin.Person.PrintDetails();
}

func (person *Person) ChangeLastName(lastName string) {
	person.LastName = lastName;
}

func runFundamentals() {
	fmt.Println("hello world")

	// array
	var arr = [5]float64{1, 2, 3, 4, 5}
	fmt.Println(arr);

	// slice
	var slice = []string{"test1", "test2", "test3"}
	slice = append(slice, "test4");
	fmt.Println(slice);

	// map
	var map1 = make(map[string]float32);
	map1["key1"] = 1;
	map1["key2"] = 2;
	map1["key3"] = 3;
	for k,v := range map1 {
		fmt.Printf("Key %s, value %.2f\n", k, v)
	}

	// struct
	var person1 = &Person {
		FirstName: "John",
		LastName: "Smith",
	}
	person1.ChangeLastName("Doe");
	person1.PrintDetails();

	var person2 = &Admin {
		Person{
			FirstName: "John",
			LastName: "Smith",
		},
		[]string{"manager", "accountant"},
	}
	person2.Person.ChangeLastName("Adams");
	person2.PrintDetails();
}