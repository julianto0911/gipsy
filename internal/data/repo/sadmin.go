package repo

import (
	"app/internal/data/entity"
)

func NewSadminRepo(username, password string) ISadminRepo {
	repo := &sadminRepo{
		sadmin: entity.NewSAdminEntity(),
	}
	repo.sadmin.InjectValue(username, password)
	return repo
}

type ISadminRepo interface {
	GetCredentials() (string, string)
}

type sadminRepo struct {
	sadmin entity.ISadmin
}

func (r *sadminRepo) GetCredentials() (string, string) {
	return r.sadmin.GetCredentials()
}
