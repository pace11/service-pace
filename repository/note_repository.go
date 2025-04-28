package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type NoteRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.NoteResponse, int, any, int64, int, int)
	Show(id uint) (*models.NoteResponse, int, string, map[string]string)
	Save(note *models.NoteDTO) (any, int, string, map[string]string)
	Update(id uint, note *models.NoteDTO) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type noteRepo struct{}

func NewNoteRepository() NoteRepository {
	return &noteRepo{}
}

func (r *noteRepo) Index(c *gin.Context, filters map[string]any) ([]models.NoteResponse, int, any, int64, int, int) {
	var notes []models.Note
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Note{}), filters)
	query = query.Order("updated_at DESC")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&notes)

	var response []models.NoteResponse
	for _, n := range notes {
		response = append(response, models.NoteResponse{
			ID:          n.ID,
			Title:       n.Title,
			Description: n.Description,
			CreatedAt:   n.CreatedAt,
			UpdatedAt:   n.UpdatedAt,
		})
	}

	return response, http.StatusOK, "note", total, page, limit
}

func (r *noteRepo) Show(id uint) (*models.NoteResponse, int, string, map[string]string) {
	var note models.Note

	if err := config.DB.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, http.StatusNotFound, "note", nil
	}

	response := &models.NoteResponse{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	}

	return response, http.StatusOK, "note", nil
}

func (r *noteRepo) Save(note *models.NoteDTO) (any, int, string, map[string]string) {
	noteCreate := models.Note{
		Title:       note.Title,
		Description: note.Description,
	}

	if err := config.DB.Save(&noteCreate).Error; err != nil {
		return nil, http.StatusInternalServerError, "note", nil
	}

	return note, http.StatusCreated, "note", nil
}

func (r *noteRepo) Update(id uint, note *models.NoteDTO) (any, int, string, map[string]string) {
	var existing models.Note

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "note", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"title":       note.Title,
		"description": note.Description,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "note", nil
	}

	return note, http.StatusOK, "note", nil
}

func (r *noteRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.Note{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "note", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "note", nil
	}

	return nil, http.StatusOK, "note", nil
}
