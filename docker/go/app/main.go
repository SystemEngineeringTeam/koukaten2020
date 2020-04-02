package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/yj", htmlhundler)
	http.ListenAndServe(":80", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
func htmlhundler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	buff := new(bytes.Buffer)
	str := "<img src=\"data:image/jpeg;base64,/9j/"
	yj := "4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDR"
	yj += "sUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nz"
	yj += "c3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAIAAgAMBIgACEQEDEQH/xAAcAAABBAMBAAAAAAAAAAA"
	yj += "AAAAAAQUGBwIDBAj/xAA4EAABAwIEAwYEBQIHAAAAAAABAAIDBBEFEiExBgdBEyJRYXGBMpGhsRQjQlLRYnIVFiSSorLw/8QA"
	yj += "GQEBAAMBAQAAAAAAAAAAAAAAAAIDBAEF/8QAIREBAQADAAMAAQUAAAAAAAAAAAECAxESITFBBBMiI1H/2gAMAwEAAhEDEQA/"
	dt := "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\">"
	dt = strings.Replace(dt, "P", "9", -1)
	str += yj + dt
	// fw := io.Writer(buff)
	dat := struct {
		Body template.HTML
	}{
		Body: template.HTML(str),
	}
	if err := tmpl.ExecuteTemplate(w, "sc", dat); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buff.Bytes()))
}
