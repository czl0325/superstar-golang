/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : superstar

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 20/08/2019 15:32:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for star_info
-- ----------------------------
DROP TABLE IF EXISTS `star_info`;
CREATE TABLE `star_info` (
  `id` int(255) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(255) DEFAULT NULL COMMENT '中文名',
  `name_en` varchar(255) DEFAULT NULL COMMENT '英文名',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `birthday` varchar(255) DEFAULT NULL COMMENT '出生日期',
  `height` double DEFAULT NULL COMMENT '身高，单位cm',
  `weight` double DEFAULT NULL COMMENT '体重，单位kg\n',
  `club` varchar(255) DEFAULT NULL COMMENT '俱乐部',
  `jersy` varchar(255) DEFAULT NULL COMMENT '球衣号码以及主打位置',
  `country` varchar(255) DEFAULT NULL COMMENT '国籍',
  `birthaddress` varchar(255) DEFAULT NULL COMMENT '出生地',
  `feature` varchar(255) DEFAULT NULL COMMENT '个人特点',
  `moreinfo` varchar(255) DEFAULT NULL COMMENT '更多介绍',
  `sys_status` int(255) NOT NULL COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(11) NOT NULL COMMENT '创建时间',
  `sys_updated` int(11) NOT NULL COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of star_info
-- ----------------------------
BEGIN;
INSERT INTO `star_info` VALUES (1, 'C罗', 'Cristiano Ronaldo', 'http://www.sinaimg.cn/lf/sports/wc_2018/player/14937.jpg', '1985-02-05', 187, 83, '尤文图斯', '', '葡萄牙', '葡萄牙', '', '', 0, 1566284925, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
