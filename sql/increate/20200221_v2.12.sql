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
INSERT INTO `deploy_strategy_list` VALUES (1,'JavaHeapDump','JavaHeapDump',1,0,1,0,'','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',0,5,NULL,'',10,0,'0000-00-00 00:00:00','0005-00-00 00:00:00');
INSERT INTO `deploy_strategy_resource_list` VALUES (1,1,'#!/bin/sh\n\ncurrent=`date \"+%Y-%m-%d %H:%M:%S\"`\ntimeStamp=`date -d \"$current\" +%s`\n\ntmp=\"/tmp/heapdumps_{SERVICENAME}_$timeStamp\"\n\nfind -L /opt/dtstack/{PRODUCTNAME}/{SERVICENAME}/*/heapdump.hprof -maxdepth 5 -size -4096M -type f -mmin -5 -print > $tmp 2>/dev/null\n\nif [ -f \"$tmp\" ];then\nfor i in `cat $tmp`\ndo\necho \'{\"file_name\":{\"desc\":\"JavaHeapDump文件名称\",\"value\":\"\'$i\'\"},\"product_name\":{\"desc\":\"所属组件\",\"value\":\"{PRODUCTNAME}\"},\"service_name\":{\"desc\":\"启动服务\",\"value\":\"{SERVICENAME}\"},\"host\":{\"desc\":\"主机IP\",\"value\":\"{HOSTIP}\"},\"generate_time\":{\"desc\":\"生成时间\",\"value\":\"\'$current\'\"},\"action\":{\"desc\":\"操作\",\"value\":\"下载\",\"instance\":\"{INSTANCEID}\",\"path\":\"\'$i\'\"}}\'\nbreak\ndone\nfi\n\nrm -f $tmp\n',0,'0000-00-00 00:00:00','0000-00-00 00:00:00');

INSERT INTO `deploy_strategy_list` VALUES (2,'服务被动拉起','服务被动拉起',1,0,1,0,'','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00',1,1,NULL,'',10,0,'0000-00-00 00:00:00','0000-00-00 00:00:00');
INSERT INTO `deploy_strategy_resource_list` VALUES (2,2,'#!/bin/sh\n\ncurrent=`date \"+%Y-%m-%d %H:%M:%S\"`\ntimeStamp=`date -d \"$current\" +%s`\n\nagentId={AGENTID}\ninterval={INTERVAL}\n\n#agentId=cd3d4341-e9ad-446d-9932-8884a034d1cb\n#interval=28\n\nlogs=/opt/dtstack/easymanager/easyagent/logs/agent.log\n\ncdate=`date -d \"-$interval hour\" +\"%Y/%m/%d %H\"`\n\nuser=`whoami`\ns=0\n\nfor i in $(cat $logs |grep \"$cdate\"|grep \"exit(exit status 1\"|grep \"$agentId\"| awk \'{print$1\"#\"$2}\')\ndo\nret=`echo $(echo $(echo \"$i\"|sed \"s/\\//-/g\")| sed \"s/AGENT-DEBUG://g\" )|sed \"s/#/ /g\"`\nresults[s]=\'{\"start_time\":{\"desc\":\"启动时间\",\"value\":\"\'$ret\'\"},\"service_name\":{\"desc\":\"启动服务\",\"value\":\"{SERVICENAME}\"},\"host\":{\"desc\":\"启动主机\",\"value\":\"{HOSTIP}\"},\"product_name\":{\"desc\":\"所属组件\",\"value\":\"{PRODUCTNAME}\"},\"run_user\":{\"desc\":\"启动用户\",\"value\":\"\'$user\'\"}}\'\ns=$[$s+1];\ndone\n\nlen=${#results[@]}\nfor ((i=$len - 1;i>=0;i--))\ndo\n    echo ${results[$i]}\ndone',0,'0000-00-00 00:00:00','0000-00-00 00:00:00');