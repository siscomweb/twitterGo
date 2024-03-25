package basedatos

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) bool {
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	// para Mongo, si con FindOne no se encuentra nada, devuelve un error. En SQL devuelve vacío
	// esto no funciona: quise hacer return (true o false) ==>return (err == nil)
	if err != nil {
		return false
	}
	return true
}
