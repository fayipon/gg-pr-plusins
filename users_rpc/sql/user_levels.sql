CREATE TABLE `user_levels` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    
    `name` VARCHAR(100) NOT NULL COMMENT '等级代号（内部使用，唯一）',
    `display_name` VARCHAR(100) NOT NULL COMMENT '显示名称（用于 UI）',

    `setting` JSON NOT NULL COMMENT '等级配置（返点比例、升级条件等）',

    `created_at` BIGINT NOT NULL COMMENT '创建时间（Unix 秒）',
    `updated_at` BIGINT NOT NULL COMMENT '更新时间（Unix 秒）',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
