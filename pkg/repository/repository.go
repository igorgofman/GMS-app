package repository

import (
	"github.com/igorgofman/GMS-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gym.SysUser) (int, error)
	GetUser(username, password string) (gym.SysUser, error)
}

type Info interface {
	Create(info gym.Info, member gym.Member, instructor gym.Instructor) (int, error)
	GetAll() ([]gym.DataToPrintInfo, error)
	GetById(infoId int) (interface{}, error)
	Delete(infoId int) error
	// Update(userId, infoId int, input gym.UpdateInfoInput) error
}

type Membership interface {
	// Create(userId, membership gym.Membership) (int, error)
	// GetAll(userId int) ([]gym.Membership, error)
	// GetById(userId, membershipId int) (gym.Membership, error)
	// Delete(userId, membershipId int) error
	// Update(userId, membershipId int, input gym.UpdateMembershipInput) error
}

type Visit interface {
	// Create(infoId int, visit gym.Visit) (int, error)
	// GetAll(userId, infoId int) ([]gym.Visit, error)
	// GetById(userId, visitId int) (gym.Visit, error)
	// Delete(userId, visitId int) error
	// Update(userId, visitId int, input gym.UpdateVisitInput) error
}

type Repository struct {
	Authorization
	Info
	Membership
	Visit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Info:          NewInfoPostgres(db),
		// Membership:      NewMembershipPostgres(db),
		// Visit:      NewVisitPostgres(db),
	}
}
