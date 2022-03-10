package models

import (
	"time"
	orm "todo-list/example/database"

	"github.com/google/uuid"
)

type UserStauts string

const (
	REGISTERED UserStauts = "registered"
	COMPLETED  UserStauts = "completed"
)

type User struct {
	Id         int64 `integer:"id"`
	Uuid       uuid.UUID
	Name       string     `string:"name"`
	Password   string     `string:"password"`
	Email      string     `string:"email"`
	Status     UserStauts `gorm:"type:string;size:32;default:'idle';check:status IN ('registered', 'completed')" json:"status"`
	Created_by int64      `gorm:"type:integer;" json:"created_by"`
	Updated_by int64      `gorm:"type:integer;" json:"updated_by"`
	Deleted_by int64      `gorm:"type:integer;" json:"deleted_by"`
	Created_at time.Time  `sql:"DEFAULT:'current_timestamp'" json:"created_at"`
	Updated_at time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updated_at"`
	Deleted_at time.Time  `sql:"DEFAULT:'current_timestamp'" json:"deleted_at"`
}

var Users []User

//添加
func (user User) Insert() (id int64, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)
	id = user.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"id", "name"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
