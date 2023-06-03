package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"main.go/models"
	"main.go/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentencia := "INSERT INTO users(User_Email,User_UUID,User_DateAdd)VALUES('" + sig.UserMail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"

	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
