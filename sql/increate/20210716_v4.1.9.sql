CREATE TABLE `deploy_upload_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `upload_type` varchar(16) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '上传类型',
  `name` varchar(2048) COLLATE utf8mb4_bin NOT NULL COMMENT '链接地址',
  `progress` decimal(10,0) DEFAULT '0' COMMENT '进度',
  `status` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状态',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_deleted` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

UPDATE `dtagent`.`import_init_moudle` SET `log_config` = '{\"apiVersion\":\"v1\",\"data\":{\"filebeat.yml\":\"logging.level: debug\\nfilebeat.inputs:\\n  - type: log\\n    tail_files: true\\n    index: \\\"log-${PRODUCT}-%{+yyyy-MM-dd}\\\"\\n    fields:\\n      namespace: ${NAMESPACE}\\n      serviceAccountName: ${SERVICE_ACCOUNT_NAME}\\n      product: ${PRODUCT}\\n      job: ${JOB}\\n      node: ${HOSTNAME}/${HOST_IP}\\n      pod_name: ${POD_NAME}\\n      pod_uid: ${POD_UID}\\n      pod_ip: ${POD_IP}\\n    tags: [ \\\"${PRODUCT}\\\",\\\"${JOB}\\\" ]\\n    paths: ${LOG_PATH}\\noutput.elasticsearch:\\n  hosts: [ \\\"${LOG_SERVER_ADDRESS}\\\" ]\\n  username: \\\"elastic\\\"\\n  password: \\\"dtstack\\\"\\n\",\"promtail.yaml\":\"client:\\n  backoff_config:\\n    max_period: 5m\\n    max_retries: 10\\n    min_period: 500ms\\n  batchsize: 1048576\\n  batchwait: 1s\\n  external_labels: {}\\n  timeout: 10s\\npositions:\\n  filename: /var/log/logs/positions.yaml\\nserver:\\n  http_listen_port: 3101\\ntarget_config:\\n  sync_period: 10s\\nscrape_configs:\\n  - job_name: test\\n    static_configs:\\n      - labels:\\n          namespace: ${NAMESPACE}\\n          serviceAccountName: ${SERVICE_ACCOUNT_NAME}\\n          product: ${PRODUCT}\\n          job: ${JOB}\\n          node: ${HOSTNAME}/${HOST_IP}\\n          pod_name: ${POD_NAME}\\n          pod_uid: ${POD_UID}\\n          pod_ip: ${POD_IP}\\n          __path__: ${LOG_PATH}\\n\"},\"kind\":\"ConfigMap\",\"metadata\":{\"name\":\"log-configmap\",\"namespace\":\"{{.NAME_SPACE}}\"}}' WHERE `id` = 1;
update workload_step set name="hs" where name="mysql-svc";
