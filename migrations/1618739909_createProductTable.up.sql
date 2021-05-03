CREATE TABLE IF NOT EXISTS "product" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "brand" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);