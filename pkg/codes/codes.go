package codes

const (
	StatusOK           = 200
	StatusCreated      = 201
	StatusBadRequest   = 400
	StatusUnauthorized = 401
	StatusForbidden    = 403
	StatusNotFound     = 404
	StatusServerError  = 500
)

const (
	LangEn = "en"
	LangID = "id"
)

var Messages = map[string]map[int]string{
	LangEn: en,
	LangID: id,
}

func GetMessage(lang string, code int) string {
	if msg, ok := Messages[lang][code]; ok {
		return msg
	}
	return ""
}
