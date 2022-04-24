BEGIN;
INSERT INTO `todo` (title, detail, expire_date) VALUES
('title1', 'detail1', '2022-04-01'),
('title2', 'detail2', '2022-04-02'),
('title3', 'detail3', '2022-04-03');
INSERT INTO `todo` (title, expire_date) VALUES
('title4', '2022-04-04'),
('title5', '2022-04-05'),
('title6', '2022-04-06');
COMMIT;
