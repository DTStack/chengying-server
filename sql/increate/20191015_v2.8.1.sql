CREATE TABLE IF NOT EXISTS `deploy_instance_event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `instance_id` int(11) UNSIGNED NOT NULL COMMENT 'instance id',
  `event_type` enum('install', 'uninstall', 'config update', 'start', 'stop', 'exec', 'error', 'unkown') NOT NULL DEFAULT 'unkown' COMMENT '事件类型',
  `content` text NOT NULL DEFAULT '' COMMENT '事件内容',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例事件列表';