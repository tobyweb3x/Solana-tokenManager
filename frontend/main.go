package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"tokenManager/frontend/pages"
)

//go:embed assets
var AssetsDir embed.FS


func main() {
	var staticFiles = fs.FS(AssetsDir)
	staticFs, err := fs.Sub(staticFiles, "assets")
	 if err != nil {
        panic(fmt.Sprintf("failed to create sub filesystem: %v", err))
    }
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := pages.Spa().Render(context.Background(),  w); err != nil {
			panic(err)
		}
    })

    fmt.Println("Starting server at port http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}