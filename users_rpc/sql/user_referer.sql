CREATE TABLE `user_referer` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    `parent_tree` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '代理树',
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '名称',
    `display_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '展示名',

    `visit_count` BIGINT NOT NULL DEFAULT 0 COMMENT '访问人数',
    `register_count` BIGINT NOT NULL DEFAULT 0 COMMENT '注册人数',
    `first_deposit_count` BIGINT NOT NULL DEFAULT 0 COMMENT '首充人数',

    `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间 (Unix 时间戳)',
    `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间 (Unix 时间戳)',

    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户推广表';