CREATE TABLE `task_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '名称',
  `describe` varchar(256) NOT NULL COMMENT '描述',
  `spec` varchar(64) NOT NULL COMMENT 'cron表达式',
  `status` int NULL DEFAULT '0' COMMENT '定时状态: 0 关闭,1 开启',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常 1逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务列表';

CREATE TABLE `task_host` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` int(11) NOT NULL COMMENT '任务id',
  `host_id` int(11) NOT NULL COMMENT '主机id',
  PRIMARY KEY (`id`),
  KEY `IDX_task_host_task_id` (`task_id`),
  KEY `IDX_task_host_host_id` (`host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务主机关联表';

CREATE TABLE `task_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `task_id` varchar(11) NOT NULL COMMENT '任务id',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `spec` varchar(64) NOT NULL COMMENT 'cron表达式',
  `ip` varchar(255) NOT NULL COMMENT '主机ip',
  `operation_id` varchar(255) NOT NULL COMMENT '操作id',
  `command` varchar(1024) NOT NULL COMMENT '执行命令',
  `exec_type` int NULL COMMENT '执行类型：0 定时执行，1 手动执行',
  `exec_status` int NULL COMMENT '执行状态: 0 未运行,1 运行中,2 正常,3 异常',
  `exec_result` longtext NOT NULL COMMENT '执行结果',
  `start_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`),
  KEY `IDX_task_log_operation_id` (`operation_id`),
  KEY `IDX_task_log_task_id` (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务日志表';

CREATE TABLE `deploy_backup_history` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
 `cluster_id` int(11) NOT NULL COMMENT '集群id',
 `db_name` varchar(255) NOT NULL COMMENT '数据库名称',
 `backup_sql` varchar(255) NOT NULL COMMENT '备份文件名称',
 `product_name` varchar(255) DEFAULT NULL COMMENT '触发此次备份的产品包名称',
 `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='备份历史记录表';

CREATE TABLE `deploy_upgrade_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `cluster_id` int(11) NOT NULL COMMENT '集群id',
  `product_name` varchar(64) NOT NULL COMMENT '产品名称',
  `source_version` varchar(64) NOT NULL COMMENT '源版本',
  `target_version` varchar(64) NOT NULL COMMENT '目标版本',
  `backup_name` varchar(64) DEFAULT '' COMMENT '备份名称，值为当前时间',
  `source_service_ip` text COMMENT '源版本服务编排信息',
  `source_config` text COMMENT '源版本服务配置信息',
  `source_multi_config` text COMMENT '源版本服务多配置信息',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `type` tinyint(4) NOT NULL COMMENT '类型，0为升级，1为回滚',
  `backup_sql` text COMMENT '备份SQL文件',
  `is_deleted` tinyint(4) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='升级历史记录表';
