/*
 Navicat Premium Data Transfer

 Source Server         : linjun
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : music

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 15/07/2018 08:25:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for TBL_MUSIC
-- ----------------------------
DROP TABLE IF EXISTS `TBL_MUSIC`;
CREATE TABLE `TBL_MUSIC` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `SRC` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for TBL_RELATEMUSIC
-- ----------------------------
DROP TABLE IF EXISTS `TBL_RELATEMUSIC`;
CREATE TABLE `TBL_RELATEMUSIC` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `USER_ID` bigint(20) DEFAULT NULL,
  `MUSIC_ID` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for TBL_USER
-- ----------------------------
DROP TABLE IF EXISTS `TBL_USER`;
CREATE TABLE `TBL_USER` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `USER_NAME` varchar(32) DEFAULT NULL,
  `AGE` varchar(32) DEFAULT NULL,
  `ACCOUNT` varchar(32) DEFAULT NULL,
  `PASSWORD` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
