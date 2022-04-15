alter table deploy_product_list add COLUMN parent_product_name varchar(255) NOT NULL COMMENT '父类产品名称';
alter table deploy_product_history add COLUMN parent_product_name varchar(255) NOT NULL COMMENT '父类产品名称';
alter table deploy_instance_list add COLUMN heart_time datetime default NULL COMMENT '心跳更新时间';
alter table deploy_instance_list add COLUMN ha_role_cmd varchar(255) NOT NULL COMMENT 'HA角色执行命令或脚本';

update deploy_product_list a,(select id , product_name from deploy_product_list) b set a.parent_product_name = b.product_name where a.id=b.id;
update deploy_product_history a,(select id , product_name from deploy_product_history) b set a.parent_product_name = b.product_name where a.id=b.id;
