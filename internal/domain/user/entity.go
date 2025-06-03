package user

type User struct {
	ID           int     
	Name         string  
	Surname      string  
	Username     string  
	Email        string  
	Phone        *string 
	Password     string  
  DepartmentID *int
}
