package Models

// Employee ... Employee Model
type Employee struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
	SEX  string `json:"sex"`
	NID  string `json:"nid"`
}

// TableName ... Employee Table
func (b *Employee) TableName() string {
	return "tb_employee"
}
