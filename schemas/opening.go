package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Remote bool
	Link string
	Salary int64
}

type OpeningResponse struct {
	Id 				uint 			`json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role 			string 		`json:"role"`
	Company 	string 		`json:"company"`
	Remote 		bool 			`json:"remote"`
	Link 			string 		`json:"link"`
	Salary 		int64 		`json:"salary"`
}
