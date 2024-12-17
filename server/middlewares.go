package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/GOAT-prod/goathttp/headers"
	"github.com/GOAT-prod/goathttp/json"
	"github.com/GOAT-prod/goatlogger"
	"github.com/golang-jwt/jwt/v5"
)

func PanicRecoveryMiddleware(logger goatlogger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger.SetTag("[PANIC RECOVERY]")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				panicErr := recover()
				if panicErr != nil {
					logger.Panic(fmt.Sprintf("произошла паника: %s", panicErr))
					jsonPanic := map[string]any{
						"panic": "что-то пошло не так",
					}
					_ = json.WriteResponse(w, http.StatusInternalServerError, jsonPanic)
				}

			}()
			next.ServeHTTP(w, r)
		})
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(headers.AccessControlAllowOriginHeader(), headers.AllowedOrigins())
		w.Header().Add(headers.AccessControlAllowMethodsHeader(), headers.AllowedMethods())
		w.Header().Add(headers.AccessControlAllowHeaders(), headers.AllowedHeaders())

		if (*r).Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CommonJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := validateToken(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateToken(token string) error {
	splitToken := strings.Split(token, " ")

	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return errors.New("invalid token")
	}

	jwtToken, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	tokenClaims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return errors.New("invalid token")
	}

	exp, ok := tokenClaims["exp"].(float64)
	if !ok {
		return errors.New("invalid token")
	}

	if time.Now().Unix() > int64(exp) {
		return errors.New("token expired")
	}

	return nil
}
