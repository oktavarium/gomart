
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user" varchar NOT NULL,
    "hash" varchar NOT NULL,
    "salt" varchar NOT NULL,
    "balance" bigint NOT NULL
);

CREATE TABLE "orders" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "order" varchar NOT NULL,
    "status" varchar NOT NULL,
    "accural" bigint NOT NULL,
    "uploaded_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdrawns" (
    "id" bigserial PRIMARY KEY,
    "order" varchar NOT NULL,
    "sum" bigint NOT NULL,
    "processed_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

