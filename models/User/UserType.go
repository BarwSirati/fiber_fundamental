package user

type (
	User struct {
		Id       int64  `gorm:"primaryKey" json:"id"`
		Username string `gorm:"type:varchar(50)" json:"username" validate:"required"`
		Password string `gorm:"type:text" json:"password" validate:"required"`
		Name     string `gorm:"type:varchar(50)" json:"name" validate:"required"`
		Lastname string `gorm:"type:varchar(50)" json:"lastname" validate:"required"`
	}
)
