package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Andesson/arkscript/utils"
	"github.com/gin-gonic/gin"
)

func ImportDinos(ctx *gin.Context) {
	var dinos []Dino
	err := utils.ReadJSON(JsonPath+"dinos.json", &dinos)
	if err != nil {
		log.Println("Erro ao ler JSON:", err)
		if strings.Contains(err.Error(), "o arquivo não existe") {
			sendError(ctx, http.StatusNotFound, "Arquivo dinos.json não encontrado")
		} else {
			sendError(ctx, http.StatusInternalServerError, "Erro ao ler o arquivo dinos.json")
		}
		return
	}

	stmt, err := db.Prepare(`
		INSERT OR REPLACE INTO dinos (dino, bp, coins, desconto, tokens)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		sendError(nil, http.StatusInternalServerError, "error creating opening on database")
	}
	defer stmt.Close()

	for _, dino := range dinos {
		_, err := stmt.Exec(dino.Dino, dino.BP, dino.Coins, dino.Desconto, dino.Tokens)
		if err != nil {
			log.Println("Erro ao inserir dino:", dino.Dino, err)
		}
	}
	fmt.Println("Dinos importados com sucesso!")
	sendSucess(ctx, "import-dino", "sucefully imported dinos")
}

func ImportItens(ctx *gin.Context) {
	var itens []Item
	err := utils.ReadJSON(JsonPath+"itens.json", &itens)
	if err != nil {
		log.Println("Erro ao ler JSON:", err)
		if strings.Contains(err.Error(), "o arquivo não existe") {
			sendError(ctx, http.StatusNotFound, "Arquivo itens.json não encontrado")
		} else {
			sendError(ctx, http.StatusInternalServerError, "Erro ao ler o arquivo itens.json")
		}
		return
	}

	for _, item := range itens {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO itens (item, bpitem, coinsItem, coinsItemBP, coinsItemDesconto, coinsItemBPDesconto)
			VALUES (?, ?, ?, ?, ?, ?)`,
			item.Item, item.BPItem, item.CoinsItem, item.CoinsItemBP, item.CoinsItemDesconto, item.CoinsItemBPDesconto)
		if err != nil {
			log.Println("Erro ao inserir item:", item.Item, err)
		}
	}
	fmt.Println("Itens importados com sucesso!")
	sendSucess(ctx, "import-item", "sucefully imported itens")
}

func ImportChibis(ctx *gin.Context) {
	var chibis []Chibi
	err := utils.ReadJSON(JsonPath+"chibis.json", &chibis)
	if err != nil {
		log.Println("Erro ao ler JSON:", err)
		if strings.Contains(err.Error(), "o arquivo não existe") {
			sendError(ctx, http.StatusNotFound, "Arquivo chibis.json não encontrado")
		} else {
			sendError(ctx, http.StatusInternalServerError, "Erro ao ler o arquivo chibis.json")
		}
		return
	}

	for _, chibi := range chibis {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO chibis (chibi, bpchibi)
			VALUES (?, ?)`,
			chibi.Chibi, chibi.BPChibi)
		if err != nil {
			log.Println("Erro ao inserir chibi:", chibi.Chibi, err)
		}
	}
	fmt.Println("Chibis importados com sucesso!")
	sendSucess(ctx, "import-chibis", "sucefully imported chibis")
}

func ImportAll(ctx *gin.Context) {
	log.Println("Iniciando importação de todos os dados...")

	ImportDinos(ctx)
	ImportItens(ctx)
	ImportChibis(ctx)

	log.Println("Todos os dados foram importados com sucesso!")
	sendSucess(ctx, "import-all", "successfully imported all data")
}
