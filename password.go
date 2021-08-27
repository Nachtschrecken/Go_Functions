// This function checks if a provided password matches the stored one
// You can use this as a single password security for your program
// If the password doesn't match, the program stops. You can change this part (line 35)
// to your favor

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

func login() {

	// Prompt password
	fmt.Println("Password: ")
	// Input password, automatically gets hashed and hidden in terminal
	pw, err := terminal.ReadPassword(0)
	// error handling
	if err != nil {
		return 
	}

	// Check if Password is Correct

	// Hash Password
	hash1, _ := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)          	
	// Compare to stored Password
	err = bcrypt.CompareHashAndPassword(hash1, []byte("yourstoredpassword"))
	if err != nil {
		// Print error if didn't match
		log.Println(err)
		// Exit if wrong
		os.Exit(0)
	} else {
		// Continue if right
		return
	}
}