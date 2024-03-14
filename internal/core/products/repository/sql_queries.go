package repository

const (
	qGetAllProductByCategoryID = `
		SELECT
			p.product_id,
			p.product_code,
			p.product_name,
			p.is_active,
			p.price,
			p.category_id,
			c.category_name,
			p.stock,
			p.created_at,
			p.created_by,
			p.updated_at,
			p.updated_by
		FROM product p
		INNER JOIN category c ON p.category_id = c.category_id
		WHERE (p.category_id = ? OR ? = 0)
	`
)