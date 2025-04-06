
CREATE TABLE if NOT EXISTS `go_shell`.`zg_ag`  (
  `code` int(255) NULL COMMENT '股票代码',
  `date` date NULL COMMENT '日期',
  `jlr` varchar(255) NULL COMMENT '净流入金额',
  `jlrzb` varchar(255) NULL COMMENT '净流入占比',
  `zljlr` varchar(255) NULL COMMENT '主力净额金额',
  `zljlrzb` varchar(255) NULL COMMENT '主力净额占比',
  `cddjlr` varchar(255) NULL COMMENT '超大单净流入金额',
  `cddjlrzb` varchar(255) NULL COMMENT '超大单净流入占比',
  `ddjlr` varchar(255) NULL COMMENT '大单净流入金额',
  `ddjlrzb` varchar(255) NULL COMMENT '大单净流入占比'
);

ALTER TABLE `go_shell`.`zg_ag` 
ADD UNIQUE INDEX `unique_IDX`(`code`, `date`, `jlr`, `jlrzb`);

ALTER TABLE `go_shell`.`zg_ag` 
ADD INDEX `code_IDX`(`code`) USING BTREE,
ADD INDEX `date_IDX`(`date`) USING BTREE;
