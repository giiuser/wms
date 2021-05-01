CREATE TABLE IF NOT EXISTS "collect" (
  "id" bigserial PRIMARY KEY,
  "status" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "collect_table" (
  "id" bigserial PRIMARY KEY,
  "collect_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "qty" integer NOT NULL
);

ALTER TABLE collect_table
    ADD CONSTRAINT fk_rcollect FOREIGN KEY (collect_id) REFERENCES collect (id) ON DELETE CASCADE;