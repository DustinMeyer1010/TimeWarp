package handler

import (
	"net/http"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("Authorization")
}
