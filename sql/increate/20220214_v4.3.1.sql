-- 平台配置
CREATE TABLE `sys_config` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `config_name` varchar(45) COLLATE utf8mb4_bin NOT NULL COMMENT '配置名称',
    `config_value` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '配置值',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0: 使用中;1: 已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 平台安全配置.密码传输加密方式
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.login_encrypt','rsa');

-- 平台安全配置.强制用户修改初始密码
-- 0 关闭(default)，1 开启
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.force_reset_password','0');

-- 平台安全配置.密码出错锁定开关
-- 0 关闭(default)，1 开启
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.account_login_lock_switch','0');

-- 平台安全配置.密码出错次数
-- 3(default)
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.account_login_limit_error','3');

-- 平台安全配置.锁定时长
-- 5(default)
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.account_login_lock_time','5');

-- 平台安全配置.自动登出时长
-- 1440(default)
INSERT INTO `sys_config`(config_name,config_value) VALUES ('platformsecurity.account_logout_sleep_time','1440');

-- 用户.初始密码修改状态
ALTER TABLE `user_list` ADD COLUMN `reset_password_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '初始密码修改状态：0 未修改，1 已修改' AFTER `status`;
UPDATE `user_list` SET `reset_password_status` = 1 WHERE 1 = 1;
