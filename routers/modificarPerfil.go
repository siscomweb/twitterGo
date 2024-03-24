package routers

import (
	"context"
	"encoding/json"
	"twitterGo/basedatos"
	"twitterGo/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos Incorrectos " + err.Error()
	}

	// status, err :=  como err estaba declarado anteriormente, en versiones nuevas de GO
	// se permite que, al usar :=, al menos una de laas variables de la izquierda sea nueva;
	// si yo utilizo alguna que ya estuvo declarada antes (como err), ya tiene su tipo de dato y solo le asigna algún valor
	status, err := basedatos.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrió un error al intentar modificar el registro. > " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado modificar el registro del usuario. "
		return r
	}

	r.Status = 200
	r.Message = "Modificación de Perfil OK!"
	return r
}
