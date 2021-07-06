package rest

import (
	"local/panda-killer/pkg/domain/usecase"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func CreateRouter(accountUsecase *usecase.AccountUsecase, transferUsecase *usecase.TransferUsecase, authUsecase *usecase.AuthUsecase) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Content-type", "application/json"))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", CreateAccount(accountUsecase))
		r.Get("/", GetAccounts(accountUsecase))
		r.Get("/{accountID}/balance", GetAccountBalance(accountUsecase))
	})

	r.Route("/transfers", func(r chi.Router) {
		r.Use(JwtAuthentication)

		r.Post("/", CreateTransfer(transferUsecase))
		r.Get("/", ListTransfers(transferUsecase))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", Login(authUsecase))
	})
	return r
}
