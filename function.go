package function

import (
	"github.com/pt-suzuki/pig_allowance_book_core/application/infrastructure/server"
	"net/http"
)

func Function (w http.ResponseWriter, r *http.Request) {
	echo := server.CreateEcho()
	echo.ServeHTTP(w, r)
}
