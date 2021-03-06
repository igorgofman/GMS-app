package gym

import "errors"

type Info struct {
	Id               int    `json:"id" db:"id"`
	First_Name       string `json:"first_name" db:"first_name" binding:"required"`
	Last_Name        string `json:"last_name" db:"last_name" binding:"required"`
	Middle_Name      string `json:"middle_name" db:"middle_name"`
	Relationship     string `json:"relationship" db:"relationship" binding:"required"`
	Phone            string `json:"phone" db:"phone" binding:"required"`
	Date_of_birth    string `json:"date_of_birth" db:"date_of_birth" binding:"required"`
	Date_of_registry string `json:"date_of_registry" db:"date_of_registry"`
	Instructor
	Member
}

type SysUser struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Member struct {
	InfoId       int
	MembershipId string `json:"membership_id" db:"membership_id"` //binding:"required"
	Expires_at   string `json:"expires_at" db:"expires_at"`       //binding:"required"
}

type Instructor struct {
	InfoId    int
	Hire_date string `json:"hire_date" db:"hire_date"`
	Salary    string `json:"salary" db:"salary"` //binding:"required"
}

type DataToPrintInfo struct {
	Id               int    `json:"id" db:"id"`
	First_Name       string `json:"first_name" db:"first_name"`
	Last_Name        string `json:"last_name" db:"last_name"`
	Relationship     string `json:"relationship" db:"relationship"`
	Phone            string `json:"phone" db:"phone"`
	Date_of_registry string `json:"date_of_registry" db:"date_of_registry"`
}

type DataToPrintMember struct {
	Id               int    `json:"id" db:"id"`
	First_Name       string `json:"first_name" db:"first_name"`
	Last_Name        string `json:"last_name" db:"last_name"`
	Middle_Name      string `json:"middle_name" db:"middle_name"`
	Relationship     string `json:"relationship" db:"relationship"`
	Phone            string `json:"phone" db:"phone"`
	Date_of_birth    string `json:"date_of_birth" db:"date_of_birth"`
	Date_of_registry string `json:"date_of_registry" db:"date_of_registry"`
	MembershipId     string `json:"membership_id" db:"membership_id"`
	Expires_at       string `json:"expires_at" db:"expires_at"`
}

type DataToPrintInstructor struct {
	Id               int    `json:"id" db:"id"`
	First_Name       string `json:"first_name" db:"first_name"`
	Last_Name        string `json:"last_name" db:"last_name"`
	Middle_Name      string `json:"middle_name" db:"middle_name"`
	Relationship     string `json:"relationship" db:"relationship"`
	Phone            string `json:"phone" db:"phone"`
	Date_of_birth    string `json:"date_of_birth" db:"date_of_birth"`
	Date_of_registry string `json:"date_of_registry" db:"date_of_registry"`
	Hire_date        string `json:"hire_date" db:"hire_date"`
	Salary           string `json:"salary" db:"salary"`
}

type UpdateInfoInput struct {
	First_Name    *string `json:"first_name"`
	Last_Name     *string `json:"last_name"`
	Middle_Name   *string `json:"middle_name"`
	Relationship  *string `json:"relationship"`
	Phone         *string `json:"phone"`
	Date_of_birth *string `json:"date_of_birth"`
	MembershipId  *string `json:"membership_id"`
	Expires_at    *string `json:"expires_at"`
	Salary        *string `json:"salary"`
}

func (i UpdateInfoInput) Validate() error {
	if i.First_Name == nil && i.Last_Name == nil && i.Middle_Name == nil && i.Relationship == nil && i.Phone == nil && i.Date_of_birth == nil && i.MembershipId == nil && i.Expires_at == nil && i.Salary == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

// type UpdateItemInput struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// 	Done        *bool   `json:"done"`
// }

// func (i UpdateItemInput) Validate() error {
// 	if i.Title == nil && i.Description == nil && i.Done == nil {
// 		return errors.New("update structure has no values")
// 	}
// 	return nil
// }
