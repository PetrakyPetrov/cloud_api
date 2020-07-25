package models

import (
	"fmt"
	"time"
)

// File ...
type File struct {
	ID       int64   `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	FolderID int64   `db:"folder_id" json:"folder_id"`
	UserID   string  `db:"user_id" json:"user_id"`
	SizeKB   float64 `db:"size_kb" json:"size_kb"`
	CreateTs int64   `db:"create_ts" json:"create_ts"`
}

// GetAllByUserID ...
func (f *File) GetAllByUserID() (file []File, err error) {

	if f.FolderID == 0 {
		_, err = dbmap.Select(&file, "SELECT * FROM files WHERE user_id=? ORDER BY id DESC",
			f.UserID,
		)
	} else {
		_, err = dbmap.Select(&file, "SELECT * FROM files WHERE user_id=? AND folder_id=? ORDER BY id DESC",
			f.UserID,
			f.FolderID,
		)
	}

	return file, err
}

// Create ...
func (f *File) Create() (file File, err error) {

	f.CreateTs = time.Now().Unix()

	_, err = dbmap.Exec(`
		INSERT INTO files (
			name, 
			folder_id, 
			user_id, 
			size_kb,
			create_ts
		) 
		VALUES (?, ?, ?, ?, ?)`,
		f.Name, f.FolderID, f.UserID, f.SizeKB, f.CreateTs,
	)

	return file, err
}

// Delete ...
func (f *File) Delete() (file File, err error) {
	fmt.Println(f)
	_, err = dbmap.Exec(`DELETE FROM files WHERE  id=? AND user_id=?`, f.ID, f.UserID)
	return file, err
}
