package models

type SecretManager struct {
	//ALT + 96 backtist
	//creamos los campos que creamos en SecretManager de AWS
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign  string `json:"jwtsign"`
	Database string `json:"database"`
}
