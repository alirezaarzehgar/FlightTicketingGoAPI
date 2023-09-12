package models

const (
	USERS_ROLE_ADMIN     = "admin"
	USERS_ROLE_EMPLOYEE  = "employee"
	USERS_ROLE_PASSENGER = "passenger"
)

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Email    string   `gorm:"not null;unique" json:"email,omitempty"`
	Password string   `gorm:"not null" json:"password,omitempty"`
	Role     string   `gorm:"not null" json:"role,omitempty"`
	Tickets  []Ticket `gorm:"many2many:user_tickets" json:"tickets,omitempty"`
}
