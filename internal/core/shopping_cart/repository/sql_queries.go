package repository

const (
	qAddProductShoppingCart = `
		INSERT INTO shopping_cart (customer_id, status_code, payment_total, created_at, created_by) 
		VALUES (?, ?, ?, ?, ?) 
		ON CONFLICT (customer_id, status_code) DO UPDATE SET payment_total = shopping_cart.payment_total + EXCLUDED.payment_total
		RETURNING *
	`

	qAddProductShoppingCartProduct = `
		INSERT INTO shopping_cart_product (shopping_cart_id, product_id, quantity, sub_total, price, created_at, created_by) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT (shopping_cart_id, product_id) DO UPDATE SET quantity = shopping_cart_product.quantity + EXCLUDED.quantity, sub_total = shopping_cart_product.sub_total + EXCLUDED.sub_total 
		RETURNING * 
	`

	qGetShoppingCart = `
		SELECT * FROM shopping_cart WHERE customer_id = ? and status_code = 'cart'
	`

	qGetShoppingCartProduct = `
		SELECT 
			scp.shopping_cart_product_id,
			scp.shopping_cart_id,
			scp.product_id,
			p.product_name,
			scp.price,
			scp.quantity,
			c.category_name,
			scp.sub_total,
			scp.created_at,
			scp.created_by,
			scp.updated_at,
			scp.updated_by
		FROM shopping_cart_product scp 
		INNER JOIN product p ON scp.product_id = p.product_id
		INNER JOIN category c ON p.category_id = c.category_id
		WHERE scp.shopping_cart_id = ? and scp.deleted_at IS NULL
	`

	qDeleteShoppingCartProduct = `
		DELETE FROM shopping_cart_product
		WHERE shopping_cart_product_id = ?
	`

	qDecreasePaymentTotal = `
		UPDATE shopping_cart
		SET payment_total = payment_total - (SELECT sub_total FROM shopping_cart_product WHERE shopping_cart_product_id = ?)
		WHERE shopping_cart_id = (SELECT shopping_cart_id FROM shopping_cart_product WHERE shopping_cart_product_id = ?)
	`

	qPaymentTransaction = `
		INSERT INTO payment_transaction (shopping_cart_id, payment_total, payment_type, created_at, created_by)
		SELECT shopping_cart_id, payment_total, 'cash', ?, ?
		FROM shopping_cart
		WHERE customer_id = ? and status_code = 'cart'
		RETURNING *
	`

	qCheckoutShoppingCart = `
		UPDATE shopping_cart
		SET status_code = 'checkout', updated_at = ?, updated_by = ?
		WHERE customer_id = ? and status_code = 'cart'
		RETURNING *
	`

	qSelectShoppingCartProduct = `
		SELECT * FROM shopping_cart_product WHERE shopping_cart_id = ?
	`

	qDecreaseStockProduct = `
		UPDATE product
		SET stock = stock - ?
		WHERE product_id = ?
		RETURNING *
	`
)