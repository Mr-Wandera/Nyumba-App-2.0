package middleware

import (
	"context"
	"net/http"
	"nyumba/auth"
	"strings"
)

type contextKey string

const (
	UserContextKey contextKey = "userClaims"
)

// AuthRequired middleware enforces valid authentication token via Cookie or Authorization header
func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ExtractClaims(r)
		if err != nil {
			if strings.Contains(r.Header.Get("Accept"), "text/html") {
				http.Redirect(w, r, "/login?error=session_expired", http.StatusSeeOther)
				return
			}
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// RoleRequired enforces specific roles (e.g. 'landlord', 'admin')
func RoleRequired(requiredRoles []string, next http.HandlerFunc) http.HandlerFunc {
	return AuthRequired(func(w http.ResponseWriter, r *http.Request) {
		claims := GetUserClaims(r)
		if claims == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		allowed := false
		for _, role := range requiredRoles {
			if strings.EqualFold(claims.Role, role) {
				allowed = true
				break
			}
		}

		if !allowed {
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ExtractClaims checks Authorization Bearer header or nyumba_token cookie
func ExtractClaims(r *http.Request) (*auth.TokenClaims, error) {
	// Check Cookie
	cookie, err := r.Cookie("nyumba_token")
	if err == nil && cookie.Value != "" {
		return auth.VerifyToken(cookie.Value)
	}

	// Check Authorization Header
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		return auth.VerifyToken(tokenStr)
	}

	return nil, auth.ErrInvalidToken
}

// GetUserClaims helper retrieves claims from context
func GetUserClaims(r *http.Request) *auth.TokenClaims {
	claims, ok := r.Context().Value(UserContextKey).(*auth.TokenClaims)
	if !ok {
		return nil
	}
	return claims
}
