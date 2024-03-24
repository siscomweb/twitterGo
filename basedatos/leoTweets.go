package basedatos

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //paginados de 20 tweets cada página
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //orden ascendente:1, descendente:-1
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(ctx) { //context vacío, porque no lo necesito
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro) //este err, el scope es dentro del for. No se repite con el anterior err
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}
