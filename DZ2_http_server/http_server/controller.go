package http_server

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type CommentController struct {
	Service *CommentService
}

func (c *CommentController) RegisterRoutes(r chi.Router) {
	r.Get("/comments", c.GetAllComments)
	r.Get("/comments/{id}", c.GetCommentByID)
	r.Post("/comments", c.CreateComment)
	r.Put("/comments/{id}", c.UpdateComment)
	r.Delete("/comments/{id}", c.DeleteComment)
}

func (c *CommentController) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := c.Service.GetComments()
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Не удалось получить комментарии")
		return
	}
	writeJSON(w, comments)
}

func (c *CommentController) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	comment, err := c.Service.GetComment(id)
	if err != nil {
		ResponseError(w, http.StatusNotFound, "Комментарий не найден")
		return
	}
	writeJSON(w, comment)
}

func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	json.NewDecoder(r.Body).Decode(&comment)

	created, err := c.Service.CreateComment(comment)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Ошибка при добавлении комментария")
		return
	}
	ResponseOk(w, http.StatusCreated, "Комментарий добавлен", created)
}

func (c *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var comment Comment
	json.NewDecoder(r.Body).Decode(&comment)

	updated, err := c.Service.UpdateComment(id, comment)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении комментария")
		return
	}
	ResponseOk(w, http.StatusOK, "Комментарий обновлён", updated)
}

func (c *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := c.Service.DeletePost(id)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Ошибка при удалении комментария")
		return
	}
	ResponseOk(w, http.StatusOK, "Комментарий удалён", nil)
}

// 🔽 Хелперы для ответов
func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ResponseOk(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]any{
		"message": message,
		"data":    data,
	}
	json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(response)
}
