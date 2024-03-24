package jwt

import (
	"errors"
	"strings"
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token inválido") //devuelvo un puntero
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// Rutina que chequea contra la BD
		_, encontrado, _ := basedatos.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return &claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("> token incorrecto")
		// esto es lo que va a llegarle a Postman
	}

	//&claims trae los datos del token: usuario, fecha expiración, etc.
	return &claims, false, string(""), err
}
