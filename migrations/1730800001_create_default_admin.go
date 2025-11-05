package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		adminEmail := os.Getenv("ADMIN_EMAIL")
		adminPassword := os.Getenv("ADMIN_PASSWORD")

		if adminEmail == "" || adminPassword == "" {
			// Skip if credentials not set
			return nil
		}

		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		// Check if admin already exists
		existing, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, adminEmail)
		if existing != nil {
			// Admin already exists, skip
			return nil
		}

		record := core.NewRecord(superusers)
		record.Set("email", adminEmail)
		record.Set("password", adminPassword)
		return app.Save(record)
	}, func(app core.App) error {
		adminEmail := os.Getenv("ADMIN_EMAIL")
		if adminEmail == "" {
			return nil
		}

		record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, adminEmail)
		if record == nil {
			return nil
		}
		return app.Delete(record)
	})
}
