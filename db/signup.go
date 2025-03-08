package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rsanchezs151/gbuser/models"
	"github.com/rsanchezs151/gbuser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")
	err := DbConnect()

	if err != nil {
		return err
	}

	// permite colocar la intrucion al final de la funcion lo que hara sera cerrar la bd
	defer Db.Close()

	sentencia := "Inser into users (User_UUID, User_Email, User_DateAdd) VALUES ('" + sig.UserUUID + "','" + sig.UserEmail + "','" + tools.FechaMySQL() + "')"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ejecucion exitosa!")
	return nil
}
