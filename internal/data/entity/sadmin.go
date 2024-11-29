package entity

func NewSAdminEntity() ISadmin {
	return &sadmin{}
}

type ISadmin interface {
	InjectValue(username, password string)
	GetCredentials() (string, string)
}

type sadmin struct {
	Username string
	Password string
}

func (s *sadmin) InjectValue(username, password string) {
	s.Username = username
	s.Password = password
}

func (s *sadmin) GetCredentials() (string, string) {
	return s.Username, s.Password
}
