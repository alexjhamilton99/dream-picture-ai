package handler

import (
	"context"
	"dream-picture-ai/pkg/sb"
	"dream-picture-ai/types"
	"net/http"
	"strings"
)

func WithAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := getAuthenticatedUser(r)
		if !user.LoggedIn {
			// path := r.URL.Path
			// http.Redirect(w, r, "/login?to="+path, http.StatusSeeOther) needs a "to" cookie that we delete after the redirect
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		accessTokenCookie, err := r.Cookie("at")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		resp, err := sb.Client.Auth.User(r.Context(), accessTokenCookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user := types.AuthenticatedUser{
			Email:    resp.Email,
			LoggedIn: true,
		}
		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
