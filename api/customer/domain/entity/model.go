package customerDomainEntity

type Tabler interface {
	TableName() string
}

type Customer struct {
	ID       int64  `gorm:"column:user_id;primary_key" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
}

func (Customer) TableName() string {
	return "customers"
}
