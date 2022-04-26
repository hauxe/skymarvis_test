package main

import (
	"context"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

type Game struct {
	s *Storage
}

type User struct {
	Name  string `db:"name"`
	Score int    `db:"score"`
}

type Response struct {
	Number int
	Users  []*User
}

func NewGame() (*Game, error) {
	s, err := NewStorage(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("create storage error: %w", err)
	}
	return &Game{
		s: s,
	}, nil
}

func (g *Game) Play(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid name"))
		return
	}
	score := rand.Intn(10000)
	// store user score
	err = g.s.Store(r.Context(), name, score)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// get leader board
	users, err := g.s.GetScores(r.Context(), 100)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	t, err := template.ParseFiles("./templates/result.tmpl.html", "./templates/header.tmpl.html") // Parse template file.
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = t.Execute(w, &Response{
		Number: score,
		Users:  users,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
