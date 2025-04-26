package controller

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

func (ctl *NoteController) GetNotes(c *gin.Context) {
	filters := map[string]any{
		"title": c.Query("title"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

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

func (ctl *NoteController) CreateNote(c *gin.Context) {
	var note models.NoteDTO

	if utils.BindAndValidate(c, &note) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&note)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *NoteController) UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note models.NoteDTO

	if utils.BindAndValidate(c, &note) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &note)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *NoteController) DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
