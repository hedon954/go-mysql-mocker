package hello_world

import (
	"database/sql"
)

type UserState struct {
	UID   string `gorm:"primaryKey;column:uid"`
	State string `gorm:"column:state"`
}

func (u UserState) TableName() string {
	return "user_state"
}

// ChangeUserStateToMatch change user state to match
func ChangeUserStateToMatch(db *sql.DB, uid string) (int64, error) {
	res, err := db.Exec("UPDATE user_state SET state = 'match' WHERE uid = ?", uid)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
