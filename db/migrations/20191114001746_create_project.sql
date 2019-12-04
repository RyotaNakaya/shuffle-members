-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `project` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `updated_at` TIMESTAMP,
    `created_at` TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS `project`;
-- +goose StatementEnd
