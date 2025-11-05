package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	_ "unlost/migrations"
)

func init() {
	// Load .env file BEFORE migrations run
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	libraryPath := os.Getenv("LIBRARY_PATH")
	if libraryPath == "" {
		log.Fatal("LIBRARY_PATH not set in .env")
	}

	thumbnailPath := os.Getenv("THUMBNAIL_PATH")
	if thumbnailPath == "" {
		log.Fatal("THUMBNAIL_PATH not set in .env")
	}

	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Background library scan
		go func() {
			time.Sleep(2 * time.Second)
			scanLibrary(app, libraryPath, thumbnailPath)
		}()

		// Serve library files
		se.Router.GET("/library/{path...}", apis.Static(os.DirFS(libraryPath), false))

		// Serve thumbnails (small and view)
		se.Router.GET("/thumbs/{path...}", apis.Static(os.DirFS(thumbnailPath), false))

		// Serve frontend
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
