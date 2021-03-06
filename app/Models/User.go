package Models

import (
	"echo-firefly-web/app/Library"
)



type User struct {
	Id           int    `xorm:"not null pk autoincr comment('主键') INT(11)" form:"id"`
	Username       string `xorm:"not null comment('用户名称') VARCHAR(32)" form:"username"`
	Address       string `xorm:"not null comment('用户名称') TEXT" form:"address"`
}
func (User) TableName() string {
	return "user"
}

func (this *User)GetAll() ([]User,error){
	query := Library.InstancetSlave("test")
	datalist := make([]User, 0)
	err := query.Find(&datalist)
	if err != nil {
		return nil,err
	}
	return datalist,nil
}