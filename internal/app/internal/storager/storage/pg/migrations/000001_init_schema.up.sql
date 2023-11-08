
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user" varchar NOT NULL,
    "hash" varchar NOT NULL,
    "salt" varchar NOT NULL,
    "balance" bigint NOT NULL DEFAULT(0)
);

CREATE TABLE "orders" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "number" varchar NOT NULL,
    "status" varchar NOT NULL DEFAULT ('NEW'),
    "accural" bigint,
    "uploaded_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdrawns" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "order_id" varchar NOT NULL,
    "sum" bigint NOT NULL,
    "processed_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "withdrawns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "withdrawns" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

