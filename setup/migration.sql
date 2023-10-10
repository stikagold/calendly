CREATE TABLE packs
(
    "id"   integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" text,
    "size" integer
);

INSERT INTO packs(name, size) VALUES ('1x250', 250);
INSERT INTO packs(name, size) VALUES ('1x500', 500);
INSERT INTO packs(name, size) VALUES ('1x1000', 1000);
INSERT INTO packs(name, size) VALUES ('1x2000', 2000);
INSERT INTO packs(name, size) VALUES ('1x5000', 5000)