ALTER TABLE `smoke_testing` ADD `test_script` VARCHAR(255) NOT NULL COMMENT '测试脚本' AFTER `operation_id`;
ALTER TABLE `smoke_testing` ADD `error_message` text NOT NULL COMMENT '错误信息' AFTER `report_url`;
