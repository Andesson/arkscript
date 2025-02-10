package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GeraScriptDino(ctx *gin.Context) {
	request := CreateDinoRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		log.Println("Erro: ", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	CreateScriptRequest := generateDino(request)
	fmt.Println(CreateScriptRequest)
	ctx.String(http.StatusOK, CreateScriptRequest)
}

func generateDino(request CreateDinoRequest) string {
	var bdDino string
	err := db.QueryRow("SELECT bp FROM dinos WHERE dino = ?", request.DinoName).Scan(&bdDino)
	if err != nil {
		log.Println("Erro: ", err)
		return ""
	}
	command := fmt.Sprintf("WOOLYCRYO %s %s %d %d %d %d %d %d %d %d %d %d %d %d",
		fmt.Sprintf("%d", request.SteamId),
		bdDino,
		request.LevelSpinner,
		request.ImprintSpinner,
		request.NeuteredSpinner,
		request.SexSpinner,
		request.HealthSpinner,
		request.StaminaSpinner,
		request.OxygenSpinner,
		request.FoodSpinner,
		request.WeightSpinner,
		request.DamageSpinner,
		request.SpeedSpinner,
		request.TorporSpinner,
	)

	return command
}
