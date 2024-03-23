package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"twitterGo/basedatos"
	"twitterGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entré a Registro")

	body := ctx.Value(models.Key("body")).(string) //convierte una variable en String
	err := json.Unmarshal([]byte(body), &t)        //&t es un puntero hacia mi modelo de Usuario
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 { //pregunto si en formato JSON me enviaron el Email
		r.Message = "Debe especificar el Email" //respuesta hacia PostMan o a la API
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 { //pregunto si largo mínimo de password es menor a 6 caracteres
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := basedatos.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese Email"
		fmt.Println(r.Message)
		return r // status 400 va siempre porque lo definimos arriba
	}

	_, status, err := basedatos.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrió un error al intengtar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200 //si llegué acá es porque todo está bien
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r

}
