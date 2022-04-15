
CREATE TABLE IF NOT EXISTS `deploy_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sid` varchar(255) NULL DEFAULT '' COMMENT '主机sid',
  `hostname` varchar(255) NULL DEFAULT '' COMMENT '主机名',
  `ip` varchar(255) NULL DEFAULT '' COMMENT '主机ip',
  `status` int(11) NULL DEFAULT '6' COMMENT '6:K8S NODE初始化成功,-6:K8S NODE初始化失败',
  `steps`  int(11) NULL DEFAULT '6' COMMENT '6:K8S NODE初始化成功',
  `group` varchar(255) DEFAULT 'default' COMMENT '组信息，兼容deploy_host',
  `errorMsg` varchar(1024) NULL DEFAULT '' COMMENT '错误信息',
  `isDeleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `updated` datetime NOT NULL COMMENT 'updated',
  `created` datetime NOT NULL COMMENT 'created',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_cluster_image_store` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `clusterId` int(11) UNSIGNED NOT NULL COMMENT '集群 id',
  `name` varchar(64) NOT NULL COMMENT '仓库名称',
  `alias` varchar(64) NOT NULL COMMENT '仓库别名',
  `address` varchar(256) NOT NULL COMMENT '仓库地址',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k8s镜像仓库表';

CREATE TABLE IF NOT EXISTS `deploy_cluster_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) NOT NULL COMMENT '集群名',
  `type` varchar(255) NOT NULL COMMENT '集群类型 hosts/kubernetes',
  `mode` int(11) NOT NULL DEFAULT '0' COMMENT '0:自建,1:导入',
  `version` varchar(255) NULL DEFAULT '' COMMENT '集群版本，主机集群为空',
  `desc` varchar(255) NULL DEFAULT '' COMMENT '集群描述',
  `tags` varchar(1024) NULL DEFAULT '' COMMENT '集群标签',
  `configs` text DEFAULT NULL COMMENT '集群个性化配置',
  `yaml` text DEFAULT NULL COMMENT '集群配置详情',
  `status` int(11) NULL DEFAULT '0' COMMENT '0:Waiting,1:Pending,2:Running,-2:Error',
  `errorMsg` varchar(1024) NULL DEFAULT '' COMMENT '错误信息',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `create_user` varchar(255) NULL DEFAULT 'admin' COMMENT '创建人',
  `update_user` varchar(255) NULL DEFAULT 'admin' COMMENT '修改人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_cluster_host_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sid` varchar(255) NULL DEFAULT '' COMMENT '主机sid',
  `clusterId` int(11) UNSIGNED NOT NULL COMMENT '集群 id',
  `roles` varchar(1024) NULL DEFAULT '' COMMENT '角色，逗号间隔',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `deploy_cluster_k8s_available` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `mode` int(11) NOT NULL DEFAULT '0' COMMENT '0:自建,1:导入',
  `version` varchar(255) NULL DEFAULT '' COMMENT 'k8s集群版本',
  `properties` text NOT NULL DEFAULT '' COMMENT '版本所允许的配置信息，分号区分配置项，冒号区分配置名，逗号区分配置选项',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_cluster_k8s_only` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `clusterId` int(11) UNSIGNED NOT NULL COMMENT '集群 id',
  `kube_config` text DEFAULT NULL COMMENT '集群配置文件，主机集群为空',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `deploy_cluster_list` (`id`, `name`, `type`, `mode`, `desc`, `tags`, `status`, `create_user`, `update_user`) VALUES ('1', 'dtstack', 'hosts', '0', '兼容EM2.0默认集群', '自动创建', '0', 'admin@dtstack.com', 'admin@dtstack.com');
INSERT INTO `deploy_cluster_host_rel` (`id`, `sid`, `clusterId`) select id, sid, '1' from deploy_host;
INSERT INTO `deploy_cluster_k8s_available` (`version`, `properties`) VALUES ('v1.16.3-rancher1-1', 'network_plugin:flannel');
ALTER TABLE sidecar_list ADD COLUMN `disk_usage_pct` decimal(6,2) DEFAULT '-1.00' COMMENT '磁盘使用率' AFTER `mem_usage`;



