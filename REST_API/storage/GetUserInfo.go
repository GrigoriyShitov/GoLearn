package storage

import (
	"context"
)

func GetUserInfo(ctx context.Context, id uint) (*User, error) { // to do (*storage.User, error)

	var user *User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error //nil

	}
	// for _, user := range Users {
	// 	if user.ID == id {
	// 		return &user, nil // to do &user
	// 	}
	// }
	return user, nil

}
