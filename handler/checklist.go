package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoowen/postTsk/checklist"
	"github.com/zoowen/postTsk/helper"
)

//tangkap parameter di handler
//handler ke service
//service yang menentukan repository method mana yang di call
//repository : FInd all, FindByUserId
//db

type checklistHandler struct {
	service checklist.Service
}

func NewChecklistHandler(service checklist.Service) *checklistHandler {
	return &checklistHandler{service}
}

func (h *checklistHandler) GetChecklists(c *gin.Context) {

	checklists, err := h.service.GetChecklist()
	if err != nil {
		response := helper.ApiResponse("Error to get Checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List Of Campaings", http.StatusOK, "success", checklist.FormatChecklists(checklists))
	c.JSON(http.StatusOK, response)
}

func (h *checklistHandler) CreateChecklist(c *gin.Context) {
	var input checklist.CreateChecklistInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create Checklist", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newChecklist, err := h.service.CreateChecklist(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create Checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to create Checklist", http.StatusOK, "success", checklist.FormatChecklist(newChecklist))
	c.JSON(http.StatusOK, response)

}

func (h *checklistHandler) DeleteChecklist(c *gin.Context) {
	checklistID := c.Param("id")

	ID, err := strconv.Atoi(checklistID)
	if err != nil {
		response := helper.ApiResponse("Invalid checklist ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteChecklist(ID)
	if err != nil {
		response := helper.ApiResponse("Error to delete checklist", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Checklist deleted successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
