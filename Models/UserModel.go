package Models

// User Table Model
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// TableName tb_user
func (b *User) TableName() string {
	return "tb_user"

}
