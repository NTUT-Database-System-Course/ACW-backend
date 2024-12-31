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
    "announcement" VARCHAR(4095),
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

INSERT INTO "vendor"("user_id", "announcement") VALUES
(3, '⭐️以下注意事項請於購買前詳細閱讀⭐️<br/>1.本預購商品係依據您的需求向廠商訂購，恕不接受預購後要求取消訂單或退款退貨。<br/>2.本商品除瑕疵品協助換貨外，恕不接受任何理由取消訂單、退款及退貨。<br/>3.預購商品建議統一將同年同月份併結帳，不同月份的預購商品恕分開計單結帳，需等候全部商品到貨後才會一起寄出，可能會造成您更久的等待。<br/>4.不同月份的預購商品如需拆單寄出，需支付運費 NT.100。<br/>5.商品預購完成後，如遇缺斷貨、廠商法生產或無法掌控之因素等，我們必須取消您的訂單時，預購款項將全額退款。<br/>6.商品出貨配送過程中難免碰撞，在不影響商品狀態下，不列為瑕疵 + 十分在意外盒完整美者，請勿下單，恕不接受因外盒碰撞申請換貨。<br/>7.預購商品到貨月份均為廠商預估告知，實際到貨日期涉及諸多因素可能會與預估月份不同，訂購人無法以延遲到貨為由取消訂單、退款。<br/>8.如廠商在產延遲或其他因素延遲到貨，因無法逐一通知，可多利用客服留言與我們詢問。<br/>9.公仔玩具、塗裝多採手工繪製 + 批次生產組裝，製作過程中難免產生細微塗裝、造型或組裝差異，差異範圍皆為 ± 小於 2cm 風險合理範圍。<br/>10.本預購商品，圖片可能為產品模擬原型，實際收到商品以實物為主。<br/>11.首抽、盒玩類商品，屬非搶奪性驚喜盲盒，無法因個人喜好挑選款式，請慎先理解此產品販售特性，再行下單！<br/>⭐️下單代表同意以上敘述。<br/>⭐️如有任何疑問請下單前詢問。<br/>'),
(4, '⭐️以下注意事項請於購買前詳細閱讀⭐️<br/>1.本預購商品係依據您的需求向廠商訂購，恕不接受預購後要求取消訂單或退款退貨。<br/>2.本商品除瑕疵品協助換貨外，恕不接受任何理由取消訂單、退款及退貨。<br/>3.預購商品建議統一將同年同月份併結帳，不同月份的預購商品恕分開計單結帳，需等候全部商品到貨後才會一起寄出，可能會造成您更久的等待。<br/>4.不同月份的預購商品如需拆單寄出，需支付運費 NT.100。<br/>5.商品預購完成後，如遇缺斷貨、廠商法生產或無法掌控之因素等，我們必須取消您的訂單時，預購款項將全額退款。<br/>6.商品出貨配送過程中難免碰撞，在不影響商品狀態下，不列為瑕疵 + 十分在意外盒完整美者，請勿下單，恕不接受因外盒碰撞申請換貨。<br/>7.預購商品到貨月份均為廠商預估告知，實際到貨日期涉及諸多因素可能會與預估月份不同，訂購人無法以延遲到貨為由取消訂單、退款。<br/>8.如廠商在產延遲或其他因素延遲到貨，因無法逐一通知，可多利用客服留言與我們詢問。<br/>9.公仔玩具、塗裝多採手工繪製 + 批次生產組裝，製作過程中難免產生細微塗裝、造型或組裝差異，差異範圍皆為 ± 小於 2cm 風險合理範圍。<br/>10.本預購商品，圖片可能為產品模擬原型，實際收到商品以實物為主。<br/>11.首抽、盒玩類商品，屬非搶奪性驚喜盲盒，無法因個人喜好挑選款式，請慎先理解此產品販售特性，再行下單！<br/>⭐️下單代表同意以上敘述。<br/>⭐️如有任何疑問請下單前詢問。<br/>'),
(5, '每周二公休<br/>⭐️訂購or現貨注意事項⭐️<br/>1.標題 月份 是日本明年發售時間。<br/>2.免訂金商品不接受任何理由取消交易。<br/>3.對外盒要求完美者請勿下單，請至店內選購。<br/>4.不接受下標店內取貨，如有需要請您到店內填寫訂單。<br/>5.免訂金預每項商品只限購一組，如需要購買多組請詢問付訂金。<br/>6.商品因大量生產，塗裝多少有些（溢色或瑕疵），能接受者再下單。<br/>7.預購商品是從日本進口，進口商品有可能延期，能接受者再進行預購。<br/>8.預購商品如遇到日本(砍量，數量不足時)，將依訂金金額優先分配(付清.訂金.取付)能接受者在進行預購。<br/>9.預購商品如遇到日本延期，廠商延遲，海關抽檢請耐心等待，不接受退訂，能接受者在進行預購。<br/>10.因外幣商品價格是以發售時的匯率進行計算，如進貨時匯率波動過大，請配合修改訂單金額，能接受在進行預購。<br/>⭐️下單代表同意以上敘述。<br/>⭐️如有任何疑問請下單前詢問。<br/>⭐️員林卡通漫畫屋感謝您的選購⭐️<br/>'),
(6, '⭐購物須知⭐<br/>1》販售動漫商品、一番賞皆為正版。<br/>2》商品皆為為工廠大量製造，脫線、刮痕、塗裝瑕疵、溢色等情形為正常現象，能接受再下單。<br/>3》運輸關係盒況可能會受到擠壓，在意外盒者請審慎評估。<br/>4》急單請選擇宅配，超商時間會比較久!!<br/>5》賣場不提供殺價❌保留服務❌<br/>6》包裹拆箱請錄影，保障您我權益^^<br/>7》外包裝袋、配件等任何問題請先聊聊詢問。<br/>'),
(7, '歡迎光臨:《進擊的巨人》<br/>★【出貨注意】<br/>下單後正常2-3出貨，特殊情況下可能需要4-5天；物流需要5-6個工作日左右才能到您手上哦！<br/>★【購物說明】<br/>1.不同批到貨的商品尺寸及顏色，皆會有些許色差，還請各位買家見諒喔！<br/>2.不同批到貨的商品尺寸為手工量測，可能會有1-3公分的誤差，請以實物為準喔！<br/>3.請確認好規格再下單，下單後若要更改訂單內容，請自行取消再重新下單！<br/>4.不要在聊聊和訂單備註顏色款式，下單前請確認顏色款式，並下單正確！<br/>5.請確認好規格再下單,若有額外需求請在“備註”中通知,訂單建立後,不可更改!<br/>6.商品因拍攝燈光原因和電腦顯色略有不同，商品圖可以參考，請以實際商品為準！<br/>出貨時間:<br/>周一--周五:08:00~ 12:00 (周六周日 休息倉庫不出貨)<br/>'),
(8, '【🏪歡迎光臨🏪】<br/>📣📣📣歡迎光臨~ 賣場均為現貨喔📣📣📣<br/>【關於訂單顯示出貨時間長的說明】<br/>👌👌訂單添加質保服務後顯示預售或者出貨時間長請水水勿擔心~ <br/>🚚🚚倉庫正常約3個工作天左右會出貨(不含例假日)正常下標5-7天左右送達。<br/>🚚🚚訂單顯示的出貨時間無參考價值。 請水水耐心等待~~ <br/>📢📢本店暫時不接急單的哦急要,要出國,時間趕的粉粉們請務必慎拍並提前告知,不接急單,請水水們預留足夠的時間哦 <br/>🙏🙏取貨後如果有問題,請及時與我們聯繫,如收到的貨物存在漏發和破損等情況,請給我們拍照留言,千萬不要著急給我們差評,我們一定會第一時間給您滿意的答復和解決方案,謝謝水水的理解~~ <br/>【😊😊友友們請注意😊😊】<br/>1.關注賣場可以領取更多優惠,水水們看到不錯的商品可以自助下標,倉庫可以先安排出貨 <br/>2.下單後不取件者,擾亂正常秩序,壹律通過法律渠道追究責任,請親們自重! 訂單提示的完成付款時間不是取貨時間,取貨時間請以收到的簡訊為準哦,切記以簡訊為準。 <br/>3.超商取貨有尺寸限制,購買體積較大的產品,建議採用宅配方式寄送,或自行分成兩筆訂單! <br/>4.正常接單發貨,賣場是小本低利潤賣家,同業惡意競爭批評或高標準常給惡評買家請繞道喔~ 每個⭐⭐⭐⭐⭐星好評對我們來說都很重要,希望我們真誠的服務都會贏得您的認可! 最後祝你們購物愉快! <br/>'),
(9, '⭐購物須知⭐<br/>1》販售動漫商品、一番賞皆為正版。<br/>2》商品皆為為工廠大量製造，脫線、刮痕、塗裝瑕疵、溢色等情形為正常現象，能接受再下單。<br/>3》運輸關係盒況可能會受到擠壓，在意外盒者請審慎評估。<br/>4》急單請選擇4大超商，蝦皮店到店時間會比較久!!<br/>5》賣場不提供保留服務❌<br/>6》包裹拆箱請錄影，保障您我權益^^<br/>⭐️下單代表同意以上敘述。<br/>⭐️如有任何疑問請下單前詢問。<br/>希望我們真誠的服務都會贏得您的認可! 最後祝你們購物愉快! <br/>');

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
(590, '此公仔為一名持劍角色，背景為黑暗森林，特效呈現出冰霜般的波動感，動態感強烈，細節精美。', 'SE公仔 柱稽古篇-時透無一郎', 100, false, 'https://diz36nn4q02zr.cloudfront.net/webapi/imagesV3/Original/SalePage/10365660/0/638689338251170000?v=1', '2024-12-25 00:00:00', 3);

INSERT INTO "own"("product_id", "tag_id") VALUES
(1, 1),
(1, 8);