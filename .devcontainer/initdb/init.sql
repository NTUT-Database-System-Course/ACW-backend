-- init all tables
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
    "description" VARCHAR(4095) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "remain" INTEGER NOT NULL,
    "disability" BOOLEAN NOT NULL,
    "image_url" VARCHAR(255),
    "build_time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "vendor_id" INTEGER NOT NULL REFERENCES vendor("user_id")
);

CREATE TABLE "favor"(
    "member_id" INTEGER NOT NULL REFERENCES "member"("user_id"),
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
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
    "phone_num" VARCHAR(255) NOT NULL,
    "member_id" INTEGER NOT NULL REFERENCES "member"("user_id"),
    "payment_method" INTEGER NOT NULL,
    "shipment_method" INTEGER NOT NULL
);

CREATE TABLE "list"(
    "order_id" INTEGER NOT NULL REFERENCES "order"("id"),
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "count" INTEGER NOT NULL,
    "time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY("order_id", "product_id")
);

CREATE TABLE "tag"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "type" INTEGER NOT NULL
);

CREATE TABLE "own"(
    "product_id" INTEGER NOT NULL REFERENCES "product"("id"),
    "tag_id" INTEGER NOT NULL REFERENCES "tag"("id"),
    PRIMARY KEY("product_id", "tag_id")
);

-- init all default values
INSERT INTO "user" ("password", "name", "username") VALUES 
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', 'member', 'member'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', 'admin', 'admin'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '鬼滅之刃 代理商', 'vendor1'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '獵人Hunter 代理商', 'vendor2'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '出租女友 代理商', 'vendor3'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '間諜家家酒 代理商', 'vendor4'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '進擊的巨人 代理商', 'vendor5'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '刀劍神域 代理商', 'vendor6'),
('$2a$10$QT7PBe0i.0EftDfL.fGMb.CpN5htUTCLx/vuxvywi3y9qnVwhVqeO', '葬送的芙莉蓮 代理商', 'vendor7');

INSERT INTO "member"("user_id", "email") VALUES
(1, 'member@example.com');

INSERT INTO "administrator"("user_id") VALUES
(2);

INSERT INTO "vendor"("user_id") VALUES
(3);

INSERT INTO "tag"("name", "type") VALUES
('鬼滅之刃', 0),
('獵人Hunter', 0),
('出租女友', 0),
('間諜家家酒', 0),
('進擊的巨人', 0),
('刀劍神域', 0),
('葬送的芙莉蓮', 0),
('公仔', 1),
('徽章', 1),
('資料夾', 1),
('鑰匙圈', 1),
('馬克杯', 1),
('掛畫', 1),
('滑鼠墊', 1),
('雨傘', 1),
('衣服', 1);

