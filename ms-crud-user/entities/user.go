package entities

import "github.com/qor/audited"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
	IsDelete bool   `json:"delete" gorm:"default:false"`
	audited.AuditedModel
}
