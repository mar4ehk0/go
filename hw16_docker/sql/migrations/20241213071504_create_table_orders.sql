-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.orders (
	"id" serial NOT NULL,
	"user_id" int NOT NULL,
	"order_date" timestamptz NOT NULL,
	"total_amount" int NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);
ALTER TABLE public.orders ADD CONSTRAINT orders_users_fk FOREIGN KEY (user_id) REFERENCES public.users("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS orders_users_fk;
DROP TABLE public.orders;
-- +goose StatementEnd
