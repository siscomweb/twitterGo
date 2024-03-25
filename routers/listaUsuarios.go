package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el parámetro 'page' como entero mayor a 0.>" + err.Error()
		return r
	}

	fmt.Println("Llegué a basedatos.LeoUsuariosTodos")
	usuarios, status := basedatos.LeoUsuariosTodos(IDUsuario, int64(pageTemp), search, typeUser)
	if !status {
		r.Message = "Error al leer los usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios a JSON"
		return r
	}

	fmt.Println("Usuarios encontrados: " + string(respJson))
	r.Status = 200
	r.Message = string(respJson)
	return r
}
