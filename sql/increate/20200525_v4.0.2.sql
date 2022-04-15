CREATE TABLE IF NOT EXISTS `deploy_cluster_product_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` int(11) NULL DEFAULT -1 COMMENT '产品 id',
  `clusterId` int(11) UNSIGNED NOT NULL COMMENT '集群 id',
  `deploy_uuid` varchar(36) NOT NULL COMMENT '部署uuid',
  `product_parsed` text NOT NULL COMMENT '已经解析的产品信息',
  `status` enum('undeployed','deploying','deployed','deploy fail','undeploying','undeploy fail') NOT NULL DEFAULT 'undeployed' COMMENT '产品状态',
  `alert_recover` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0，不恢复告警，1，恢复告警',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '部署人id',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `deploy_time` datetime COMMENT '部署时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE deploy_product_list ADD COLUMN `product_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '产品包类型,0,传统包, 1, k8s包' AFTER `schema`;
ALTER TABLE deploy_service_ip_list ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;
ALTER TABLE deploy_service_ip_node ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;
ALTER TABLE deploy_schema_field_modify ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;
ALTER TABLE deploy_unchecked_service ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;
ALTER TABLE deploy_product_history ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;
ALTER TABLE deploy_product_history ADD COLUMN `product_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT ',0,, 1, k8s' AFTER `product_version`;
ALTER TABLE deploy_instance_list ADD COLUMN `cluster_id` int(11) unsigned NOT NULL COMMENT '集群id' AFTER `id`;


alter table deploy_service_ip_list drop index `product_service_name`;
alter table deploy_service_ip_list add unique key `cluster_product_service_name` (`cluster_id`,`product_name`,`service_name`);
alter table deploy_service_ip_node drop index `product_service_ip`;
alter table deploy_service_ip_node add unique key `cluster_product_service_ip` (`cluster_id`,`product_name`,`service_name`,`ip`);
alter table deploy_schema_field_modify drop index `names_path`;
alter table deploy_schema_field_modify add unique key `cluster_names_path` (`cluster_id`,`product_name`,`service_name`,`field_path`);
alter table deploy_unchecked_service drop index `pid`;
alter table deploy_unchecked_service add unique key `cluster_pid` (`cluster_id`,`pid`);
alter table deploy_instance_list drop index `pid_service_name`;
alter table deploy_instance_list add unique key `cluster_pid_service_name` (`cluster_id`,`pid`,`service_name`,`sid`);

-- 查询主机集群id, 迁移数据到deploy_cluster_product_rel
select id from deploy_cluster_list;
-- 替换下述cluster_id 为查询到的主机集群id, 假设为1
INSERT INTO `deploy_cluster_product_rel` (`pid`, `clusterId`, `deploy_uuid`, `product_parsed`, `status`, `alert_recover`, `user_id`, `deploy_time`, `update_time`, `create_time`) select id, '1', deploy_uuid, product_parsed, status, alert_recover, user_id, deploy_time, deploy_time, create_time from deploy_product_list where is_current_version=1;
update deploy_service_ip_list set cluster_id=1 where 1=1;
update deploy_service_ip_node set cluster_id=1 where 1=1;
update deploy_schema_field_modify set cluster_id=1 where 1=1;
update deploy_unchecked_service set cluster_id=1 where 1=1;
update deploy_product_history set cluster_id=1 where 1=1;
update deploy_instance_list set cluster_id=1 where 1=1;
