package auth

import "fmt"

// if you write func name with capital letter then it will be exported and can be used outside the package
// this is the scope of the function
func LoginWithCredentials(username, password string) {
	fmt.Println("Login user using", username, password)
}


