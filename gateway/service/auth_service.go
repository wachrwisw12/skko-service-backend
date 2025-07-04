package service

import (
	"database/sql"

	"github.com/yeawyow/gateway/db"
)

type User struct {

	UserID   string
	FullName string
}

func AuthLineService(line_id string) (*User, error) {
	query := "SELECT u.user_id, concat(u.prename,u.user_first_name,u.user_last_name) AS fullname,u.line_id,u.office_id  FROM co_user  u WHERE  u.line_id= ? LIMIT 1"
	row := db.DB.QueryRow(query, line_id)

	var user User
	err := row.Scan(&user.UserID, &user.FullName)
	if err == sql.ErrNoRows {
		// ผู้ใช้ไม่พบ
		return nil, nil
	} else if err != nil {
		// error จาก database
		return nil, err
	}

	return &user, nil
}
