package models

import (
	"time"
)

// Folder ...
type Folder struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	UserID   int64  `db:"user_id" json:"user_id"`
	CreateTs int64  `db:"create_ts" json:"create_ts"`
	UpdateTs int64  `db:"update_ts" json:"update_ts"`
}

// GetAllByUserID ...
func (f *Folder) GetAllByUserID() (folder []Folder, err error) {

	_, err = dbmap.Select(&folder, "SELECT * FROM folders WHERE user_id=? ORDER BY id DESC", f.UserID)
	return folder, err
}

// Create ...
func (f *Folder) Create() (folder Folder, err error) {

	f.CreateTs = time.Now().Unix()
	f.UpdateTs = f.CreateTs

	_, err = dbmap.Exec(`
		INSERT INTO folders (
			name,  
			user_id,
			create_ts,
			update_ts
		) 
		VALUES (?, ?, ?, ?)`,
		f.Name, f.UserID, f.CreateTs, f.UpdateTs,
	)

	return folder, err
}

// Delete ...
func (f *Folder) Delete() (folder Folder, err error) {
	_, err = dbmap.Exec(`DELETE FROM folders WHERE user_id=? AND id=?`, f.UserID, f.ID)
	_, err = dbmap.Exec(`DELETE FROM files WHERE user_id=? AND folder_id=?`, f.UserID, f.ID)
	return folder, err
}
