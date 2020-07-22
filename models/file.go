package models

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

	_, err = dbmap.Select(&file, "SELECT * FROM files WHERE user_id=?", f.UserID)
	return file, err
}
