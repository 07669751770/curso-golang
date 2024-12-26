package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:senha123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(?,?)")

	nomes := make(map[int64]string)

	nomes[2001] = "Bia"
	nomes[2002] = "Valentina"
	nomes[2003] = "João"

	insertError := false

	for id, nome := range nomes {
		res, err := stmt.Exec(id, nome)
		if err != nil {
			insertError = true
		} else {
			lid, _ := res.LastInsertId()
			fmt.Printf("%s foi inserido com id: %d\n", nome, lid)
		}
	}

	if insertError {
		tx.Rollback()
		panic("Erro ao inserir registros")
	}
	tx.Commit()

}
