package repository

const (
	createStatus = `INSERT INTO status (name, description, created_by, updated_by) 
					VALUES ($1, $2, $3, $4) 
					RETURNING *`

	updateStatus = `UPDATE status 
					SET name = $1, name),
					description = $2, description),
					active = $3, active),
					order_number = $4, order_number),
					updated_by = $5, updated_by)
					WHERE id = $6
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
					WHERE status_id = $1`

	deleteStatus = `DELETE FROM status WHERE id = $1`

	getTotalCount = `SELECT COUNT(id) FROM status`

	getStatus = `SELECT id, created_by, updated_by, created_at, updated_at, deleted_at, updated_at, name, description, active, order_number
				FROM status 
				ORDER BY order_number, created_at, updated_at OFFSET $1 LIMIT $2`

	findByTitleCount = `SELECT COUNT(*)
					FROM status
					WHERE title ILIKE '%' || $1 || '%'`

	findByTitle = `SELECT status_id, author_id, title, content, image_url, category, updated_at, created_at
					FROM status
					WHERE title ILIKE '%' || $1 || '%'
					ORDER BY title, created_at, updated_at
					OFFSET $2 LIMIT $3`
)
