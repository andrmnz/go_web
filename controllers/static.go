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
