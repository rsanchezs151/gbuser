package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config //aws config estructura que tiene todo lo necesario para configurar el inicio de session
var err error

func InizializarAWS() {
	//TODO retorna sin ningun tipo de configuracion
	Ctx = context.TODO()
	//Configurar config estructuira de tipo AWSConfig, error tambien tiene valores
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error al cargar la configuration .aws/config" + err.Error())
	}
}
