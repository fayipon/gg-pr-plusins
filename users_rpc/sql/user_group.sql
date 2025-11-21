CREATE TABLE `user_groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL COMMENT '代号（内部使用，唯一）',
  `display_name` varchar(128) NOT NULL COMMENT '显示名称（外部显示）',
  `setting` text NOT NULL COMMENT 'JSON 配置',
  `created_at` bigint NOT NULL COMMENT '创建时间 (unix)',
  `updated_at` bigint NOT NULL COMMENT '更新时间 (unix)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
