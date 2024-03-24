package routers

import (
	"context"
	"twitterGo/basedatos"
	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
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

	status, err := basedatos.InsertoRelacion(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar insertar la relación >" + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relación"
		return r
	}

	r.Status = 200
	r.Message = "Alta de Relación OK!"
	return r
}
