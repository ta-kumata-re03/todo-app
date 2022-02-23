-- +migrate Up
CREATE TABLE `todo` (
  `id` INTEGER NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(30) NOT NULL,
  `detail` VARCHAR(255),
  `expire_date` TEXT NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS todo;
