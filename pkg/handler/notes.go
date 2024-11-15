package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"notes/pkg/request"
	"notes/pkg/xerror"
	"strconv"
)

func (h *Handler) CreateNote(c *gin.Context) {
	var input request.CreateNoteRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	output, err := h.Services.Notes.Create(input)
	if handleServiceError(c, err) {
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) UpdateNote(c *gin.Context) {
	var input request.UpdateNoteRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	output, err := h.Services.Notes.Update(input)
	if handleServiceError(c, err) {
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid ID format")
		return
	}

	output, err := h.Services.Notes.Delete(id)
	if handleServiceError(c, err) {
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) GetNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid ID format")
		return
	}

	output, err := h.Services.Notes.Get(id)
	if handleServiceError(c, err) {
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) GetNotes(c *gin.Context) {
	output, err := h.Services.Notes.All()
	if handleServiceError(c, err) {
		return
	}

	c.JSON(http.StatusOK, output)
}

func handleServiceError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	var nfe *xerror.NotFoundError
	var vle *xerror.ValidationError

	switch {
	case errors.As(err, &nfe):
		newErrorResponse(c, http.StatusNotFound, nfe.Error())
	case errors.As(err, &vle):
		newErrorResponse(c, http.StatusBadRequest, vle.Error())
	default:
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return true
}
