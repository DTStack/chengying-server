ALTER TABLE `dtagent`.`deploy_kube_base_product_list`
    CHANGE COLUMN `base_clusterId` `rely_namespace` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '依赖namespace' AFTER `namespace`;


create table if not exists operation_list
(
    id               int auto_increment
    primary key,
    cluster_id       int                                 not null comment '集群 id',
    operation_id     varchar(255)                        null comment '操作 id',
    operation_type   int                                 not null comment '1. 产品包部署 2. 产品包升级 3. 产品包启动 4. 服务启动 5. 服务滚动重启 6. 主机初始化 7. Kerberos开启 8. Kerberos关闭',
    operation_status int                                 not null comment '1 进行中 2 正常 3 失败',
    object_type      int                                 not null comment '对象类型 1：产品包 2：服务 3：主机',
    object_value     varchar(255)                        not null comment '对象值，用于页面回显',
    create_time      timestamp default CURRENT_TIMESTAMP null,
    end_time         timestamp                           null comment '结束时间',
    duration         float                               null comment '持续时间 单位秒',
    update_time      timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint operation_list_operation_id_uindex
    unique (operation_id)
    )
    comment '操作清单表' charset = utf8;


create table if not exists exec_shell_list
(
    id           int auto_increment
    primary key,
    cluster_id   int                                 not null comment '集群 id',
    exec_id      varchar(255)                        not null comment 'shell 执行 id ',
    operation_id varchar(255)                        not null comment '操作id',
    shell_type   int                                 not null comment '具体shell 类型 1 服务安装 2 服务启动 3 执行脚本',
    product_name varchar(255)                        null comment '所属产品包',
    service_name varchar(255)                        null comment '所属服务',
    sid          varchar(255)                        null comment '主机',
    seq          int                                 null comment 'exec seq',
    exec_status  int                                 null comment '1 进行中 2 正常 3 失败',
    create_time  timestamp default CURRENT_TIMESTAMP null,
    end_time     timestamp                           null comment '结束时间',
    duration     float                               null comment '持续时间 单位 秒',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint exec_shell_list_exec_id_uindex
    unique (exec_id)
    )
    comment 'shell 执行记录表' charset = utf8;

CREATE TABLE if not exists `deploy_inspect_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) NOT NULL COMMENT '巡检报告名称',
  `status` varchar(64) NOT NULL COMMENT '状态',
  `progress` tinyint(3) DEFAULT NULL COMMENT '进度',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `cluster_id` int(11) NOT NULL COMMENT '集群id',
  `file_path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='巡检报告表';
