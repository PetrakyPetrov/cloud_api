package models

import (
	"github.com/petrakypetrov/cloud_clients_api/libs"
)

// User ...
type User struct {
	ID                 int64   `db:"id" json:"id"`
	Email              string  `db:"email" json:"email"`
	Password           string  `db:"password" json:"password"`
	StorageGB          float64 `db:"storage_gb" json:"storage_gb"`
	AvailableStorageGB float64 `db:"available_storage_gb" json:"available_storage_gb"`
	Token              string  `db:"token" json:"token"`
	CreateTs           int64   `db:"create_ts" json:"create_ts"`
	UpdateTs           int64   `db:"update_ts" json:"update_ts"`
}

var dbmap = libs.DBmap

// Get ...
func (u *User) Get() (user []User, err error) {

	_, err = dbmap.Select(&user, "select * from users")
	return user, err
}
