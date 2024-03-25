package routers

import (
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := basedatos.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar la relación >" + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado borrar la relación"
		return r
	}

	r.Status = 200
	r.Message = "Baja Relación OK!"
	return r

}
