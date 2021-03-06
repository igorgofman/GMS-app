package gym

import "errors"

type Membership struct {
	Id           int     `json:"id" db:"id"`
	Title        string  `json:"title" db:"title" binding:"required"`
	Price        string  `json:"price" db:"price" binding:"required"`
	Duration     string  `json:"duration" db:"duration" binding:"required"`
	InstructorId *string `json:"instructor_id" db:"instructor_id"`
}

type Visit struct {
	Id      int     `json:"id" db:"id"`
	InfoId  string  `json:"visitor_id" db:"visitor_id"`
	Came_at string  `json:"came_at" db:"came_at"`
	Left_at *string `json:"left_at" db:"left_at"`
}

type UpdateMembershipInput struct {
	Title        *string `json:"title" db:"title"`
	Price        *string `json:"price" db:"price"`
	Duration     *string `json:"duration" db:"duration"`
	InstructorId *string `json:"instructor_id" db:"instructor_id"`
}

func (i UpdateMembershipInput) Validate() error {
	if i.Title == nil && i.Price == nil && i.Duration == nil && i.InstructorId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateVisitInput struct {
	InfoId  *string `json:"visitor_id"`
	Came_at *string `json:"came_at"`
	Left_at *string `json:"left_at"`
}

func (i UpdateVisitInput) Validate() error {
	if i.InfoId == nil && i.Came_at == nil && i.Left_at == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
