CREATE TABLE IF NOT EXISTS `deploy_instance_update_record` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `update_uuid` char(36) NOT NULL DEFAULT '' COMMENT '更新记录uuid',
  `instance_id` int(11) unsigned NOT NULL COMMENT '实例id',
  `sid` char(36) NOT NULL DEFAULT '' COMMENT '主机 ID (UUID)',
  `ip` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '组件名称',
  `product_name_display` varchar(255) NOT NULL DEFAULT '' COMMENT '组件显示名称',
  `product_version` varchar(255) NOT NULL DEFAULT '' COMMENT '产品版本',
  `group` varchar(255) NOT NULL DEFAULT '' COMMENT '组名称',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
  `service_name_display` varchar(255) NOT NULL DEFAULT '' COMMENT '服务显示名称',
  `service_version` varchar(255) NOT NULL DEFAULT '' COMMENT '服务版本',
  `status` varchar(32) NOT NULL DEFAULT '' COMMENT '实例状态',
  `status_message` varchar(1024) NOT NULL DEFAULT '' COMMENT '状态详细信息',
  `progress` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '进度',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `deploy_uuid` (`update_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=1998 DEFAULT CHARSET=utf8 COMMENT='实例部署记录表';

CREATE TABLE IF NOT EXISTS `deploy_product_update_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id',
  `namespace` varchar(255) NOT NULL COMMENT 'cluster namespace',
  `update_uuid` char(36) NOT NULL COMMENT '更新uuid',
  `parent_product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '组件名称',
  `product_name_display` varchar(255) NOT NULL DEFAULT '' COMMENT '组件显示名称',
  `product_version` varchar(255) NOT NULL DEFAULT '' COMMENT '产品版本',
  `product_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT ',0,, 1, k8s',
  `status` char(32) NOT NULL DEFAULT '' COMMENT '产品状态',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
  `update_start_time` datetime DEFAULT NULL COMMENT 'deploy_start_time',
  `update_end_time` datetime DEFAULT NULL COMMENT 'deploy_end_time',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '部署人id',
  `package_name` varchar(255) NOT NULL COMMENT '补丁包名称',
  `update_dir` varchar(255) NOT NULL COMMENT '目标目录',
  `backup_dir` varchar(255) NOT NULL COMMENT '备份目录',
  `product_id` int(11) NOT NULL COMMENT '产品包id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- 解决 mysql 默认唯一索引不区分大小写 导致配置大小写不敏感问题
alter table deploy_schema_field_modify modify field_path varchar(255) binary default '' not null comment '字段路径';
