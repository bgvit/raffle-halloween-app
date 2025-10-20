package controller

import (
	"math/rand/v2"
	"raffle-backend/cmd/web/model"
	"raffle-backend/internal/database"

	"github.com/gin-gonic/gin"
)

func createRaffle(c *gin.Context) {

	n := uint(rand.Uint32N(9))
	raffle := model.Raffle{
		WinningNumber: n,
		WasSelected:   false,
	}
	db := database.Get()

	err := db.Model(&model.Raffle{}).Create(raffle)
	if err != nil {
		return
	}
}
