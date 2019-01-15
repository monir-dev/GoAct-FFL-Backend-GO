package model

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string 	`json:"created_at"`
	UpdatedAt string 	`json:"updated_at"`
}

// func (User) TableName() string {
// 	return "users"
// }

type Users struct {
	ID            int `json:"id"`
	Name          string
	Email         string
	StaffId       string
	Displayname   string
	Desg          string
	ServiceGroup  string
	Company       string
	Dept          string
	Country       string
	Location      string
	UserType      string
	ParentId      int `gorm:"column:parent_id"`
	CompanyId     int `gorm:"column:company_id"`
	depertmentId  int `gorm:"column:depertment_id"`
	Password      string
	Approved      string
	RememberToken string
	CreatedAt     string
	UpdatedAt     string
}

func (Users) TableName() string {
	return "user_not_using"
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type Permission struct {
	ID          int `json:"id"`
	Name        string
	DisplayName string
	Description string
	CreatedBy   int
	CreatedAt   string
	UpdatedAt   string
}

type PermissionRole struct {
	PermissionId int `gorm:"column:permission_id"`
	RoleId       int `gorm:"column:role_id"`
	Accesses     string
}

func (PermissionRole) TableName() string {
	return "permission_role"
}

type RoleUser struct {
	RoleId int `gorm:"column:role_id"`
	UserId int `gorm:"column:user_id"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
