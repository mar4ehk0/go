-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.orders_products (
	order_id int NOT NULL,
	product_id int NOT NULL
);
ALTER TABLE public.orders_products ADD CONSTRAINT orders_products_products_fk FOREIGN KEY (product_id) REFERENCES public.products("id");
ALTER TABLE public.orders_products ADD CONSTRAINT orders_products_orders_fk FOREIGN KEY (order_id) REFERENCES public.orders("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS orders_products_products_fk;
DROP INDEX IF EXISTS orders_products_orders_fk;
DROP TABLE public.orders_products;
-- +goose StatementEnd
