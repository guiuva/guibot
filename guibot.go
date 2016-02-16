package main

import (
	"net/http"
	"os"

	"github.com/adrm/grama"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		grama.Send(grama.Reply{ /*TODO Rellenar*/ })
	})
	grama.Start(os.Getenv("GUIBOT_TOKEN"), "GUIAsistenteBot", os.Getenv("GUIBOT_PORT"))
}
