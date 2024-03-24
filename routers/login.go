package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"twitterGo/basedatos"
	"twitterGo/jwt"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o Contraseña Inválidos " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido."
		return r
	}
	userData, existe := basedatos.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o Contraseña Inválidos " + err.Error()
		return r
	}

	//función que genera un Json Web Token
	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrió un error al intentar generar el token correspondiente " + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrió un error al intentar formatear el token a JSON > " + err2.Error()
		return r
	}

	//devolvemos una cookie con el token
	//para que, cuando se vuelva a logear, no demore con el proceso
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	//convertir la cookie en String
	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
