CREATE TABLE public.users (
	"id" serial NOT NULL,
	"name" varchar(255) NULL,
	"email" varchar(190) NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT users_email_unique UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE public.products (
	"id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
	"price" int NOT NULL,
	CONSTRAINT products_name_unique UNIQUE (name),
	CONSTRAINT products_pk PRIMARY KEY (id)
);

CREATE TABLE public.orders (
	"id" serial NOT NULL,
	"user_id" int NOT NULL,
	"order_date" timestamptz NOT NULL,
	"total_amount" int NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);
ALTER TABLE public.orders ADD CONSTRAINT orders_users_fk FOREIGN KEY (user_id) REFERENCES public.users("id");


CREATE TABLE public.orders_products (
	order_id int NOT NULL,
	product_id int NOT NULL
);
ALTER TABLE public.orders_products ADD CONSTRAINT orders_products_products_fk FOREIGN KEY (product_id) REFERENCES public.products("id");
ALTER TABLE public.orders_products ADD CONSTRAINT orders_products_orders_fk FOREIGN KEY (order_id) REFERENCES public.orders("id");