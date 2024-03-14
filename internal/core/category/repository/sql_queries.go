package repository

const (
	qGetAllCategory = `
		SELECT category_id, category_name, is_active, created_at, created_by, updated_at, updated_by, deleted_at
		FROM category
		WHERE deleted_at IS NULL
	`
)