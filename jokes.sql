/*
 Navicat Premium Data Transfer

 Source Server         : aliyun
 Source Server Type    : MySQL
 Source Server Version : 50715
 Source Host           : ubuntu-0706.mysql.rds.aliyuncs.com:3306
 Source Schema         : jokes

 Target Server Type    : MySQL
 Target Server Version : 50715
 File Encoding         : 65001

 Date: 05/09/2017 22:52:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crawls
-- ----------------------------
DROP TABLE IF EXISTS `crawls`;
CREATE TABLE `crawls` (
  `c_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `unique` varchar(50) DEFAULT NULL,
  `source` varchar(50) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `is_del` bigint(20) DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`c_id`),
  UNIQUE KEY `uix_crawls_unique` (`unique`)
) ENGINE=InnoDB AUTO_INCREMENT=159 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for img_jokes
-- ----------------------------
DROP TABLE IF EXISTS `img_jokes`;
CREATE TABLE `img_jokes` (
  `i_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `c_id` bigint(20) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `img_list` varchar(255) DEFAULT NULL,
  `source_list` varchar(255) DEFAULT NULL,
  `img_count` bigint(20) DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`i_id`)
) ENGINE=InnoDB AUTO_INCREMENT=724 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for text_jokes
-- ----------------------------
DROP TABLE IF EXISTS `text_jokes`;
CREATE TABLE `text_jokes` (
  `t_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `c_id` bigint(20) DEFAULT NULL,
  `source` varchar(50) DEFAULT NULL,
  `content` text,
  `create_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`t_id`)
) ENGINE=InnoDB AUTO_INCREMENT=588 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
