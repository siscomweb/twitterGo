package basedatos

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO() // me devuelve un context vacío

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuario") //en MongoDB, colección es como una tabla

	condition := bson.M{"email": email} //esto es como un WHERE

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado) //el primero que encuentre. Lo que encuentra hay que decodificarlo
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID //true significa que el usuario Existe
}
