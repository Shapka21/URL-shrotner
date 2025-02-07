package save

import (
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	Save(urlToSave string, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
