
CREATE TABLE if NOT EXISTS `go_shell`.`zg_ag`  (
  `code` varchar(255) NULL COMMENT '股票代码',
  `date` varchar(255) NULL COMMENT '日期',
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


----

					
CREATE TABLE if NOT EXISTS `go_shell`.`zg_ag_sdgd`  (
  `code` varchar(255) NULL COMMENT '股票代码',
  `date` varchar(255) NULL COMMENT '报告日期',
  `gdmc` varchar(255) NULL COMMENT '股东名称',
  `gdbh` varchar(255) NULL COMMENT '股东编号',
  `gdcgsl` varchar(255) NULL COMMENT '股东持股数量',
  `gdlb` varchar(255) NULL COMMENT '股东类别',
  `gfxz` varchar(255) NULL COMMENT '股份性质',
  `wzdn1` varchar(255) NULL COMMENT '未知代码1',
  `wzdn2` varchar(255) NULL COMMENT '未知代码2',
  `wzdn3` varchar(255) NULL COMMENT '未知代码3',
  `wzdn4` varchar(255) NULL COMMENT '未知代码4'
);

ALTER TABLE `go_shell`.`zg_ag_sdgd` 
ADD UNIQUE INDEX `unique_IDX`(`code`, `date`, `gdmc`, `gdcgsl`);

ALTER TABLE `go_shell`.`zg_ag_sdgd` 
ADD INDEX `code_IDX`(`code`) USING BTREE,
ADD INDEX `date_IDX`(`date`) USING BTREE;
----