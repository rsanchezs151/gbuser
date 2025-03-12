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

	sentencia := "Insert into users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ejecucion exitosa!")
	return nil
}
