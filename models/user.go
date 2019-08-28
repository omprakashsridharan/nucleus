package models

//User structure
type User struct {
	Model
	Name string `gorm:"type:varchar(200)" json:"name" validate:"required"`
}

//TableName return name of database table
func (u *User) TableName() string {
	return "user"
}
