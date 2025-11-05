package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		// Get existing users collection
		users, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		// Add username field
		users.Fields.Add(
			&core.TextField{
				Name:     "username",
				Required: true,
				Max:      100,
				Pattern:  `^[a-zA-Z0-9_-]+$`,
			},
		)

		if err := app.Save(users); err != nil {
			return err
		}

		// Create media collection
		media := core.NewBaseCollection("media")
		media.ListRule = types.Pointer("owner = @request.auth.id")
		media.ViewRule = types.Pointer("owner = @request.auth.id")
		media.CreateRule = types.Pointer("@request.auth.id != ''")
		media.UpdateRule = types.Pointer("owner = @request.auth.id")
		media.DeleteRule = types.Pointer("owner = @request.auth.id")

		media.Fields.Add(
			&core.TextField{
				Name:     "path",
				Required: true,
				Max:      500,
			},
			&core.TextField{
				Name:     "hash",
				Required: true,
				Max:      64,
			},
			&core.RelationField{
				Name:         "owner",
				Required:     true,
				MaxSelect:    1,
				CollectionId: users.Id,
			},
			&core.DateField{
				Name:     "taken_at",
				Required: false,
			},
			&core.BoolField{
				Name:     "missing",
				Required: false,
			},
		)

		if err := app.Save(media); err != nil {
			return err
		}

		// Create albums collection
		albums := core.NewBaseCollection("albums")
		albums.ListRule = types.Pointer("owner = @request.auth.id")
		albums.ViewRule = types.Pointer("owner = @request.auth.id")
		albums.CreateRule = types.Pointer("@request.auth.id != ''")
		albums.UpdateRule = types.Pointer("owner = @request.auth.id")
		albums.DeleteRule = types.Pointer("owner = @request.auth.id")

		albums.Fields.Add(
			&core.TextField{
				Name:     "name",
				Required: true,
				Max:      100,
			},
			&core.RelationField{
				Name:         "owner",
				Required:     true,
				MaxSelect:    1,
				CollectionId: users.Id,
			},
			&core.RelationField{
				Name:         "media",
				Required:     false,
				MaxSelect:    0,
				CollectionId: media.Id,
			},
		)

		return app.Save(albums)
	}, func(app core.App) error {
		// Rollback: delete collections
		albums, _ := app.FindCollectionByNameOrId("albums")
		if albums != nil {
			app.Delete(albums)
		}

		media, _ := app.FindCollectionByNameOrId("media")
		if media != nil {
			app.Delete(media)
		}

		users, _ := app.FindCollectionByNameOrId("users")
		if users != nil {
			users.Fields.RemoveById(users.Fields.GetByName("username").GetId())
			app.Save(users)
		}

		return nil
	})
}
