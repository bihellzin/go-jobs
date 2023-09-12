package handler

import (
	"fmt"
	"net/http"

	"github.com/bihellzin/go-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(
			ctx,
			http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id),
		)
		return
	}
	sendSuccess(ctx, "show-opening", opening)
}
