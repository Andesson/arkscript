package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GeraScriptChibi(ctx *gin.Context) {
	request := CreateChibiRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		log.Println("Erro: ", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	CreateScriptRequest := generateChibi(request)
	fmt.Println(CreateScriptRequest)
	ctx.String(http.StatusOK, CreateScriptRequest)
}

func generateChibi(request CreateChibiRequest) string {
	var bdChibi string
	err := db.QueryRow("SELECT bpchibi FROM chibis WHERE chibi = ?", request.Dino).Scan(&bdChibi)
	if err != nil {
		log.Println("Erro: ", err)
		return ""
	}
	command := fmt.Sprintf("woolyitem %s %d %d %d %d %d %d",
		bdChibi,
		1,
		0,
		0,
		0,
		0,
		0,
	)

	return command
}
