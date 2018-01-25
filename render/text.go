package render

import (
	"github.com/gopub/utils"
	"log"
	"net/http"
)

// Text render text into writer
func Text(writer http.ResponseWriter, text string) {
	writer.Header()["Content-Type"] = []string{utils.MIMETEXT + "; charset=utf-8"}
	data := []byte(text)
	err := utils.WriteAll(writer, data)
	if err != nil {
		log.Println("[WINE] Render error:", err)
	}
}
