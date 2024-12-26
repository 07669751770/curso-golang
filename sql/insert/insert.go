package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:senha123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	nomes := []string{"Nuno", "Clara", "Leonardo"}

	ids := make([]int64, len(nomes))

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")

	for id, nome := range nomes {
		res, _ := stmt.Exec(nome)
		lid, _ := res.LastInsertId()
		ids[id] = lid
	}
	for _, idInserido := range ids {
		println(idInserido)
	}
}
