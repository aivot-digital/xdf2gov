package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"goxdf2gov/src/models/xdf2_models"
	"goxdf2gov/src/processors"
	"io"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, "Stammdatenschema konnte nicht korrekt verarbeitet werden.", err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var stammdatenschema xdf2_models.Stammdatenschema0102
		err = xml.Unmarshal(requestBody, &stammdatenschema)
		if err != nil {
			writeError(w, "Stammdatenschema konnte nicht korrekt verarbeitet werden.", err.Error(), http.StatusBadRequest)
			return
		}

		if stammdatenschema.Stammdatenschema.Identifikation.Id == "" {
			writeError(w, "Stammdatenschema hat keine ID", "", http.StatusBadRequest)
			return
		}

		form, err := processors.ProcessStammdatenschema(stammdatenschema.Stammdatenschema)
		if err != nil {
			writeError(w, "Stammdatenschema konnte nicht korrekt verarbeitet werden.", err.Error(), http.StatusBadRequest)
			return
		}

		writeJson(w, form, http.StatusOK)
	})

	err := http.ListenAndServe("0.0.0.0:9595", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeError(w http.ResponseWriter, err string, details string, status int) {
	writeJson(w, map[string]string{"error": err, "details": details}, status)
}

func writeJson(w http.ResponseWriter, data any, status int) {
	bytes, err := json.Marshal(data)
	if err != nil {
		writeError(w, "Fehler beim Serialisieren der Daten.", err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}
