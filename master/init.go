package master

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/simpleRestApiLA/master/account"
	"github.com/vivaldy22/simpleRestApiLA/middleware"
)

func InitRouters(db *sql.DB, r *mux.Router) {
	r.Use(middleware.ActivityLogMiddleware)
	r.Use(middleware.CORSMiddleware)

	accRepo := account.NewRepo(db)
	accUseCase := account.NewUseCase(accRepo)
	account.NewController(accUseCase, r)
}
