package data

var UserSeed = []struct {
	Name           string
	Surname        string
	Username       string
	Email          string
	Phone          *string
	Password       string
}{
	{
		Name:           "ระบบ",
		Surname:        "ผู้ดูแล",
		Username:       "admin",
		Email:          "admin@example.com",
		Phone:          nil,
		Password:       "admin123",
	},
	{
		Name:           "มนุษย์",
		Surname:        "จัดการ",
		Username:       "manager",
		Email:          "manager@example.com",
		Phone:          nil,
		Password:       "manager123",
	},
	{
		Name:           "บุคคล",
		Surname:        "ทั่วไป",
		Username:       "user",
		Email:          "user@example.com",
		Phone:          nil,
		Password:       "user123",
	},
}
