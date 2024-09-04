package model

type Employee struct {
	Base
	Name     *string `json:"name,omitempty" gorm:"type:varchar(150)"`
	Position *string `json:"position,omitempty" gorm:"type:varchar(100)"`
	Address  *string `json:"address,omitempty" gorm:"type:varchar(255)"`
}
