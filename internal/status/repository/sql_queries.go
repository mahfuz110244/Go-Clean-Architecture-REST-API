package repository

const (
	createStatus = `INSERT INTO status (name, description, created_by, updated_by) 
					VALUES ($1, $2, $3, $4) 
					RETURNING *`

	updateStatus = `UPDATE status 
					SET description = $1,
						active = $2,
						order_number = $3,
						updated_by = $4
					WHERE id = $5
					RETURNING *`

	getStatusByID = `SELECT n.id,
							n.created_by,
							n.updated_by,
							n.created_at,
							n.updated_at,
							n.deleted_at,
							n.name,
							n.description,
							n.active,
							n.order_number
					FROM status n
					WHERE id = $1`

	deleteStatus = `UPDATE status
					SET deleted_at = CURRENT_TIMESTAMP
					WHERE id = $1`

	getTotalCount = `SELECT COUNT(id) FROM status WHERE deleted_at IS NULL`

	getStatus = `SELECT id, name, description, active, order_number
				FROM status 
				WHERE deleted_at IS NULL
				ORDER BY order_number, created_at, updated_at OFFSET $1 LIMIT $2`

	findByTitleCount = `SELECT COUNT(*)
					FROM status
					WHERE description ILIKE '%' || $1 || '%'`

	findByTitle = `SELECT id, name, description, active, order_number
					FROM status
					WHERE description ILIKE '%' || $1 || '%'
					ORDER BY order_number, created_at, updated_at
					OFFSET $2 LIMIT $3`
)
