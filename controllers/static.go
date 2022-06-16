package controllers

import (
	"net/http"
)

func StaticHandler(tmpl Template) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  }
}

func FAQ(tmpl Template) http.HandlerFunc {
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