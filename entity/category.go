package entity

type Category struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(200)"`
}
