ALTER TABLE deploy_schema_field_modify MODIFY COLUMN `field` text NOT NULL COMMENT '字段值';
ALTER TABLE deploy_host ADD COLUMN `group` varchar(255) DEFAULT 'default' COMMENT '组信息';

CREATE TABLE IF NOT EXISTS `deploy_unchecked_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) UNSIGNED NOT NULL COMMENT 'product id',
  `unchecked_services` varchar(255) NOT NULL COMMENT '登录密码',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='未勾选服务表';