package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rsanchezs151/gbuser/models"
	"github.com/rsanchezs151/gbuser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

/*
	crear bariable Db en mayuscula ya que es publica

// todo lo relacionado con BD se maneja con tipo punteo
// por tema de velocidad, ya que va a direccion de menoria
// donde estara la conexion y el buffer de la base de datos
*/
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conecion exitosa a BD")
	return nil

}

// formatear el json a String
func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gbecommerce"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
