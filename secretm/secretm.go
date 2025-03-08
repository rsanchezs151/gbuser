package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rsanchezs151/gbuser/awsgo"
	"github.com/rsanchezs151/gbuser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido secreto " + nombreSecret)
	//svc (secretmanagmentclient) = inicialización del servicio secretsmanager(cfg aws.Config)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura de secret ok " + nombreSecret)
	return datosSecret, nil
}
