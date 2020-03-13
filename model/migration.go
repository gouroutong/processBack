package model

func migration() {
	DB.
		AutoMigrate(&User{}).
		AutoMigrate(&Form{}).
		AutoMigrate(&Process{}).
		AutoMigrate(&Node{})
}
