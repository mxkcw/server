package common

import "time"

type UserInfo struct {
	ID            int64     `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	Password      string    `json:"password" db:"password"`
	NickName      string    `json:"nick_name" db:"nick_name"`
	Phone         string    `json:"phone" db:"phone"`
	Avatar        string    `json:"avatar" db:"avatar"`
	Sex           int       `json:"sex" db:"sex"`
	Memo          string    `json:"memo" db:"memo"`
	LastLoginIP   string    `json:"last_login_ip" db:"last_login_ip"`
	LastLoginTime time.Time `json:"last_login_time" db:"last_login_time"`
	Status        int       `json:"status" db:"status"`
	CreateTime    time.Time `json:"create_time" db:"create_time"`
	UpdateTime    time.Time `json:"update_time" db:"update_time"`
	IsDeleted     bool      `json:"is_deleted" db:"is_deleted"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
