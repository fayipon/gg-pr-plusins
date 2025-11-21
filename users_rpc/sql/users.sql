CREATE TABLE `users` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    
    `account` VARCHAR(64) NOT NULL COMMENT '用户账号（唯一）',
    `password` VARCHAR(255) NOT NULL COMMENT '密码（bcrypt/argon2）',
    `transaction_password` VARCHAR(255) DEFAULT NULL COMMENT '交易密码（可选）',
    
    `level_id` BIGINT UNSIGNED DEFAULT 0 COMMENT 'VIP 等级 ID',
    `group_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '用户组别 ID',
    
    `email_verified_at` BIGINT DEFAULT NULL COMMENT 'Email 认证时间',
    `mobile_verified_at` BIGINT DEFAULT NULL COMMENT '手机认证时间',
    `kyc_verified_at` BIGINT DEFAULT NULL COMMENT 'KYC 认证时间',
    
    `parent_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '直属上级 ID（代理制度）',
    `parent_tree` VARCHAR(1024) DEFAULT NULL COMMENT '父节点路径，用 / 分隔',
    `depth` INT DEFAULT 0 COMMENT '代理深度（如第 3 层）',
    
    `referer_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '推荐人 ID（可能不同于代理）',
    
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '用户状态 1=正常 0=禁用',
    
    `created_at` BIGINT NOT NULL COMMENT '创建时间（Unix 秒）',
    `updated_at` BIGINT NOT NULL COMMENT '更新时间（Unix 秒）',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_account` (`account`),
    
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_parent_tree` (`parent_tree`),
    KEY `idx_referer_id` (`referer_id`),
    KEY `idx_level_id` (`level_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
