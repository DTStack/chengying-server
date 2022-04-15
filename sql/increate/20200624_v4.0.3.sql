CREATE TABLE IF NOT EXISTS `deploy_kube_base_product_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` varchar(255) DEFAULT '' COMMENT 'pid',
  `clusterId` int(11) unsigned NOT NULL COMMENT 'cluster id',
  `namespace` varchar(255) DEFAULT '' COMMENT 'k8s部署模式，命名空间',
  `base_clusterId` varchar(36) NOT NULL COMMENT '依赖集群',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `pid_clusterId_name_base` (`pid`, `clusterId`, `namespace`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- CREATE TABLE IF NOT EXISTS `deploy_kube_product_product_rel` (
--   `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
--   `pid` varchar(255) DEFAULT '' COMMENT 'pid',
--   `clusterId` int(11) unsigned NOT NULL COMMENT 'cluster id',
--   `namespace` varchar(255) DEFAULT '' COMMENT 'k8s部署模式，命名空间',
--   `base_product_name` varchar(36) NOT NULL COMMENT '依赖产品',
--   `deploy_id` varchar(36) NOT NULL COMMENT '依赖组件部署id',
--   `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
--   `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
--   `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `pid_clusterId_name_base` (`pid`, `service_name`, `sid`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `deploy_kube_product_lock` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` varchar(255) DEFAULT '' COMMENT 'pid',
  `clusterId` int(11) unsigned NOT NULL COMMENT 'cluster id',
  `namespace` varchar(255) DEFAULT '' COMMENT 'k8s部署模式，命名空间',
  `is_deploy` int(11) NOT NULL DEFAULT 0 COMMENT '1: deploying | 0: not deploy',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `deploy_cluster_kube_pod_list` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` int(11) UNSIGNED NOT NULL COMMENT 'product id',
  `clusterId` int(11) unsigned NOT NULL COMMENT 'cluster id',
  `namespace` varchar(255) DEFAULT '' COMMENT '命名空间',
  `product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_version` varchar(255) NOT NULL DEFAULT '' COMMENT '产品版本',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `service_version` varchar(255) NOT NULL DEFAULT '' COMMENT '服务版本',
  `pod_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'pod uid',
  `pod_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'pod names',
  `pod_key` varchar(255) NOT NULL DEFAULT '' COMMENT 'pod informer cache key',
  `self_link` text NOT NULL DEFAULT '' COMMENT 'pod self link',
  `host_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip',
  `pod_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip',
  `phase` enum('Pending', 'Running', 'Succeeded', 'Failed', 'Unknown') NOT NULL DEFAULT 'Pending' COMMENT 'Pod状态',
  `message` varchar(1024) NOT NULL DEFAULT '' COMMENT '状态详细信息',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_kube_service_list` (
 `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` int(11) UNSIGNED NOT NULL COMMENT 'product id',
  `clusterId` int(11) unsigned NOT NULL COMMENT 'cluster id',
  `namespace` varchar(255) DEFAULT '' COMMENT '命名空间',
  `product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_version` varchar(255) NOT NULL DEFAULT '' COMMENT '产品版本',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `service_version` varchar(255) NOT NULL DEFAULT '' COMMENT '服务版本',
  `cluster_ip` varchar(255) NOT NULL DEFAULT '' COMMENT 'cluster ip of the service',
  `type` enum('ClusterIP', 'NodePort', 'LoadBalancer', 'ExternalName') NOT NULL DEFAULT 'ClusterIP' COMMENT 'service type',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE deploy_cluster_image_store ADD COLUMN `is_default` tinyint(1) NOT NULL COMMENT '1:default image store' AFTER `clusterId`;
ALTER TABLE deploy_cluster_product_rel ADD COLUMN `namespace` varchar(255) NOT NULL COMMENT 'cluster namespace' AFTER `clusterId`;
ALTER TABLE deploy_product_history ADD COLUMN `namespace` varchar(255) NOT NULL COMMENT 'cluster namespace' AFTER `cluster_id`;
ALTER TABLE deploy_instance_list ADD COLUMN `namespace` varchar(255) NOT NULL default '' COMMENT 'cluster namespace' AFTER `cluster_id`;
ALTER TABLE deploy_service_ip_list ADD COLUMN `namespace` varchar(255) NOT NULL default '' COMMENT 'cluster namespace' AFTER `cluster_id`;
ALTER TABLE deploy_service_ip_node ADD COLUMN `namespace` varchar(255) NOT NULL default '' COMMENT 'cluster namespace' AFTER `cluster_id`;
ALTER TABLE deploy_cluster_product_rel MODIFY COLUMN `pid` int(11);

alter table deploy_instance_list drop index `cluster_pid_service_name`;
alter table deploy_service_ip_list drop index `cluster_product_service_name`;
alter table deploy_service_ip_list add unique key `cluster_ns_product_service_name` (`cluster_id`,`namespace`,`product_name`,`service_name`);
alter table deploy_service_ip_node drop index `cluster_product_service_ip`;
alter table deploy_service_ip_node add unique key `cluster_ns_product_service_ip` (`cluster_id`,`namespace`,`product_name`,`service_name`,`ip`);
