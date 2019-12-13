-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `member_tag` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `project_id` INT NOT NULL,
    `member_id` INT NOT NULL,
    `tag_id` INT NOT NULL,
    `weight` INT,
    `updated_at` TIMESTAMP,
    `created_at` TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (project_id)
        references project (id)
        on delete restrict
        on update restrict,
    FOREIGN KEY (member_id)
        references member (id)
        on delete restrict
        on update restrict
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS `member_tag`;
-- +goose StatementEnd
