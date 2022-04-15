
CREATE TABLE IF NOT EXISTS `addons_list`(
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) NOT NULL COMMENT 'addon类型',
  `desc` varchar(255) NOT NULL COMMENT 'addon描述',
  `version` varchar(255) NOT NULL COMMENT 'addon版本',
  `os` varchar(255) NOT NULL COMMENT '支持OS类型, linux|windows|docker',
  `schema` text NOT NULL COMMENT 'addon 运行相关schema',
  `isDeleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除',
  `updated` datetime NOT NULL COMMENT '更新时间',
  `created` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `addons_name_version` (`type`, `version`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS  `deploy_addons_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `aid` int(11) NOT NULL COMMENT 'addon id',
  `sid` varchar(255) NOT NULL DEFAULT '' COMMENT 'sid',
  `agentId` varchar(255) NOT NULL DEFAULT '' COMMENT 'agentId',
  `config` text NOT NULL COMMENT '配置',
  `addon_type` varchar(255) NOT NULL DEFAULT 'standalone' COMMENT 'addon类型,standalone|rpc|api',
  `addon_version` varchar(255) NOT NULL DEFAULT '' COMMENT 'addon版本',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `status_message` varchar(255) NOT NULL DEFAULT '' COMMENT 'error信息',
  `isDeleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除',
  `updated` datetime NOT NULL COMMENT '更新时间',
  `created` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;