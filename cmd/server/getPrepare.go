package main

import (
	"log"
	"net/http"
	"os/exec"
)

func getPrepare(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	pdfConversion := exec.Command("wkhtmltopdf", "localhost:8080/app/report.html", "app/report.pdf")

	if err := pdfConversion.Run(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	downloadLinkHTML := `
<a href="report.pdf"
     download="report.pdf">
    Download!
</a>
`
	if _, err := w.Write([]byte(downloadLinkHTML)); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
