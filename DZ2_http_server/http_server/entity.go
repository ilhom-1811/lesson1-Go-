package http_server

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
