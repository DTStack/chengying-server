
alter table deploy_product_list add COLUMN deploy_uuid varchar(36) NOT NULL default '' COMMENT '部署uuid';
alter table deploy_product_list add COLUMN product_name_display varchar(255) not null default '' COMMENT '组件显示名';

alter table deploy_product_history add COLUMN deploy_uuid char(36) NOT NULL default '' COMMENT '部署uuid';
alter table deploy_product_history add COLUMN deploy_start_time datetime COMMENT '部署开始时间';
alter table deploy_product_history add COLUMN deploy_end_time datetime COMMENT '部署结束时间';
alter table deploy_product_history add COLUMN product_name_display varchar(255) not null default '' COMMENT '组件显示名';

alter table deploy_instance_list add COLUMN service_name_display varchar(255) not null default '' COMMENT '服务显示名';

alter table deploy_instance_record add COLUMN product_name_display varchar(255) not null default '' COMMENT '组件显示名';
alter table deploy_instance_record add COLUMN service_name_display varchar(255) not null default '' COMMENT '服务显示名';


update deploy_product_list a,(select id , product_name from deploy_product_list) b set a.product_name_display = b.product_name where a.id=b.id;
update deploy_product_history a,(select id , product_name from deploy_product_history) b set a.product_name_display = b.product_name where a.id=b.id;
update deploy_instance_list a,(select id , service_name from deploy_instance_list) b set a.service_name_display = b.service_name where a.id=b.id;
update deploy_instance_record a,(select id , product_name, service_name from deploy_instance_record) b set a.product_name_display = b.product_name,a.service_name_display = b.service_name where a.id=b.id;




-- 历史记录表，在升级2.6版本的时候与原来的是不兼容的
delete from deploy_product_history where deploy_uuid = ''

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
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0正常 1逻辑删除',
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

insert into role_list(id,role_name,role_value,role_desc) values(1,"管理员","1","管理员角色");
insert into role_list(id,role_name,role_value,role_desc) values(2,"运维人员","2","运维人员角色");

insert into user_list ( `phone`, `password`, `id`, `company`, `username`, `email`, `full_name`) values ( '11111111111', 'admin123', '1', 'dtstack', 'admin@dtstack.com', 'admin@dtstack.com', 'admin');

insert into user_role(id,role_id,user_id) values(1,1,1);

