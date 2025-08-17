package handler

import (
	"fmt"
	"net/http"
)

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Authorization"))
}
