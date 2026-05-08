package models

import "time"

type User struct {
	Id         int64     `orm:auto`
	Name       string    `orm:size(64)`
	Age        int       `orm:null`
	SubmitTime time.Time `orm:"type(datetime)"`
}

func (u *User) TableName() string {
	return "t_user"
}
