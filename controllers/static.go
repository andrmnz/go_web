package controllers

import (
	"net/http"

	"github.com/andrmnz/go_web/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  }
}

func FAQ(tmpl views.Template) http.HandlerFunc {
  questions := []struct{
    Question string
    Answer string
  }{
    {
    Question: "question uno",
    Answer: "answer uno",
    },
    {
    Question: "question dos",
    Answer: "answer dos",
    },
    {
    Question: "question tre",
    Answer: "answer tre",
    },
  }
  
  return func (w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, questions)
  }
}