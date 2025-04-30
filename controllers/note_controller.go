package controllers

import (
	"net/http"
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	Repo repository.NoteRepository
}

func NewNoteController(repo repository.NoteRepository) *NoteController {
	return &NoteController{Repo: repo}
}

// GetNotes
// @Summary Get list of notes
// @Description Retrieve all notes with pagination
// @Tags Notes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Router /notes [get]
func (ctl *NoteController) GetNotes(c *gin.Context) {
	filters := map[string]any{
		"title": c.Query("title"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// GetNote
// @Summary Get a single note by ID
// @Description Get detail of a note by ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} utils.StandardResponses
// @Router /note/{id} [get]
func (ctl *NoteController) GetNote(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
		return
	}

	data, code, entity, errors := ctl.Repo.Show(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// CreateNote
// @Summary Create a new note
// @Description Create a new note with title and description
// @Tags Notes
// @Accept json
// @Produce json
// @Param payload body models.NoteDTO true "Create note payload"
// @Success 201 {object} utils.StandardResponses
// @Router /note [post]
func (ctl *NoteController) CreateNote(c *gin.Context) {
	var note models.NoteDTO

	if utils.BindAndValidate(c, &note) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&note)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// UpdateNote
// @Summary Update a note by ID
// @Description Update the title or description of a note
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param payload body models.NoteDTO true "Update note payload"
// @Success 200 {object} utils.StandardResponses
// @Router /note/{id} [patch]
func (ctl *NoteController) UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note models.NoteDTO

	if utils.BindAndValidate(c, &note) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &note)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// DeleteNote
// @Summary Delete a note by ID
// @Description Delete a note from database
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} utils.StandardResponses
// @Router /note/{id} [delete]
func (ctl *NoteController) DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
