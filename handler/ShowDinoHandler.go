package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Andesson/arkscript/utils"
)

func ImportDinos(db *sql.DB) error {
	var dinos []Dino
	err := utils.ReadJSON(JsonPath+"dinos.json", &dinos)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO dinos (dino, bp, coins, desconto, tokens)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, dino := range dinos {
		_, err := stmt.Exec(dino.Dino, dino.BP, dino.Coins, dino.Desconto, dino.Tokens)
		if err != nil {
			log.Println("Erro ao inserir dino:", dino.Dino, err)
		}
	}
	fmt.Println("Dinos importados com sucesso!")
	return nil
}
