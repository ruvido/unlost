package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func scanLibrary(app *pocketbase.PocketBase, libraryPath string) error {
	log.Println("Scanning library...")

	mediaCollection, err := app.FindCollectionByNameOrId("media")
	if err != nil {
		return err
	}

	users, err := app.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	supportedExts := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
		".mp4": true, ".mov": true,
	}

	// Track files found on disk
	filesOnDisk := make(map[string]bool)
	added := 0

	// Phase 1: Scan filesystem, add missing files to DB
	err = filepath.Walk(libraryPath, func(fullPath string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		ext := strings.ToLower(filepath.Ext(fullPath))
		if !supportedExts[ext] {
			return nil
		}

		relPath, err := filepath.Rel(libraryPath, fullPath)
		if err != nil {
			return err
		}

		filesOnDisk[relPath] = true

		// Parse path: username/YYYY/MM/DD/filename.ext
		parts := strings.Split(filepath.ToSlash(relPath), "/")
		if len(parts) < 5 {
			return nil
		}

		username := parts[0]

		// Find user by username
		var userId string
		err = app.DB().
			Select("id").
			From(users.Name).
			AndWhere(dbx.NewExp("username = {:username}", dbx.Params{"username": username})).
			Limit(1).
			Row(&userId)

		if err != nil {
			return nil // Skip if user not found
		}

		// Check if record exists by path
		var existingId string
		err = app.DB().
			Select("id").
			From(mediaCollection.Name).
			AndWhere(dbx.NewExp("path = {:path}", dbx.Params{"path": relPath})).
			Limit(1).
			Row(&existingId)

		if err == nil {
			// Record exists, mark as not missing
			record, _ := app.FindRecordById(mediaCollection.Name, existingId)
			if record != nil {
				record.Set("missing", false)
				app.Save(record)
			}
			return nil
		}

		// Calculate file hash for integrity
		file, err := os.Open(fullPath)
		if err != nil {
			return nil
		}
		defer file.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return nil
		}
		fileHash := hex.EncodeToString(hasher.Sum(nil))

		// Create new record
		record := core.NewRecord(mediaCollection)
		record.Set("path", relPath)
		record.Set("hash", fileHash)
		record.Set("owner", userId)
		record.Set("taken_at", info.ModTime())
		record.Set("missing", false)

		if err := app.Save(record); err != nil {
			return nil
		}

		added++
		return nil
	})

	if err != nil {
		return err
	}

	// Phase 2: Mark DB records as missing if file doesn't exist
	var allRecords []*core.Record
	err = app.RecordQuery(mediaCollection).All(&allRecords)
	if err != nil {
		return err
	}

	marked := 0
	for _, record := range allRecords {
		path := record.GetString("path")
		if !filesOnDisk[path] {
			record.Set("missing", true)
			app.Save(record)
			marked++
		}
	}

	log.Printf("Library scan complete. Added: %d, Marked missing: %d", added, marked)
	return nil
}
