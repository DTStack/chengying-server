alter table deploy_host
    add role_list varchar(255) null comment '角色列表';

create table host_role
(
    id         int auto_increment
        primary key,
    cluster_id int          null comment '集群 id',
    role_name  varchar(255) collate utf8_bin not null comment '角色名',
    role_type  int          not null comment '角色类型 1 默认 2 自定义',
    constraint host_role_list_cluster_id_role_name_uindex
        unique (cluster_id, role_name)
);
create table if not exists deploy_product_select_history
(
    cluster_id int          not null
        primary key,
    pid_list   varchar(255) null comment 'pid 清单'
);

create table if not exists deploy_uuid
(
    id          int auto_increment
        primary key,
    uuid        varchar(255)                       not null,
    type        int                                not null comment '部署类型： 1 手动；2 自动；3 自动部署中的子产品的 uuid',
    parent_uuid varchar(255)                       null comment '自动部署 uuid',
    pid         varchar(255)                       null comment '产品包 id',
    create_time datetime default CURRENT_TIMESTAMP not null
);


-- 执行完上面的建表语句后
--  查询当前所有未删除的集群id
-- select  id  from deploy_cluster_list a where is_deleted = 0 ;
--  用查询出的每个 id 替换以下插入语句中的每个 cluster_id_value 然后执行 insert 语句
insert into host_role (cluster_id, role_name, role_type) values (1,'web',1),(1,'manager',1),(1,'worker',1);

-- deploy_unchecked_service增加namespace字段和索引
ALTER TABLE `dtagent`.`deploy_unchecked_service`
ADD COLUMN `namespace` varchar(255) NOT NULL COMMENT 'k8s多命名空间' AFTER `create_time`;

ALTER TABLE `dtagent`.`deploy_unchecked_service`
DROP INDEX `cluster_pid`,
ADD UNIQUE INDEX `cluster_pid`(`cluster_id`, `pid`, `namespace`) USING BTREE;


-- deploy_schema_field_modify增加namespace字段和索引
ALTER TABLE `dtagent`.`deploy_schema_field_modify`
ADD COLUMN `namespace` varchar(255) NOT NULL COMMENT 'k8s多命名空间区分' AFTER `create_time`;

ALTER TABLE `dtagent`.`deploy_schema_field_modify`
DROP INDEX `cluster_names_path`,
ADD UNIQUE INDEX `cluster_names_path`(`cluster_id`, `product_name`, `service_name`, `field_path`, `namespace`) USING BTREE;
