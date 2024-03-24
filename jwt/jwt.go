package jwt

import (
	"context"
	"time"
	"twitterGo/models"

	"github.com/golang-jwt/jwt/v5"
)

func GeneroJWT(ctx context.Context, t models.Usuario) (string, error) {
	//en parámetro String devuelvo el Token generado
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)
	miClave := []byte(jwtSign)

	//Json Web Token, esta compuesto de 3 partes: Headers, payload y firma
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.Sitioweb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), //duración de 1 día
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave) //miClave sirve para decodificar
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
