/*
  Bcrypt Password hasher
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

	if len(os.Args) == 4 && os.Args[1] == "hash" {
		// hash pass
		cost, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Println(err)
			os.Exit(3)
		}
		hash, err := hashAndSalt([]byte(os.Args[2]), cost)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(hash)
		os.Exit(0)
	} else if len(os.Args) == 4 && os.Args[1] == "check" {
		// check
		pass := []byte(os.Args[3])
		//fmt.Println("hash: ", string(os.Args[2]), " and pass ", string(os.Args[3]))
		err := comparePasswords(os.Args[2], pass)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(true)
		os.Exit(1)
	} else {
		// No recognisable args
		printHowTo()
		os.Exit(0)
	}
}

// Print How-to
func printHello() {
	fmt.Println(`
	This program use bcrypt and variable cost to generate hash for Password.

	This implement the cutting edge and most recommend way of hash generation for password.

	Since this program is basically compiled into bytecode it can generate
	very high cost hash at a very high speed.

	Remember that password bruteforcer aren't using you cheap single thread
	php code. Therefore, you shouldn't probably use PHP for hashing.

	`)
}

func printHowTo() {
	fmt.Println(`Argument invalid.
Usage:

./bcrypt-password-hash hash 'pass' cost
	- hash is the function to use
	- pass is the password to hash
	- cost (Integer) between 4 and 31 is the hardning cost to use...
	- anything under 10 is not recommended. The higher the harder to break.

./bcrypt-password-hash check 'hash' 'pass'
	- check is the function to use
	- hash is saved hash to check against
	- pass is the clear text password to compare with the hash

In all cases an exist return value other than 0 mean that an error occurs

	`)
}

// Take Password and cost and return hash
func hashAndSalt(pwd []byte, cost int) (strHash string, err error) {

	hash, intErr := bcrypt.GenerateFromPassword(pwd, cost)
	strHash = string(hash)
	err = intErr
	return
}

// Compare hash and Password
func comparePasswords(hashedPwd string, plainPwd []byte) (err error) {
	byteHash := []byte(hashedPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return
}
