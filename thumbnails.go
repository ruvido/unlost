package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

var thumbnailSizes = map[string]int{
	"small": 400,
	"view":  1920,
}

// generateThumbnails creates small and view thumbnails for the given image
func generateThumbnails(libraryPath, thumbnailPath, relPath string) error {
	originalPath := filepath.Join(libraryPath, relPath)

	// Skip non-image files
	ext := strings.ToLower(filepath.Ext(relPath))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return fmt.Errorf("unsupported format: %s", ext)
	}

	// Open original image with auto-orientation
	src, err := imaging.Open(originalPath, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	// Generate both sizes
	for size, dimension := range thumbnailSizes {
		thumbPath := filepath.Join(thumbnailPath, size, relPath)

		// Create directory structure
		if err := os.MkdirAll(filepath.Dir(thumbPath), 0755); err != nil {
			return fmt.Errorf("failed to create thumbnail dir: %w", err)
		}

		// Generate thumbnail (fit within square, maintain aspect ratio)
		thumb := imaging.Fit(src, dimension, dimension, imaging.Lanczos)

		// Save with JPEG quality 85
		if err := imaging.Save(thumb, thumbPath, imaging.JPEGQuality(85)); err != nil {
			return fmt.Errorf("failed to save thumbnail %s: %w", size, err)
		}
	}

	return nil
}

// thumbnailExists checks if both small and view thumbnails exist
func thumbnailExists(thumbnailPath, relPath string) bool {
	smallPath := filepath.Join(thumbnailPath, "small", relPath)
	viewPath := filepath.Join(thumbnailPath, "view", relPath)

	_, errSmall := os.Stat(smallPath)
	_, errView := os.Stat(viewPath)

	return errSmall == nil && errView == nil
}
