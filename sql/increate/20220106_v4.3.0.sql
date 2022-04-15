CREATE TABLE `product_backup_config` (
 `id` INT(11) NOT NULL AUTO_INCREMENT,
 `cluster_id` VARCHAR(255) COLLATE UTF8MB4_BIN NOT NULL COMMENT '集群id',
 config_path VARCHAR(255) COLLATE UTF8MB4_BIN NOT NULL COMMENT '备份路径',
 create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
 PRIMARY KEY (id)
) ENGINE=INNODB DEFAULT CHARSET=UTF8MB4 COLLATE = UTF8MB4_BIN COMMENT='组件备份路径记录表';