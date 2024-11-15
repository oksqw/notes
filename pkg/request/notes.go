package request

type CreateNoteRequest struct {
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text" binding:"required"`
}

type UpdateNoteRequest struct {
	Id    int    `json:"id" db:"id" binding:"required"`
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text" binding:"required"`
}
