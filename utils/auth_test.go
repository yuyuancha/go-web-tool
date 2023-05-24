package utils

import "fmt"

func ExampleGenerateHash() {
	// This function is used to generate a hash encrypted string.
	// Applicable to the storage of user's password, etc.

	password := "password"
	hash, err := GenerateHash(password)
	if err != nil {
		// do something to handle error exception.
	}

	fmt.Println("hash:", hash)
	// hash: $2a$10$HeDQcsMbLXJr/eoOctA00eN6Sx8d7RkVVFVefEK5KBtGIl2bwKrU2
}

func ExampleCompareHash() {
	// This function is used to compare input and hash consistent.
	// Applicable to compare password when user login, etc.

	password := "password"
	hash := "$2a$10$HeDQcsMbLXJr/eoOctA00eN6Sx8d7RkVVFVefEK5KBtGIl2bwKrU2"

	isOk, err := CompareHash(password, hash)
	if err != nil {
		// do something to handle error exception.
	}

	if !isOk {
		// do something to handle unauthorized exception.
	}

	// user login.
}
