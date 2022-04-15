ALTER TABLE deploy_product_list ADD COLUMN `schema` text NOT NULL DEFAULT '' COMMENT '产品原始schema';
update deploy_product_list set `schema`=`product` where `schema`='';
