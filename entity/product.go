package entity

type Product struct {
	Id int64 `grom:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"productName"`
	Description string `gorm:"type:text" json:"description"`
}