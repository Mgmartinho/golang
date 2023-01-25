package main

import (
	"database/sql"
	"fmt"
)

// Criar uma rotina para ler todos os itens do banco
func main() {

	var name = "Gui"

	var response = dowithNameButReturnAValue(name)

	fmt.Println("Hello Word: " + response)
}

func do() {
	fmt.Println("")
}

func DBCOnnect() {
	//connection mysql

	db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}

func dowithName(name string) {

}

func dowithNameButReturnAValue(name string) string {
	return name
}
