package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection string
	dsn := "root:senha123@/cursogo"

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update statement
	updateStmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")

	alunos := make(map[int64]string)

	alunos[2001] = "Davi"
	alunos[1] = "Caio"
	alunos[2003] = "Antonio"
	alunos[2002] = "Anna Clara"

	listaErros := []error{}
	listaSucessos := []string{}

	for id, nome := range alunos {
		res, err := updateStmt.Exec(nome, id)
		if err != nil {
			log.Fatal(err)
			listaErros = append(listaErros, err)
		} else {
			lid, _ := res.RowsAffected()
			listaSucessos = append(listaSucessos, fmt.Sprintf("%s foi atualizado linhas afetadas: %d\n", nome, lid))
		}

	}

	if len(listaErros) > 0 {
		for _, err := range listaErros {
			log.Fatal(err)
		}
	}
	if len(listaSucessos) > 0 {
		for _, sucesso := range listaSucessos {
			fmt.Println(sucesso)
		}
	}
}
