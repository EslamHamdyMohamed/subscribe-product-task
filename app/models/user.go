package models

import (
	"subscribe-product/database"
	userModel "subscribe-product/database/models/user"
)

func FindOrFailUser(userId uint) (bool, error, userModel.User) {
	var user userModel.User
	db := database.DB
	err := db.Where("id", userId).
		Find(&user).Error
	if user.ID == 0 {
		return false, err, user
	}
	return true, err, user
}
