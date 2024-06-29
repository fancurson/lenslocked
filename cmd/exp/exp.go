package main

import (
	"html/template"
	"os"
)

type User struct {
	Bio  string
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Jon Calhoun",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Age:  19,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
