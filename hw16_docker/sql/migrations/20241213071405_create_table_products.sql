-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.products (
	"id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
	"price" int NOT NULL,
	CONSTRAINT products_name_unique UNIQUE (name),
	CONSTRAINT products_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.products
-- +goose StatementEnd
