package basedatos

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx := context.TODO()
	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("usuario")

	registro := make(map[string]interface{}) //make crea un mapa de tipo strings cuyos valores serán de tipo interface
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellido) > 0 {
		registro["apellido"] = u.Apellido
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.Sitioweb) > 0 {
		registro["sitioweb"] = u.Sitioweb
	}

	//registro de actualización para Mongo
	updtString := bson.M{
		"$set": registro, //para Mongo, $set es como UPDATE
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	//filtro es como el WHERE
	filtro := bson.M{"_id": bson.M{"$eq": objID}} //$eq operador Equal

	_, err := col.UpdateOne(ctx, filtro, updtString) //updtString viene con otodos los campos
	if err != nil {
		return false, err
	}

	return true, nil

}
