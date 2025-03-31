package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// To declair a user method. Like a class method.
func (u user) save() {
	fmt.Printf("Saving user %s in datasoruce\n", u.name)
}

func (u user) checkAge() bool {
	return u.age >= 18
}

func (u *user) DoBirthday() {
	u.age++
}

func main() {
	user1 := user{"Leonardo Suave", 29}
	user1.save()	// Not neccessary call by user (name struct) because user1 is a user struct isntance

	newUser := user{"Judite", 17}
	newUser.save()
	newUser.DoBirthday()
	checkNewUserAge := newUser.checkAge()
	fmt.Println(checkNewUserAge)
}