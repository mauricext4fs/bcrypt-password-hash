/*
 * Bcrypt Password hasher
 */
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Always print the How-to
	printHello()

	// hash pass
	if os.Args[1] == "hash" && len(os.Args) == 4 {
		cost, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Println(err)
			os.Exit(3)
		}
		hash := hashAndSalt([]byte(os.Args[2]), cost)
		fmt.Println(hash)
		os.Exit(0)
		// check
	} else if os.Args[1] == "check" && len(os.Args) == 4 {
		pass := []byte(os.Args[3])
		//fmt.Println("hash: ", string(os.Args[2]), " and pass ", string(os.Args[3]))
		fmt.Println(comparePasswords(os.Args[2], pass))
		os.Exit(0)
		// Not implemented
	} else {
		fmt.Println("Argument invalid. Usage: ")
		fmt.Println("./bcrypt-password-hash hash 'pass' cost")
		fmt.Println(" - hash is the function to use")
		fmt.Println(" - pass is the password to hash")
		fmt.Println(" - cost (Integer) between 4 and 31 is the hardning cost to use...")
		fmt.Println(" anything under 10 is not recommended. The higher the harder to break.")
		fmt.Println(" ")
		fmt.Println("./bcrypt-password-hash check 'hash' 'pass'")
		fmt.Println(" - check is the function to use")
		fmt.Println(" - hash is saved hash to check against")
		fmt.Println(" - pass is the clear text password to compare with the hash")
		fmt.Println(" ")
		fmt.Println("In all cases an exist return value other than 0 mean that an error occurs")
		fmt.Println(" ")
		os.Exit(0)
	}
}

// Print How-to
func printHello() {
	fmt.Println(" ")
	fmt.Println("This program use bcrypt and variable cost to generate hash for Password")
	fmt.Println(" ")
	fmt.Println("This implement the cutting edge and most recommend way of hash generation for password")
	fmt.Println(" ")
	fmt.Println("Since this program is basically compiled into bytecode it can generate ")
	fmt.Println("very high cost hash at a very high speed")
	fmt.Println("")
	fmt.Println("Remember that bruteforcer password aren't usually using you cheap ")
	fmt.Println("single thread shitty php code")
	fmt.Println(" ")
	fmt.Println(" ")
}

// Take Password and cost and return hash
func hashAndSalt(pwd []byte, cost int) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(hash)
}

// Compare hash and Password
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// convert string  to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
