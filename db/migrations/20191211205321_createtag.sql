-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `tag` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `project_id` INT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP,
    `created_at` TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE (`name`),
    FOREIGN KEY (project_id)
        references project (id)
        on delete restrict
        on update restrict
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS `tag`;
-- +goose StatementEnd
