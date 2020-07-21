package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/petrakypetrov/cloud_clients_api/libs"
)

// StartedStorageGB ...
const StartedStorageGB = 5

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

	_, err = dbmap.Select(&user, "SELECT * FROM users")
	return user, err
}

// Create ...
func (u *User) Create() (user User, err error) {

	currentUnixTs := int64(time.Now().Unix())
	rawTokenText := fmt.Sprintf("%s%d", u.Email, currentUnixTs)

	hasher := md5.New()
	hasher.Write([]byte(rawTokenText))

	u.Token = hex.EncodeToString(hasher.Sum(nil))
	u.StorageGB = StartedStorageGB
	u.AvailableStorageGB = StartedStorageGB
	u.CreateTs = currentUnixTs
	u.UpdateTs = currentUnixTs

	_, err = dbmap.Exec(`
		INSERT INTO users (
			email, 
			password, 
			storage_gb, 
			available_storage_gb,
			token,
			create_ts,
			update_ts
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		u.Email, u.Password, u.StorageGB, u.AvailableStorageGB, u.Token, u.CreateTs, u.UpdateTs,
	)

	if err == nil {
		user.Token = u.Token
	}

	return user, err
}

// GetByPass ...
func (u *User) GetByPass() (user []User, err error) {

	_, err = dbmap.Select(&user, "select * from users where password=?", u.Password)
	return user, err
}

// GetByToken ...
func (u *User) GetByToken() (user []User, err error) {

	_, err = dbmap.Select(&user, "select * from users where token=?", u.Token)
	return user, err
}
