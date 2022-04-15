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

-- 4.1.0开始数据库密码字段加密
-- 加密方式为md5 明文密码
-- DT#passw0rd2019 的 md5 为 ca6590a271539cc89e2cc20bd6b58518

UPDATE user_list SET password='ca6590a271539cc89e2cc20bd6b58518' where id = 1;
