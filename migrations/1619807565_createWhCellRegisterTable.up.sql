CREATE TABLE IF NOT EXISTS "wh_cell_register" (
  "id" bigserial PRIMARY KEY,
  "cell_id" integer NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL,
  "document_id" bigint NOT NULL,
  "document_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);