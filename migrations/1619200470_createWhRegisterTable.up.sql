CREATE TABLE IF NOT EXISTS "wh_register" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL,
  "document_id" bigint NOT NULL,
  "document_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);