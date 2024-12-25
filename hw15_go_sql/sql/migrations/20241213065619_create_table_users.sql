-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users (
	"id" serial NOT NULL,
	"name" varchar(255) NULL,
	"email" varchar(190) NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT users_email_unique UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
