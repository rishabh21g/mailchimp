package types

type Recipient struct {
	Name  string
	Email string
}

type User struct {
	ID       uint   `gorm:"primaryKey; autoIncrement"`
	Fullname string `json:"fullname" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	Password string `json:"password" binding:"required,min=8"`
}
