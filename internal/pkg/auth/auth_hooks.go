package auth

import (
	"abude-backend/internal/common"

	"gorm.io/gorm"
)

func AssignEditor(db *gorm.DB) {
	if db.Statement.Error != nil {
		return
	}
	// Check if the table has a `editor_id` column using GORM's Migrator
	if db.Migrator().HasColumn(db.Statement.Model, "editor_id") {
		// Extract the Editor from context
		creds, ok := db.Statement.Context.Value(CredsKey).(common.Creds)
		if !ok {
			return
		}

		// Add a WHERE condition to filter by EditorID
		db.Statement.SetColumn("EditorID", creds.ID, true)
	}
}
