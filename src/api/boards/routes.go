package boards

import (
	"net/http"

	"GOLDEN_LOTUS/api"
)

type Handler struct{}
type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type Data struct {
	Boards []Board `json:"boards"`
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", getBoards)
	return r
}

func getBoards(w http.ResponseWriter, r *http.Request) {
	var data Data
	data = Data{
		Boards: []Board{
			{
				ID:          1,
				Name:        "test",
				DisplayName: "test",
			},
		},
	}
	response := map[string]any{
		"status": "OK",
		"data":   data,
	}
	writeJSON(w, response)
}
