
alter table deploy_product_list add COLUMN user_id int(11) NOT NULL default 0 COMMENT '部署人id';
alter table deploy_product_history add COLUMN user_id int(11) not null default 0 COMMENT '部署人id';;

update deploy_product_list set user_id = 1 where user_id = 0;
update deploy_product_history set user_id = 1 where user_id = 0;

insert into role_list(id,role_name,role_value,role_desc) values(-1,"system",-1,"system");
