CREATE TABLE IF NOT EXISTS "allocation" (
  "id" bigserial PRIMARY KEY,
  "status" smallint NOT NULL DEFAULT 0,
  "document_id" bigint NOT NULL,
  "document_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "allocation_table" (
  "id" bigserial PRIMARY KEY,
  "allocation_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL,
  "cell_id" integer NOT NULL
);

ALTER TABLE allocation_table
    ADD CONSTRAINT fk_allocation FOREIGN KEY (allocation_id) REFERENCES allocation (id) ON DELETE CASCADE;