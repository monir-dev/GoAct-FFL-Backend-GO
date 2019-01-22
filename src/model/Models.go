package model

type User struct {
	ID        int       `json:"id"`
	RoleId	  int	    `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string 	`json:"created_at"`
	UpdatedAt string 	`json:"updated_at"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Permission struct {
	ID          int 	`json:"id"`
	Name        string 	`json:"name"`
	Module		string 	`json:"module"`
	CreatedAt   string	`json:"created_at"`
	UpdatedAt   string	`json:"updated_at"`
}

type PermissionRole struct {
	PermissionId int `gorm:"column:permission_id"`
	RoleId       int `gorm:"column:role_id"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}

