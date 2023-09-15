package handler

import (
	"fmt"
	"net/http"

	"github.com/bihellzin/go-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.ShouldBindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
