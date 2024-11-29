package ucsystem

import (
	"app/internal/data/repo"
	"errors"
)

func NewSystemUseCase(sadminRepo repo.ISadminRepo) ISystemUC {
	return &systemUC{
		sadminRepo: sadminRepo,
	}
}

type ISystemUC interface {
	VerifySAdminLogin(input InputSadmin) error
}

type systemUC struct {
	sadminRepo repo.ISadminRepo
}

func (uc *systemUC) VerifySAdminLogin(input InputSadmin) error {
	user, password := uc.sadminRepo.GetCredentials()
	if user != input.Username || password != input.Password {
		return errors.New("invalid credentials")
	}
	return nil
}
