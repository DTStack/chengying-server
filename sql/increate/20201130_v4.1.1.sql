
CREATE TABLE IF NOT EXISTS `deploy_namespace_client` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `yaml` text,
  `namespace_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_namespace_event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(255) DEFAULT NULL,
  `reason` varchar(255) DEFAULT NULL,
  `resource` varchar(255) DEFAULT NULL,
  `message` varchar(1500) DEFAULT NULL,
  `namespace_id` int(11) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `deploy_namespace_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(255) NOT NULL,
  `namespace` varchar(255) NOT NULL,
  `registry_id` int(11) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `port` varchar(255) DEFAULT NULL,
  `cluster_id` int(11) NOT NULL,
  `status` varchar(255) DEFAULT NULL,
  `is_deleted` int(11) DEFAULT NULL,
  `user` varchar(255) DEFAULT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `import_init_moudle` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `service_account` varchar(1000) DEFAULT NULL,
  `role` varchar(1000) DEFAULT NULL,
  `role_binding` varchar(1000) DEFAULT NULL,
  `operator` varchar(1000) DEFAULT NULL,
  `is_deleted` int(2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

INSERT INTO `import_init_moudle` VALUES (1, '{\"apiVersion\":\"v1\",\"kind\":\"ServiceAccount\",\"metadata\":{\"name\":\"dtstack\",\"namespace\":\"{{.NAME_SPACE}}\"}}', '{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"Role\",\"metadata\":{\"name\":\"dtstack-admin\",\"namespace\":\"{{.NAME_SPACE}}\"},\"rules\":[{\"apiGroups\":[\"*\"],\"resources\":[\"*\"],\"verbs\":[\"*\"]}]}', '{\"apiVersion\":\"rbac.authorization.k8s.io/v1\",\"kind\":\"RoleBinding\",\"metadata\":{\"name\":\"dtstack-admin-binding\",\"namespace\":\"{{.NAME_SPACE}}\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"Role\",\"name\":\"dtstack-admin\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"dtstack\",\"namespace\":\"{{.NAME_SPACE}}\"}]}', '{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"name\":\"mole-operator\",\"namespace\":\"{{.NAME_SPACE}}\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"name\":\"mole-operator\"}},\"template\":{\"metadata\":{\"labels\":{\"name\":\"mole-operator\"}},\"spec\":{\"containers\":[{\"command\":[\"mole-operator\"],\"env\":[{\"name\":\"WATCH_NAMESPACE\",\"value\":\"{{.NAME_SPACE}}\"},{\"name\":\"POD_NAME\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.name\"}}},{\"name\":\"OPERATOR_NAME\",\"value\":\"mole-operator\"}],\"image\":\"{{.REGISTRY}}/mole:v1.0.17\",\"imagePullPolicy\":\"Always\",\"name\":\"mole-operator\",\"resources\":{\"limits\":{\"cpu\":\"500m\",\"memory\":\"500Mi\"}}}],\"imagePullSecrets\":[{\"name\":\"{{.SECRET_NAME}}\"}]}}}}', 0);
