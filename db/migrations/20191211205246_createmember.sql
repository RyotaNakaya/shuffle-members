-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `member` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `project_id` INT NOT NULL,
    `employee_number` INT,
    `name` VARCHAR(50) NOT NULL,
    `name_kana` VARCHAR(50) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `shuffle_count` INT,
    `updated_at` TIMESTAMP,
    `created_at` TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE (`email`),
    FOREIGN KEY (project_id)
        references project (id)
        on delete restrict
        on update restrict
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS `member`;
-- +goose StatementEnd
