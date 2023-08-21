package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoowen/postTsk/checklistItem"
	"github.com/zoowen/postTsk/helper"
)

//tangkap parameter di handler
//handler ke service
//service yang menentukan repository method mana yang di call
//repository : FInd all, FindByUserId
//db

type checklistItemHandler struct {
	service checklistItem.Service
}

func NewChecklistItemHandler(service checklistItem.Service) *checklistItemHandler {
	return &checklistItemHandler{service}
}

func (h *checklistItemHandler) GetChecklistsItem(c *gin.Context) {
	checklistID := c.Param("checklistId")
	checklistIDInt, err := strconv.Atoi(checklistID)

	checklists, err := h.service.GetChecklistItem(checklistIDInt)
	if err != nil {
		response := helper.ApiResponse("Error to get Checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List Of ListItem", http.StatusOK, "success", checklistItem.FormatChecklists(checklists))
	c.JSON(http.StatusOK, response)
}

func (h *checklistItemHandler) CreateChecklistItem(c *gin.Context) {
	checklistID := c.Param("checklistId")
	checklistIDInt, _ := strconv.Atoi(checklistID)
	var input checklistItem.CreateChecklistInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create Checklist", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newChecklist, err := h.service.CreateChecklistItem(input, checklistIDInt)

	if err != nil {
		response := helper.ApiResponse("Failed to create Checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to create Checklist", http.StatusOK, "success", checklistItem.FormatChecklist(newChecklist))
	c.JSON(http.StatusOK, response)

}

func (h *checklistItemHandler) DeleteChecklist(c *gin.Context) {
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
