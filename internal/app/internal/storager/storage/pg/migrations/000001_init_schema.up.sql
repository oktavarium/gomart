
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "hash" varchar NOT NULL,
    "salt" varchar NOT NULL,
    "balance" real NOT NULL DEFAULT 0,
    "withdrawn" real NOT NULL DEFAULT 0
);

CREATE TABLE "orders" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "number" varchar NOT NULL,
    "status" varchar NOT NULL DEFAULT ('NEW'),
    "accrual" real NOT NULL DEFAULT 0,
    "uploaded_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdrawals" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "order" varchar NOT NULL,
    "sum" real NOT NULL,
    "processed_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "withdrawals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
