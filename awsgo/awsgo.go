package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	Ctx = context.TODO() //crea un context vacío
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("Error al grabar la configuración .aws/config " + err.Error())
	}
}
