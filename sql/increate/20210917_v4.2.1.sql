CREATE TABLE IF NOT EXISTS `smoke_testing` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `cluster_id` int(11) NOT NULL COMMENT '集群id',
    `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
    `operation_id` varchar(255) NOT NULL COMMENT '操作id',
    `exec_status` int NULL COMMENT '1 进行中,2 正常,3 失败',
    `report_url` varchar(255) NOT NULL COMMENT '报告地址',
    `create_time` timestamp DEFAULT CURRENT_TIMESTAMP NULL  COMMENT '开始时间',
    `end_time` timestamp NULL COMMENT '结束时间',
    PRIMARY KEY(`id`)
    ) COMMENT '冒烟测试记录表' ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `service_health_check` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cluster_id` int(11) NOT NULL COMMENT '集群id',
  `product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `pid` int(11) NOT NULL COMMENT 'product id',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `agent_id` varchar(50) NOT NULL COMMENT 'agent id',
  `sid` varchar(50) NOT NULL COMMENT '主机id',
  `ip` varchar(50) NOT NULL COMMENT '主机ip',
  `script_name` varchar(255) NOT NULL COMMENT '脚本名称',
  `script_name_display` varchar(255) NOT NULL COMMENT '脚本显示名称',
  `auto_exec` tinyint(4) NOT NULL COMMENT '自动执行开关状态',
  `period` varchar(10) NOT NULL COMMENT '执行间隔时间',
  `retries` int(11) DEFAULT NULL COMMENT '执行重试次数',
  `exec_status` int(11) DEFAULT NULL COMMENT '0 未就绪,1 进行中,2 正常,3 失败',
  `error_message` varchar(1000) NOT NULL COMMENT '执行失败的错误信息',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '执行开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '执行结束时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='健康检查记录表';
