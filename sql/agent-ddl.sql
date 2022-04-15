CREATE DATABASE IF NOT EXISTS `dtagent` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE dtagent;

CREATE TABLE IF NOT EXISTS `agent_list` (
  `id` char(36) NOT NULL COMMENT 'Agent ID (UUID)',
  `sidecar_id` char(36) NOT NULL COMMENT 'sidecar id',
  `type` tinyint(1) NOT NULL COMMENT 'agent类型',
  `name` char(32) NULL DEFAULT '' COMMENT 'agent 名称',
  `version` char(32) NOT NULL DEFAULT '' COMMENT 'agent版本',
  `is_uninstalled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已被卸载',
  `deploy_date` datetime DEFAULT NULL COMMENT 'agent部署时间',
  `auto_deployment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动部署的',
  `last_update_date` datetime DEFAULT NULL COMMENT '最近更新时间',
  `auto_updated` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动升级的',
  PRIMARY KEY (`id`),
  KEY `uuid` (`sidecar_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Sidecar管控的Agent信息表';

CREATE TABLE IF NOT EXISTS `operation_history` (
  `seq` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '操作序列',
  `op_name` varchar(20) NOT NULL DEFAULT '' COMMENT '操作名称',
  `op_time` datetime NOT NULL COMMENT '操作时间',
  `target` char(36) NOT NULL DEFAULT '' COMMENT '目标id（sidecar id）',
  `send_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '操作状态',
  `op_result` tinyint(1) DEFAULT '1' COMMENT '操作执行状态',
  `op_return_msg` mediumblob DEFAULT NULL COMMENT '操作返回内容',
  `finish_time` datetime DEFAULT NULL COMMENT '操作执行结束时间',
  PRIMARY KEY (`seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `progress_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `seq` int(11) unsigned NOT NULL COMMENT '对应操作序列号',
  `ts` datetime NOT NULL COMMENT '事件时间',
  `progress` decimal(5,2) NOT NULL DEFAULT '0.00' COMMENT '进度百分比',
  `sidecar_id` char(36) NOT NULL DEFAULT '' COMMENT 'sidecar id',
  `agent_id` char(36) DEFAULT '' COMMENT 'agent id',
  `msg` varchar(100) DEFAULT '' COMMENT '附带信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `sidecar_list` (
  `id` char(36) NOT NULL COMMENT 'Sidecar ID (UUID)',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'Sidecar状态',
  `disabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被屏蔽',
  `name` varchar(32) DEFAULT '' COMMENT 'Sidecar备注名',
  `version` varchar(32) DEFAULT '' COMMENT 'Sidecar版本',
  `host` varchar(255) DEFAULT '' COMMENT '主机域名或者ip',
  `os_type` varchar(16) DEFAULT '' COMMENT '目标系统类型,linux，windows等',
  `is_ecs` tinyint(1) DEFAULT '0' COMMENT '是否是ECS',
  `os_platform` varchar(64) DEFAULT '' COMMENT 'os完整的名称',
  `os_version` varchar(64) DEFAULT '' COMMENT 'os版本号',
  `cpu_serial` varchar(64) DEFAULT '' COMMENT 'cpu型号',
  `cpu_cores` tinyint(8) DEFAULT '0' COMMENT 'cpu内核数',
  `mem_size` bigint(20) unsigned DEFAULT '0' COMMENT '内存容量',
  `swap_size` bigint(20) unsigned DEFAULT '0' COMMENT '交换空间容量',
  `deploy_date` datetime DEFAULT NULL COMMENT 'Sidecar部署时间',
  `auto_deployment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动部署的',
  `last_update_date` datetime DEFAULT NULL COMMENT '最近更新时间',
  `auto_updated` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动升级的',
  `server_host` varchar(255) DEFAULT '' COMMENT 'api server ip',
  `server_port` int(11) DEFAULT 0 COMMENT 'api server port',
  `ssh_host` varchar(255) DEFAULT '' COMMENT '安装用的ssh主机域名或者ip',
  `ssh_user` varchar(60) DEFAULT '' COMMENT 'ssh用户名',
  `ssh_password` varchar(100) DEFAULT '' COMMENT 'ssh密码',
  `ssh_port` int(11) DEFAULT '22' COMMENT 'ssh端口',
  `cpu_usage` decimal(6,2) DEFAULT '-1' COMMENT 'cpu使用率',
  `mem_usage` bigint(20) DEFAULT '-1' COMMENT '物理内存使用',
  `swap_usage` bigint(20) DEFAULT '-1' COMMENT '交换空间使用',
  `load1` float DEFAULT '-1' COMMENT 'cpu load1',
  `uptime` double DEFAULT '-1' COMMENT '系统启动时间',
  `disk_usage` text DEFAULT NULL COMMENT '各个硬盘使用率',
  `net_usage` text DEFAULT NULL COMMENT '各个网卡统计',
  `local_ip` varchar(255) DEFAULT '' COMMENT '主机ip',
  PRIMARY KEY (`id`),
  KEY `uuid` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Sidecar客户端';

CREATE TABLE IF NOT EXISTS `deploy_callback` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto',
  `time` int(11) unsigned NOT NULL DEFAULT '0',
  `client_id` varchar(64) DEFAULT '' COMMENT '客户端的UUID',
  `install_type` varchar(32) DEFAULT '' COMMENT '安装类型 [sidecar 等等]',
  `install_res` varchar(32) DEFAULT '' COMMENT '安装标识信息[success,failed 等等]',
  `msg` varchar(255) DEFAULT '' COMMENT '安装结论信息',
  `request_url` varchar(2000) DEFAULT '' COMMENT '访问的回调原始请求',
  `ip` varchar(32) DEFAULT '' COMMENT 'ip地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='一键安装部署回调数据表';

CREATE TABLE IF NOT EXISTS `deploy_host` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sid` varchar(255) NULL DEFAULT '' COMMENT '主机sid',
  `hostname` varchar(255) NULL DEFAULT '' COMMENT '主机名',
  `ip` varchar(255) NULL DEFAULT '' COMMENT '主机ip',
  `status` int(11) NULL DEFAULT '0' COMMENT '1:管控安装成功,-1:管控安装失败,2:script安装成功,-2:script安装失败,3:主机初始化成功,-3:主机初始化失败',
  `steps`  int(11) NULL DEFAULT '0' COMMENT '0:默认;1:管控安装成功;2:script wrapper安装成功;3:主机初始化成功',
  `errorMsg` varchar(1024) NULL DEFAULT '' COMMENT '错误信息',
  `group` varchar(255) DEFAULT 'default' COMMENT '组信息',
  `isDeleted` int(11) NOT NULL DEFAULT '0' COMMENT '0:未删除,1:已删除',
  `updated` datetime NOT NULL COMMENT 'updated',
  `created` datetime NOT NULL COMMENT 'created',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_product_list` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `deploy_uuid` varchar(36) NOT NULL COMMENT '部署uuid',
  `parent_product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_name` varchar(255) NOT NULL COMMENT '组件名称',
  `product_name_display` varchar(255) NOT NULL default '' COMMENT '组件显示名称',
  `product_version` varchar(255) NOT NULL COMMENT '产品版本',
  `product` text NOT NULL COMMENT '产品信息',
  `product_parsed` text NOT NULL COMMENT '已经解析的产品信息',
  `is_current_version` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否当前版本',
  `status` enum('undeployed', 'deploying', 'deployed', 'deploy fail', 'undeploying', 'undeploy fail') NOT NULL DEFAULT 'undeployed' COMMENT '产品状态',
  `alert_recover` tinyint(1) NOT NULL default '0' COMMENT '0，不恢复告警，1，恢复告警',
  `user_id` int(11) NOT NULL default 0 COMMENT '部署人id',
  `deploy_time` datetime COMMENT '部署时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
  `schema` text NOT NULL COMMENT '产品原始schema',
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_name_version` (`product_name`, `product_version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_product_history` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `deploy_uuid` char(36) NOT NULL COMMENT '部署uuid',
  `parent_product_name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '组件名称',
  `product_name_display` varchar(255) NOT NULL DEFAULT '' COMMENT '组件显示名称',
  `product_version` varchar(255) NOT NULL DEFAULT '' COMMENT '产品版本',
  `status` char(32) NOT NULL DEFAULT '' COMMENT '产品状态',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
  `deploy_start_time` datetime COMMENT 'deploy_start_time',
  `deploy_end_time` datetime COMMENT 'deploy_end_time',
  `user_id` int(11) NOT NULL default 0 COMMENT '部署人id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_schema_field_modify` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '组件名称',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
  `field_path` varchar(255) NOT NULL DEFAULT '' COMMENT '字段路径',
  `field` text NOT NULL COMMENT '字段值',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update_time',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `namesz_path` (`product_name`, `service_name`, `field_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_service_ip_list` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '组件名称',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
  `ip_list` varchar(1024) NOT NULL DEFAULT '' COMMENT 'IP列表',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update_time',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create_time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_service_name` (`product_name`, `service_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_instance_list` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `agent_id` char(36) NOT NULL DEFAULT '' COMMENT 'agent id',
  `sid` char(36) NOT NULL COMMENT '主机 ID (UUID)',
  `pid` int(11) UNSIGNED NOT NULL COMMENT 'product id',
  `ip` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip',
  `group` varchar(255) NOT NULL DEFAULT '' COMMENT '组名称',
  `prometheus_port` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'prometheus port',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `service_name_display` varchar(255) NOT NULL default '' COMMENT '服务显示名称',
  `service_version` varchar(255) NOT NULL DEFAULT '' COMMENT '服务版本',
  `schema` text NOT NULL  COMMENT 'instance schema',
  `ha_role_cmd` varchar(255) NOT NULL COMMENT 'HA角色执行命令或脚本',
  `health_state` tinyint(2) NOT NULL DEFAULT -2 COMMENT '健康状态,0:不健康,1:健康,-1:未设置,-2:等待',
  `status` enum('installing', 'installed', 'install fail', 'uninstalling', 'uninstalled', 'uninstall fail', 'running', 'run fail', 'stopping', 'stopped', 'stop fail') NOT NULL DEFAULT 'installing' COMMENT '实例状态',
  `status_message` varchar(1024) NOT NULL DEFAULT '' COMMENT '状态详细信息',
  `heart_time` datetime DEFAULT NULL COMMENT '心跳更新时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `pid_service_name` (`pid`, `service_name`, `sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例列表';

CREATE TABLE IF NOT EXISTS `deploy_instance_record` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `deploy_uuid` char(36) NOT NULL DEFAULT '' COMMENT '部署记录uuid',
  `instance_id` int(11) UNSIGNED NOT NULL COMMENT '实例id',
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
  `progress` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '进度',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY (`deploy_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例部署记录表';

CREATE TABLE IF NOT EXISTS deploy_service_ip_node
(
  id int(11) unsigned auto_increment comment 'id' primary key,
  product_name varchar(255) default '' not null comment '组件名称',
  service_name varchar(255) default '' not null comment '服务名称',
  ip varchar(20) default '' not null comment 'IP列表',
  node_id int default '0' not null comment '序号',
  created_at timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
  constraint product_service_ip unique (product_name, service_name, ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='NodeId记录表';

CREATE TABLE IF NOT EXISTS `user_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(256) NOT NULL COMMENT '登录密码',
  `company` varchar(128) NOT NULL DEFAULT '' COMMENT '用户所属公司',
  `full_name` varchar(128) NOT NULL DEFAULT '' COMMENT '姓名',
  `email` varchar(255) NOT NULL COMMENT '邮箱地址',
  `phone` varchar(255) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0 启动，1 禁用',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `role_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_name` varchar(256) NOT NULL COMMENT '角色名称',
  `role_value` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'ADMIN(1), 运维(2)',
  `role_desc` varchar(256) NOT NULL DEFAULT '' COMMENT '角色描述',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常 1逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

CREATE TABLE IF NOT EXISTS `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0正常 1逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';

CREATE TABLE IF NOT EXISTS `deploy_unchecked_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) UNSIGNED NOT NULL COMMENT 'product id',
  `unchecked_services` varchar(255) NOT NULL COMMENT '未勾选服务',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='未勾选服务表';

insert into role_list(id,role_name,role_value,role_desc) values(-1,"system",-1,"system");
insert into role_list(id,role_name,role_value,role_desc) values(1,"管理员","1","管理员角色");
insert into role_list(id,role_name,role_value,role_desc) values(2,"运维人员","2","运维人员角色");
insert into user_list ( `phone`, `password`, `id`, `company`, `username`, `email`, `full_name`) values ( '11111111111', 'DT#passw0rd2019', '1', 'dtstack', 'admin@dtstack.com', 'admin@dtstack.com', 'admin');
insert into user_role(id,role_id,user_id) values(1,1,1);

CREATE TABLE IF NOT EXISTS `deploy_instance_event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `instance_id` int(11) UNSIGNED NOT NULL COMMENT 'instance id',
  `event_type` enum('install', 'uninstall', 'config update', 'start', 'stop', 'exec', 'error', 'unkown') NOT NULL DEFAULT 'unkown' COMMENT '事件类型',
  `content` text NOT NULL DEFAULT '' COMMENT '事件内容',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例事件列表';



CREATE TABLE IF NOT EXISTS `deploy_instance_runtime_event` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `event_type` varchar(256) NOT NULL DEFAULT '' COMMENT '事件类型',
    `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
    `parent_product_name` varchar(255) NOT NULL COMMENT '父产品名称',
    `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
    `host` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip',
    `content` text NOT NULL COMMENT '事件内容描述' ,
    `isDeleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='事件列表';

CREATE TABLE IF NOT EXISTS  `deploy_strategy_list` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '策略名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '策略简介',
  `property` int(1) NOT NULL DEFAULT '0' COMMENT '0:服务类型、1:主机类型',
  `strategy_type` int(1) NOT NULL DEFAULT '0' COMMENT '0:脚本、1:编码',
  `deploy_status` int(1) NOT NULL DEFAULT '0' COMMENT '0:未发布、1:发布',
  `exe_status` int(1) NOT NULL DEFAULT '0' COMMENT '0:正常、1:异常',
  `error_message` text NOT NULL DEFAULT '' COMMENT '调度状态',
  `start_date` datetime NOT NULL COMMENT '生效日期',
  `end_date` datetime NOT NULL COMMENT '结束日期',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `cron_period` int(11) NOT NULL DEFAULT '0' COMMENT '调度周期, 0:分钟、1:小时、2:天',
  `cron_interval` int(11) NOT NULL DEFAULT '1' COMMENT '调度间隔时间',
  `cron_time` datetime DEFAULT NULL COMMENT '具体调度时间',
  `params` text NOT NULL DEFAULT '' COMMENT '参数，逗号间隔',
  `time_out` bigint(10) NOT NULL DEFAULT '-1' COMMENT '超时设置，单位s',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0否、1是',
  `gmt_create` datetime NOT NULL COMMENT '创建日期',
  `gmt_modified` datetime NOT NULL COMMENT '最近更新日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='策略表';


CREATE TABLE IF NOT EXISTS  `deploy_strategy_resource_list` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `strategy_id` bigint(11) unsigned NOT NULL COMMENT '策略id',
  `content` text NOT NULL DEFAULT '' COMMENT '资源内容',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0否、1是',
  `gmt_create` datetime NOT NULL COMMENT '创建日期',
  `gmt_modified` datetime NOT NULL COMMENT '最近更新日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='策略资源表';


CREATE TABLE IF NOT EXISTS  `deploy_strategy_assign_list` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `strategy_id` bigint(11) unsigned NOT NULL COMMENT '策略id',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
  `parent_product_name` varchar(255) NOT NULL COMMENT '父亲产品名称',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称,逗号间隔',
  `host` varchar(255) NOT NULL DEFAULT '' COMMENT '主机ip,逗号间隔',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0否、1是',
  `gmt_create` datetime NOT NULL COMMENT '创建日期',
  `gmt_modified` datetime NOT NULL COMMENT '最近更新日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='策略分配表';


INSERT INTO deploy_instance_runtime_event(event_type,content) VALUES('JavaHeapDump','{}');
INSERT INTO deploy_instance_runtime_event(event_type,content) VALUES('服务被动拉起','{}');
INSERT INTO deploy_strategy_list VALUES (1,'JavaHeapDump','JavaHeapDump',1,0,1,0,'','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,5,NULL,'',10,0,'0000-00-00 00:00:00','0005-00-00 00:00:00');
INSERT INTO deploy_strategy_resource_list VALUES (1,1,'#!/bin/sh\n\ncurrent=`date \"+%Y-%m-%d %H:%M:%S\"`\ntimeStamp=`date -d \"$current\" +%s`\n\ntmp=\"/tmp/heapdumps_{SERVICENAME}_$timeStamp\"\n\nfind -L /opt/dtstack/{PRODUCTNAME}/{SERVICENAME}/*/heapdump.hprof -maxdepth 5 -size -4096M -type f -mmin -5 -print > $tmp 2>/dev/null\n\nif [ -f \"$tmp\" ];then\nfor i in `cat $tmp`\ndo\necho \'{\"file_name\":{\"desc\":\"JavaHeapDump文件名称\",\"value\":\"\'$i\'\"},\"product_name\":{\"desc\":\"所属组件\",\"value\":\"{PRODUCTNAME}\"},\"service_name\":{\"desc\":\"启动服务\",\"value\":\"{SERVICENAME}\"},\"host\":{\"desc\":\"主机IP\",\"value\":\"{HOSTIP}\"},\"generate_time\":{\"desc\":\"生成时间\",\"value\":\"\'$current\'\"},\"action\":{\"desc\":\"操作\",\"value\":\"下载\",\"instance\":\"{INSTANCEID}\",\"path\":\"\'$i\'\"}}\'\nbreak\ndone\nfi\n\nrm -f $tmp\n',0,'0000-00-00 00:00:00','0000-00-00 00:00:00');

INSERT INTO `deploy_strategy_list` VALUES (2,'服务被动拉起','服务被动拉起',1,0,1,0,'','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',1,1,NULL,'',10,0,'0000-00-00 00:00:00','0000-00-00 00:00:00');
INSERT INTO `deploy_strategy_resource_list` VALUES (2,2,'#!/bin/sh\n\ncurrent=`date \"+%Y-%m-%d %H:%M:%S\"`\ntimeStamp=`date -d \"$current\" +%s`\n\nagentId={AGENTID}\ninterval={INTERVAL}\n\n#agentId=cd3d4341-e9ad-446d-9932-8884a034d1cb\n#interval=28\n\nlogs=/opt/dtstack/easymanager/easyagent/logs/agent.log\n\ncdate=`date -d \"-$interval hour\" +\"%Y/%m/%d %H\"`\n\nuser=`whoami`\ns=0\n\nfor i in $(cat $logs |grep \"$cdate\"|grep \"exit(exit status 1\"|grep \"$agentId\"| awk \'{print$1\"#\"$2}\')\ndo\nret=`echo $(echo $(echo \"$i\"|sed \"s/\\//-/g\")| sed \"s/AGENT-DEBUG://g\" )|sed \"s/#/ /g\"`\nresults[s]=\'{\"start_time\":{\"desc\":\"启动时间\",\"value\":\"\'$ret\'\"},\"service_name\":{\"desc\":\"启动服务\",\"value\":\"{SERVICENAME}\"},\"host\":{\"desc\":\"启动主机\",\"value\":\"{HOSTIP}\"},\"product_name\":{\"desc\":\"所属组件\",\"value\":\"{PRODUCTNAME}\"},\"run_user\":{\"desc\":\"启动用户\",\"value\":\"\'$user\'\"}}\'\ns=$[$s+1];\ndone\n\nlen=${#results[@]}\nfor ((i=$len - 1;i>=0;i--))\ndo\n    echo ${results[$i]}\ndone',0,'0000-00-00 00:00:00','0000-00-00 00:00:00');

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


INSERT INTO `deploy_cluster_list` (`id`, `name`, `type`, `mode`, `desc`, `tags`, `status`, `create_user`, `update_user`, `update_time`, `create_time`) VALUES ('1', 'dtstack', 'hosts', '0', '兼容EM2.0默认集群', '自动创建', '0', 'admin@dtstack.com', 'admin@dtstack.com', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `deploy_cluster_host_rel` (`id`, `sid`, `clusterId`,  `update_time`, `create_time`) select id, sid, '1','0000-00-00 00:00:00','0000-00-00 00:00:00' from deploy_host;
INSERT INTO `deploy_cluster_k8s_available` (`version`, `properties`) VALUES ('v1.16.3-rancher1-1', 'network_plugin:flannel');
ALTER TABLE sidecar_list ADD COLUMN `disk_usage_pct` decimal(6,2) DEFAULT '-1.00' COMMENT '磁盘使用率' AFTER `mem_usage`;

CREATE TABLE IF NOT EXISTS `safety_audit_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `operator` varchar(255) NOT NULL COMMENT '操作人',
  `module` varchar(255) NOT NULL COMMENT '操作模块',
  `operation` varchar(255) NOT NULL COMMENT '动作',
  `ip` varchar(255) NOT NULL COMMENT '来源ip',
  `content` text NOT NULL COMMENT '详细内容',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `audit_item_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `module` varchar(255) NOT NULL COMMENT '操作模块',
  `operation` varchar(255) NOT NULL COMMENT '动作',
  `is_deleted` int(11) NOT NULL DEFAULT '0' COMMENT '1: 已删除',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

UPDATE role_list SET role_list.role_name='Administrator', role_list.role_desc='超级管理员，具备产品所有操作权限' WHERE role_list.role_value='1';
UPDATE role_list SET role_list.role_name='Cluster Operator', role_list.role_desc='集群操作人员，一般指运维人员，具有安装部署、集群运维、监控告警功能操作权限' WHERE role_list.role_value='2';
INSERT INTO role_list(role_name, role_value, role_desc) VALUES('Cluster Reader', 4, '普通用户，只有集群的只读权限');

INSERT INTO audit_item_list(module, operation) VALUES('产品访问', '进入EM');
INSERT INTO audit_item_list(module, operation) VALUES('产品访问', '退出EM');
INSERT INTO audit_item_list(module, operation) VALUES('用户管理', '创建账号');
INSERT INTO audit_item_list(module, operation) VALUES('用户管理', '禁用账号');
INSERT INTO audit_item_list(module, operation) VALUES('用户管理', '启用账号');
INSERT INTO audit_item_list(module, operation) VALUES('用户管理', '移除账号');
INSERT INTO audit_item_list(module, operation) VALUES('用户管理', '重置密码');
INSERT INTO audit_item_list(module, operation) VALUES('集群管理', '创建集群');
INSERT INTO audit_item_list(module, operation) VALUES('集群管理', '编辑集群');
INSERT INTO audit_item_list(module, operation) VALUES('集群管理', '删除集群');
INSERT INTO audit_item_list(module, operation) VALUES('集群管理', '添加主机');
INSERT INTO audit_item_list(module, operation) VALUES('集群管理', '删除主机');
INSERT INTO audit_item_list(module, operation) VALUES('部署向导', '产品部署');
INSERT INTO audit_item_list(module, operation) VALUES('部署向导', '产品卸载');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '组件停止');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '组件启动');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '服务停止');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '服务启动');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '服务滚动重启');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '服务参数修改');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '服务参数新增');
INSERT INTO audit_item_list(module, operation) VALUES('集群运维', '配置下发');
