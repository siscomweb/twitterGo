package basedatos

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(psw string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), costo)
	//lo que hace bcrypt, encripta un string. El costo es la cantidad de vueltas que va a realizar
	//el módulo de encriptación para encriptar el password. Lo encripta el número de veces que deseamos.
	//8 indica que encriptará 8 veces el passw. Mientras más veces, más demora en hacerlo y consume más recursos
	//8 es el número recomendado. 6 es aceptable.
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil //devuelvo nil en el param error, porque NO hubo error si llegué aquí
}
