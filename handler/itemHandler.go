package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GeraScriptItem(ctx *gin.Context) {
	request := CreateItemRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		log.Println("Erro: ", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	CreateScriptRequest := generateItem(request)
	fmt.Println(CreateScriptRequest)
	ctx.String(http.StatusOK, CreateScriptRequest)
}

func generateItem(request CreateItemRequest) string {
	var bdItem string
	err := db.QueryRow("SELECT bpitem FROM itens WHERE item = ?", request.Item).Scan(&bdItem)
	if err != nil {
		log.Println("Erro: ", err)
		return ""
	}
	command := fmt.Sprintf("WOOLYITEM %s %d %d %d %d %d %d",
		bdItem,
		request.Quantidade,
		request.Qualidade,
		request.BPOrSela,
		request.Durabilidade,
		request.Dano,
		request.Armadura,
	)

	return command
}
