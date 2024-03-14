-- public.category definition

-- Drop table

-- DROP TABLE public.category;

CREATE TABLE public.category (
	category_id serial4 NOT NULL,
	category_name varchar NULL,
	is_active bool NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT category_category_name_category_name1_key UNIQUE (category_name) INCLUDE (category_name),
	CONSTRAINT category_pkey PRIMARY KEY (category_id)
);

INSERT INTO category (category_name, is_active, created_at, created_by)
VALUES
    ('Elektronik', true, NOW(), 'System'),
    ('Pakaian', true, NOW(), 'System'),
    ('Perlengkapan Rumah', true, NOW(), 'System'),
    ('Makanan dan Minuman', true, NOW(), 'System'),
    ('Otomotif', true, NOW(), 'System');


-- public.customers definition

-- Drop table

-- DROP TABLE public.customers;

CREATE TABLE public.customers (
	customer_id serial4 NOT NULL,
	username varchar NULL,
	"password" varchar NULL,
	email varchar NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT customers_pkey PRIMARY KEY (customer_id),
	CONSTRAINT customers_username_username1_key UNIQUE (username) INCLUDE (username)
);


-- public.product definition

-- Drop table

-- DROP TABLE public.product;

CREATE TABLE public.product (
	product_id serial4 NOT NULL,
	product_code varchar NULL,
	product_name varchar NULL,
	is_active bool NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	price int8 NULL,
	category_id int4 NULL,
	stock int4 NULL,
	CONSTRAINT product_pkey PRIMARY KEY (product_id),
	CONSTRAINT product_product_code_product_code1_key UNIQUE (product_code) INCLUDE (product_code),
	CONSTRAINT product_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category(category_id)
);

INSERT INTO product (product_code, product_name, is_active, price, category_id, stock, created_at, created_by)
VALUES
    ('PROD_ELEC_1', 'Smartphone', true, 1000000, 1, 50, NOW(), 'System'),
    ('PROD_ELEC_2', 'Laptop', true, 1500000, 1, 30, NOW(), 'System'),
    ('PROD_CLTH_1', 'T-Shirt', true, 20000, 2, 100, NOW(), 'System'),
    ('PROD_CLTH_2', 'Jeans', true, 40000, 2, 80, NOW(), 'System'),
    ('PROD_HOME_1', 'Panci', true, 30000, 3, 50, NOW(), 'System'),
    ('PROD_HOME_2', 'Pengering Rambut', true, 250000, 3, 40, NOW(), 'System'),
    ('PROD_FOOD_1', 'Susu', true, 5000, 4, 200, NOW(), 'System'),
    ('PROD_FOOD_2', 'Roti', true, 2500, 4, 300, NOW(), 'System'),
    ('PROD_AUTO_1', 'Oli Mesin', true, 15000, 5, 100, NOW(), 'System'),
    ('PROD_AUTO_2', 'Ban Mobil', true, 500000, 5, 80, NOW(), 'System');


-- public.shopping_cart definition

-- Drop table

-- DROP TABLE public.shopping_cart;

CREATE TABLE public.shopping_cart (
	shopping_cart_id int4 NOT NULL DEFAULT nextval('shopping_cart_shopping_card_id_seq'::regclass),
	customer_id int4 NULL,
	status_code varchar NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	payment_total int4 NULL,
	CONSTRAINT shopping_cart_pkey PRIMARY KEY (shopping_cart_id),
	CONSTRAINT unique_customer_status UNIQUE (customer_id, status_code),
	CONSTRAINT shopping_cart_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customers(customer_id)
);


-- public.shopping_cart_product definition

-- Drop table

-- DROP TABLE public.shopping_cart_product;

CREATE TABLE public.shopping_cart_product (
	shopping_cart_product_id int4 NOT NULL DEFAULT nextval('shopping_cart_product_shopping_card_product_id_seq'::regclass),
	product_id int4 NULL,
	shopping_cart_id int4 NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	quantity int4 NULL,
	sub_total int8 NULL,
	price int8 NULL,
	CONSTRAINT shopping_cart_product_pkey PRIMARY KEY (shopping_cart_product_id),
	CONSTRAINT unique_shopping_cart UNIQUE (product_id, shopping_cart_id),
	CONSTRAINT shopping_cart_product_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(product_id),
	CONSTRAINT shopping_cart_product_shopping_card_id_fkey FOREIGN KEY (shopping_cart_id) REFERENCES public.shopping_cart(shopping_cart_id)
);


-- public.payment_transaction definition

-- Drop table

-- DROP TABLE public.payment_transaction;

CREATE TABLE public.payment_transaction (
	payment_transaction_id serial4 NOT NULL,
	shopping_cart_id int4 NULL,
	payment_total int4 NULL,
	created_by varchar NULL,
	created_at timestamptz NULL,
	updated_by varchar NULL,
	updated_at timestamptz NULL,
	deleted_by varchar NULL,
	deleted_at timestamptz NULL,
	payment_type varchar NULL,
	CONSTRAINT payment_transaction_pkey PRIMARY KEY (payment_transaction_id),
	CONSTRAINT payment_transaction_shopping_card_id_fkey FOREIGN KEY (shopping_cart_id) REFERENCES public.shopping_cart(shopping_cart_id)
);