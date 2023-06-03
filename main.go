package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"main.go/awsgo"
	"main.go/bd"
	"main.go/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}
func EjecutoLambda(contexto context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()
	if !ValidoParametro() {
		fmt.Println("Error en los parametros debe enviar SecretManager")
		err := errors.New("Error en los envios de parametros debe enviar Secret Manager")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserMail = att
			fmt.Println("Email = " + datos.UserMail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("sub = " + datos.UserUUID)
		}

	}
	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("Error al leer el secret" + err.Error())
		return event, err
	}
	err = bd.SignUp(datos)
	return event, err

}

func ValidoParametro() bool {

	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro

}
