CREATE TABLE IF NOT EXISTS "receipt" (
  "id" bigserial PRIMARY KEY,
  "status" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "receipt_table" (
  "id" bigserial PRIMARY KEY,
  "receipt_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL
);

ALTER TABLE receipt_table
    ADD CONSTRAINT fk_receipt FOREIGN KEY (receipt_id) REFERENCES receipt (id) ON DELETE CASCADE;
