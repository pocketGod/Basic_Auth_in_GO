package Repositories

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pocketGod/Basic_Auth_in_GO/Models"
)

const userDBFile = "./Resources/users.json"

func GetUserByUsername(username string) (*Models.User, error) {
	users, err := readUsersFromFile()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, nil // User not found
}

func readUsersFromFile() ([]Models.User, error) {
	var users []Models.User
	file, err := os.Open(userDBFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
