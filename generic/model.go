package generic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

// User ...
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// UserBrief ...
type UserBrief struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

// UserShort ...
type UserShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetUserBSON ...
func GetUserBSON() []User {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(byteValue, &users)

	return users.Users
}
