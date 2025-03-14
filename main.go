package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/rsanchezs151/gbuser/awsgo"
	"github.com/rsanchezs151/gbuser/db"
	"github.com/rsanchezs151/gbuser/models"
)

func main() {
	lambda.Start(EjecWSLambda)
}

func EjecWSLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InizializarAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros, debe enviar el 'SecretName'")
		err := errors.New("error: En los parámetros debe enviar SecretName ")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret = " + err.Error())
		return event, err
	}

	err = db.SignUp(datos)

	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
