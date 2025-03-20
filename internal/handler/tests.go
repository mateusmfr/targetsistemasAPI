package handler

import (
	"net/http"
	"strings"

	"targetsistemas/internal/service"
)

func N1Handler(w http.ResponseWriter, r *http.Request) {
	log := strings.Builder{}

	service.N1(&log)

	w.Write([]byte(log.String()))
}

func N2Handler(w http.ResponseWriter, r *http.Request) {
	log := strings.Builder{}

	service.N2(&log)

	w.Write([]byte(log.String()))
}

func N3Handler(w http.ResponseWriter, r *http.Request) {
	log := strings.Builder{}

	service.N3(&log)

	w.Write([]byte(log.String()))
}

func N4Handler(w http.ResponseWriter, r *http.Request) {
	log := strings.Builder{}

	service.N4(&log)

	w.Write([]byte(log.String()))
}

func N5Handler(w http.ResponseWriter, r *http.Request) {
	log := strings.Builder{}

	service.N5(&log)

	w.Write([]byte(log.String()))
}
