package repository

const (
	qRegisterCustomer = `
		INSERT INTO customers (username, password, email, created_at, created_by) 
		VALUES (?, ?, ?, ?, ?) 
		RETURNING *
	`

	qLoginCustomer = `
		SELECT customer_id, username, password, email FROM customers WHERE username = ? and deleted_at IS NULL
	`
)