-- TODO : add more products later
INSERT INTO "product"("price", "description", "name", "remain", "disability", "image_url", "build_time", "vendor_id") VALUES
(590, '此公仔為一名持劍角色，背景為黑暗森林，特效呈現出冰霜般的波動感，動態感強烈，細節精美。', 'SE公仔 柱稽古篇-時透無一郎', 100, false, 'https://imgur.com/a/YqzBXeK', '2024-12-25 00:00:00', 3),
(590, '', 'SE公仔 柱稽古篇-伊黑小芭內', 50, false, 'https://imgur.com/a/G1U1Slh', '2024-12-25 00:01:20', 3),
(590, '', 'SE公仔 柱稽古篇-不死川實彌', 50, false, 'https://imgur.com/a/ujKRgpo', '2024-12-25 00:03:10', 3),
(590, '', 'MAXIMATIC 宇隨天元', 40, false, 'https://imgur.com/a/xnM6FYT', '2024-12-26 10:00:10', 3),
(590, '', '鬼之裝EX 憎珀天', 75, false, 'https://imgur.com/a/MC1eq1G', '2024-12-26 10:20:10', 3),
(590, '', '鬼之裝EX 妓夫太郎', 75, false, 'https://imgur.com/a/5t1npHX', '2024-12-26 10:22:10', 3),
(590, '', '鬼之裝EX 墮姬', 75, false, 'https://imgur.com/a/OEwsCkX', '2024-12-24 09:24:50', 3),
(150, '', '第一彈徽章-竈門炭治郎', 150, false, 'https://imgur.com/a/2AXKTAN', '2024-12-21 00:01:00', 3),
(150, '', '第一彈徽章-竈門彌豆子', 150, false, 'https://imgur.com/a/MdkUQaG', '2024-12-21 00:02:00', 3),
(150, '', '第一彈徽章-我妻善逸', 150, false, 'https://imgur.com/a/cDHjluP', '2024-12-21 00:03:00', 3),
(150, '', '第一彈徽章-嘴平伊之助', 150, false, 'https://imgur.com/a/I4GCSbb', '2024-12-21 00:04:00', 3),
(150, '', '第一彈徽章-富岡義勇', 150, false, 'https://imgur.com/a/5x3usI9', '2024-12-21 00:05:00', 3),
(150, '', '第一彈徽章-煉獄杏壽郎', 150, false, 'https://imgur.com/a/Gzs7HlU', '2024-12-21 00:06:00', 3),
(150, '', '第一彈徽章-胡蝶忍', 150, false, 'https://imgur.com/a/0jPjywp', '2024-12-21 00:07:00', 3),
(150, '', '第一彈徽章-栗花落香奈乎', 150, false, 'https://imgur.com/a/zB599tf', '2024-12-21 00:08:00', 3),
(150, '', 'A4資料夾-竈門炭治郎', 200, false, 'https://imgur.com/a/NVj8Ept', '2024-12-21 01:00:00', 3),
(150, '', 'A4資料夾-竈門彌豆子', 200, false, 'https://imgur.com/a/zqW8VkK', '2024-12-21 01:01:00', 3),
(150, '', 'A4資料夾-我妻善逸', 200, false, 'https://imgur.com/a/RXwDWNE', '2024-12-21 01:02:00', 3),
(150, '', 'A4資料夾-嘴平伊之助', 200, false, 'https://imgur.com/a/TWgpI2C', '2024-12-21 01:03:00', 3),
(150, '', 'A4資料夾-富岡義勇', 200, false, 'https://imgur.com/a/i0GQ3i0', '2024-12-21 01:04:00', 3),
(150, '', 'A4資料夾-煉獄杏壽郎', 200, false, 'https://imgur.com/a/wi99Swa', '2024-12-21 01:05:00', 3),
(150, '', 'A4資料夾-胡蝶忍', 200, false, 'https://imgur.com/a/j1U9ugn', '2024-12-21 01:06:00', 3),
(150, '', 'A4資料夾-栗花落香奈乎', 200, false, 'https://imgur.com/a/ZIfMyPD', '2024-12-21 01:07:00', 3),
(150, '', 'A4資料夾-宇隨天元', 200, false, 'https://imgur.com/a/UykfVee', '2024-12-21 01:08:00', 3),
(150, '', 'A4資料夾-時透無一郎', 200, false, 'https://imgur.com/a/LKAdlkt', '2024-12-21 01:09:00', 3),
(150, '', 'A4資料夾-甘露寺蜜璃', 200, false, 'https://imgur.com/a/4JYHHde', '2024-12-21 01:10:00', 3),
(135, '', 'Q版鑰匙圈-竈門炭治郎', 80, false, 'https://imgur.com/a/PZPCi7M', '2024-12-21 02:00:00', 3),
(135, '', 'Q版鑰匙圈-我妻善逸', 80, false, 'https://imgur.com/a/dVGYljl', '2024-12-21 02:01:00', 3),
(135, '', 'Q版鑰匙圈-嘴平伊之助', 80, false, 'https://imgur.com/a/DFsN26d', '2024-12-21 02:02:00', 3),
(135, '', 'Q版鑰匙圈-富岡義勇', 80, false, 'https://imgur.com/a/YfF4Kv5', '2024-12-21 02:03:00', 3);

INSERT INTO "own"("product_id", "tag_id") VALUES
(1, 1),
(1, 8);