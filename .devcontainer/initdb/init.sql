CREATE TABLE "payment"(
    "id" SERIAL PRIMARY KEY,
    "method" INTEGER NOT NULL,
    "detail" VARCHAR(255) NOT NULL
);

CREATE TABLE "shipment"(
    "id" SERIAL PRIMARY KEY,
    "method" INTEGER NOT NULL,
    "detail" VARCHAR(255) NOT NULL
);

CREATE TABLE "user"(
    "id" SERIAL PRIMARY KEY,
    "password" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE "member"(
    "user_id" INTEGER NOT NULL REFERENCES "user"("id"),
    "address" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL,
    "phone_num" VARCHAR(255),
    "payment_id" INTEGER REFERENCES "payment"("id"),
    "shipment_id" INTEGER REFERENCES "shipment"("id"),
    PRIMARY KEY("user_id")
);

CREATE TABLE "vendor"(
    "user_id" INTEGER NOT NULL REFERENCES "user"("id"),
    PRIMARY KEY("user_id")
);

CREATE TABLE "administrator"(
    "user_id" INTEGER NOT NULL REFERENCES "user"("id"),
    PRIMARY KEY("user_id")
);

CREATE TABLE "product"(
    "id" SERIAL PRIMARY KEY,
    "price" INTEGER NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "remain" INTEGER NOT NULL,
    "build_time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "vendor_id" INTEGER NOT NULL REFERENCES vendor("user_id")
);

CREATE TABLE "favor"(
    "member_id" INTEGER NOT NULL REFERENCES "member"("user_id"),
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "timedate" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY("member_id", "product_id")
);

CREATE TABLE "cart"(
    "member_id" INTEGER NOT NULL REFERENCES "member"("user_id"),
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "count" INTEGER NOT NULL,
    PRIMARY KEY("member_id", "product_id")
);

CREATE TABLE "order"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "state" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "member_id" INTEGER NOT NULL REFERENCES "member"("user_id"),
    "vendor_id" INTEGER NOT NULL REFERENCES "vendor"("user_id"),
    "payment_id" INTEGER NOT NULL REFERENCES "payment"("id"),
    "shipment_id" INTEGER NOT NULL REFERENCES "shipment"("id")
);

CREATE TABLE "list"(
    "order_id" INTEGER NOT NULL REFERENCES "order"("id"),
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "count" INTEGER NOT NULL,
    "time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY("order_id", "product_id")
);