package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		media, err := app.FindCollectionByNameOrId("media")
		if err != nil {
			return err
		}

		// Add thumbs_generated field
		media.Fields.Add(
			&core.BoolField{
				Name:     "thumbs_generated",
				Required: false,
			},
		)

		return app.Save(media)
	}, func(app core.App) error {
		// Rollback: remove thumbs_generated field
		media, err := app.FindCollectionByNameOrId("media")
		if err != nil {
			return err
		}

		field := media.Fields.GetByName("thumbs_generated")
		if field != nil {
			media.Fields.RemoveById(field.GetId())
			app.Save(media)
		}

		return nil
	})
}
