BEGIN;
INSERT INTO `workload_definition` VALUES (7, 'plugin', 'v1', '[{\"key\":\"Image\",\"ref\":\"spec.workloadpatrs.0.steps.0.object.image\"}]', 1);
INSERT INTO `workload_part` VALUES (7, 'image-push', 'job', '{}', 7);
INSERT INTO `workload_step` VALUES (32, 'image', 'container', 'bound', '{\"image\":\"\",\"command\":[\"/bin/sleep\"],\"args\":[\"5\"],\"resources\":{\"limits\":{\"cpu\":\"50m\",\"memory\":\"10Mi\"},\"requests\":{\"cpu\":\"0m\",\"memory\":\"0Mi\"}}}', 7);
COMMIT;
