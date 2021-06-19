package Models

import (
	"fmt"
	"mvc_pattern/Config"

	_ "github.com/go-sql-driver/mysql"
)

// Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}

// Insert new user
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// Fetch only one user by id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}

// Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)

	Config.DB.Save(user)

	return nil
}

// Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)

	return nil
}
