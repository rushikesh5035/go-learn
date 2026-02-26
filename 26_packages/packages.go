package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rushikesh5035/podcast/auth"
	"github.com/rushikesh5035/podcast/user"
)


func main() {
	auth.LoginWithCredentials("rushikesh", "password123")

	session := auth.GetSession()
	
	// auth.extractSession() // this will give error because it is a private function and cannot be accessed outside the package

	fmt.Println(session)

	user := user.User{
		Email: "user@gmail.com",
		Name: "Rushikesh",
	}

	// fmt.Println(user)

	// print user details in RED color using color package
	color.Red("User email: %s, User name: %s", user.Email, user.Name)
}


// to use the open sourse package -> go to `https://pkg.go.dev/` and search for the package you want to use and then follow the instructions to install it and use it in your code.

// to install package -> 
// go get github.com/fatih/color 
// check this package inside the go.mod file and then import it in your code and use it.

// `go mod tidy` -> this command will remove the unused packages from the go.mod file and go.sum file. It will also add the missing packages to the go.mod file and go.sum file. It is a good practice to run this command after installing or removing any package.