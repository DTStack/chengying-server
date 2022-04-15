## deploy_notify_event添加product_stopped字段
ALTER TABLE deploy_notify_event ADD COLUMN product_stopped TINYINT DEFAULT '0' COMMENT '组件是否停止';

CREATE TABLE `deploy_switch_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) NOT NULL COMMENT '开关名称',
  `product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `status` varchar(64) NOT NULL COMMENT '状态',
  `status_message` varchar(1024) DEFAULT NULL COMMENT '状态详细信息',
  `progress` tinyint(3) DEFAULT NULL COMMENT '进度',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `cluster_id` int(11) NOT NULL COMMENT '集群id',
  `switch_type` varchar(16) NOT NULL COMMENT '开关操作类型，on/off',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='开关记录表';

alter table deploy_cluster_product_rel modify product_parsed LongText not null;
