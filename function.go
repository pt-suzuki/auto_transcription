package function

import (
	"github.com/pt-suzuki/auto_transcription/infrastructure/server"
	"net/http"
)

func Function(w http.ResponseWriter, r *http.Request) {
	echo := server.CreateEcho()
	echo.ServeHTTP(w, r)
}
