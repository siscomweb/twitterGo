package routers

import (
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	err := basedatos.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar el Tweet >" + err.Error()
		return r
	}

	r.Message = "Eliminar Tweet OK!"
	r.Status = 200
	return r
}
