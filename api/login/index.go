// Package handler is the Vercel function behind /auth/login (vercel.json rewrite).
package handler

import (
	"net/http"

	"github.com/TheVicky1/GitVerse/web"
)

func Handler(w http.ResponseWriter, r *http.Request) { web.Serve(w, r, (*web.Config).Login) }
