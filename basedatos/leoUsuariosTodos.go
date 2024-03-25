package basedatos

import (
	"context"
	"fmt"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	//[]*models.Usuario, porque retorna una colección de datos que apunta al modelo Usuario
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("usuario")

	fmt.Println("Me conecté con MongoDB y la colección 'relacion'")
	var results []*models.Usuario
	opciones := options.Find()
	opciones.SetLimit(20) //paginados de 20 tweets cada página
	opciones.SetSkip((page - 1) * 20)

	query := bson.M{ //es como WHERE LIKE
		"nombre": bson.M{"$regex": `(?i)` + search}, //búsqueda por Expresión Regular
	}

	//todo lo que no es FindOne devuelve un cursor
	cur, err := col.Find(ctx, query, opciones)
	if err != nil {
		fmt.Println("No se encontró nada. >" + err.Error())
		return results, false
	}

	var incluir bool
	for cur.Next(ctx) {
		fmt.Println("Entré al for cur . Next(ctx)")
		var s models.Usuario

		err := cur.Decode(&s)
		if err != nil {
			fmt.Println("Decode = " + err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado := ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		//ID de usuario viene en el Token, pero estoy leyendo la base de Mongo
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			//para excluir ciertos campos. Como están seteados como omitempty, si los pongo vacíos no los toma en cuenta
			s.Password = ""
			s.Banner = ""
			s.Avatar = ""
			results = append(results, &s)
		}

	}

	err = cur.Err()
	fmt.Println("El cursor debería contener datos.")
	if err != nil {
		fmt.Println("cur.Err() = " + err.Error())
		return results, false
	}

	fmt.Println("Aquí cierro el cursor.")
	cur.Close(ctx)
	return results, true
}
