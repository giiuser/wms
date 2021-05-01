CREATE TABLE IF NOT EXISTS "waybill" (
  "id" bigserial PRIMARY KEY,
  "status" smallint NOT NULL DEFAULT 0,
  "document_id" bigint NOT NULL,
  "document_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "waybill_table" (
  "id" bigserial PRIMARY KEY,
  "waybill_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL
);

ALTER TABLE waybill_table
    ADD CONSTRAINT fk_waybill FOREIGN KEY (waybill_id) REFERENCES waybill (id) ON DELETE CASCADE;