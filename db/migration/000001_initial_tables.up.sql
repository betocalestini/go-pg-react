CREATE TABLE IF NOT EXISTS  "users" (
    "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ, 
    "deleted_at" TIMESTAMPTZ 
);


CREATE TABLE IF NOT EXISTS "categories" (
    "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
    "user_id" INTEGER NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ, 
    "deleted_at" TIMESTAMPTZ
    );
ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


CREATE TABLE IF NOT EXISTS "accounts" (
    "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
    "user_id" INTEGER NOT NULL,
    "category_id" INTEGER NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "value" INTEGER NOT NULL,
    "date" DATE NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ, 
    "deleted_at" TIMESTAMPTZ
    );
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");