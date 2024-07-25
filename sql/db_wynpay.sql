/*
 Navicat Premium Dump SQL

 Source Server         : 114.117.164.200
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : 114.117.164.200:3308
 Source Schema         : db_wynpay

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 25/07/2024 14:59:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for site_formation
-- ----------------------------
DROP TABLE IF EXISTS `site_formation`;
CREATE TABLE `site_formation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `api_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'apiKey and wegitID',
  `currency_code` varchar(255) DEFAULT NULL COMMENT 'currency code',
  `utm_source` varchar(255) DEFAULT NULL COMMENT 'source',
  `utm_medium` varchar(255) DEFAULT NULL COMMENT 'medium- video - article - other',
  `utm_campaign` varchar(255) DEFAULT NULL COMMENT 'campaign',
  `url` varchar(255) DEFAULT NULL COMMENT 'url',
  `gmt_create` datetime DEFAULT NULL,
  `gmt_modified` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='link database';

-- ----------------------------
-- Records of site_formation
-- ----------------------------
BEGIN;
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (60, 'ee3d61985c8d4807af526b2d6965b455', '11123', 'youtube', 'video', '', 'https://wynpay.io?appKey=ee3d61985c8d4807af526b2d6965b455&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=11123', '2024-07-22 11:29:50', '2024-07-22 11:29:50');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (61, '1f248e01e5d04ead984b4a39b186f86d', '182929', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=1f248e01e5d04ead984b4a39b186f86d&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=182929', '2024-07-22 11:31:32', '2024-07-22 11:31:32');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (62, '74730e834a464338a512ed385d0a147f', '1', '3', '2', '', 'https://www.wynpay.io/en?appKey=74730e834a464338a512ed385d0a147f&utm_source=3&utm_medium=2&utm_campaign=&currency_code=1', '2024-07-22 15:38:21', '2024-07-22 15:38:21');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (63, '6d96df4c78d348daa16e6fd27b12c8a4', '1112311', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=6d96df4c78d348daa16e6fd27b12c8a4&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=1112311', '2024-07-22 15:38:44', '2024-07-22 15:38:44');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (64, 'c8d3c8ae676445cbaaa213d772ba9061', '1', '1', '1', '', 'https://www.wynpay.io/en?appKey=c8d3c8ae676445cbaaa213d772ba9061&utm_source=1&utm_medium=1&utm_campaign=&currency_code=1', '2024-07-22 15:39:28', '2024-07-22 15:39:28');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (65, 'fce32288321f4580aa6ac951941439dd', '99002829', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=fce32288321f4580aa6ac951941439dd&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=99002829', '2024-07-22 15:42:12', '2024-07-22 15:42:12');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (66, '8852c88b92254daf80646d0e58f1b3d8', '2', '1', '1', '', 'https://www.wynpay.io/en?appKey=8852c88b92254daf80646d0e58f1b3d8&utm_source=1&utm_medium=1&utm_campaign=&currency_code=2', '2024-07-22 15:57:22', '2024-07-22 15:57:22');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (67, '3cfcdc4345944a9eb13d9371887238e6', '2', '1', '1', '', 'https://www.wynpay.io/en?appKey=3cfcdc4345944a9eb13d9371887238e6&utm_source=1&utm_medium=1&utm_campaign=&currency_code=2', '2024-07-22 16:36:01', '2024-07-22 16:36:01');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (68, '16afc4dc9414472980a45744204f8e74', '1', '1', '1', '', 'https://www.wynpay.io/en?appKey=16afc4dc9414472980a45744204f8e74&utm_source=1&utm_medium=1&utm_campaign=&currency_code=1', '2024-07-22 16:40:03', '2024-07-22 16:40:03');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (69, 'd915dac310094fbf9c705c381e999470', '111232222', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=d915dac310094fbf9c705c381e999470&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=111232222', '2024-07-22 16:40:20', '2024-07-22 16:40:20');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (70, '7f0e13591f954d9b9542d53287af4823', '222222', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=7f0e13591f954d9b9542d53287af4823&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=222222', '2024-07-22 16:40:30', '2024-07-22 16:40:30');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (71, 'db7a5fd8ed604cc1a62c90d38e440987', '11111111', 'youtube', 'video', '', 'https://www.wynpay.io/en?appKey=db7a5fd8ed604cc1a62c90d38e440987&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=11111111', '2024-07-22 16:49:04', '2024-07-22 16:49:04');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (72, '9ced0b2bf9344fa19acb64c12103dd6b', '2222222', 'youtube', 'video', '', 'https://www.wynpay.io/en/individual?appKey=9ced0b2bf9344fa19acb64c12103dd6b&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=2222222', '2024-07-22 16:50:12', '2024-07-22 16:50:12');
INSERT INTO `site_formation` (`id`, `api_key`, `currency_code`, `utm_source`, `utm_medium`, `utm_campaign`, `url`, `gmt_create`, `gmt_modified`) VALUES (73, 'dab005b4ed834aed9fb978dc7a2354a2', '123123', 'youtube', 'video', '', 'https://www.wynpay.io/en/individual?appKey=dab005b4ed834aed9fb978dc7a2354a2&utm_source=youtube&utm_medium=video&utm_campaign=&currency_code=123123', '2024-07-25 14:33:26', '2024-07-25 14:33:26');
COMMIT;

-- ----------------------------
-- Table structure for site_log
-- ----------------------------
DROP TABLE IF EXISTS `site_log`;
CREATE TABLE `site_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `unique_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'api_key',
  `device_type` varchar(255) DEFAULT NULL COMMENT 'device',
  `region` varchar(255) DEFAULT NULL COMMENT 'nation',
  `referer` varchar(255) DEFAULT NULL COMMENT 'form address',
  `utm_source` varchar(255) DEFAULT NULL COMMENT 'from source',
  `utm_medium` varchar(255) DEFAULT NULL COMMENT 'from medium',
  `utm_campaign` varchar(255) DEFAULT NULL COMMENT 'from campaign',
  `gmt_create` datetime DEFAULT NULL,
  `gmt_modified` datetime DEFAULT NULL,
  `everyday` date DEFAULT NULL COMMENT 'everyday',
  `count` int DEFAULT NULL COMMENT 'frequency',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='site_log';

-- ----------------------------
-- Records of site_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(500) DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(100) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(17) DEFAULT NULL COMMENT '电话号码',
  `avatar` varchar(200) DEFAULT NULL COMMENT '头像',
  `sex` tinyint(1) DEFAULT NULL COMMENT '性别',
  `memo` varchar(100) DEFAULT NULL COMMENT '备注',
  `last_login_ip` varchar(50) DEFAULT NULL COMMENT '最后一次登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后一次登录时间',
  `status` tinyint DEFAULT NULL COMMENT '状态：1为正常，0为禁止',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '删除标记（0:不可用 1:可用）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Records of user_info
-- ----------------------------
BEGIN;
INSERT INTO `user_info` (`id`, `username`, `password`, `nick_name`, `phone`, `avatar`, `sex`, `memo`, `last_login_ip`, `last_login_time`, `status`, `create_time`, `update_time`, `is_deleted`) VALUES (33, 'admin', '22850bded62a4795605b7358229c7855', 'admin', '11111000000', NULL, NULL, NULL, NULL, NULL, NULL, '2024-07-18 07:27:23', '2024-07-19 05:37:37', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
