/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : localhost
 Source Database       : student

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : utf-8

 Date: 07/16/2017 06:44:56 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `Student`
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `id` varchar(11) NOT NULL COMMENT '学生编号',
  `class_num` varchar(11) NOT NULL COMMENT ‘班级编号’,
  `score` varchar(11) NOT NULL COMMENT '分数',
  UNIQUE (`Id`),
  INDEX `student_id` (`id`) comment ''
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
--  Table structure for `class`
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `class_num` varchar(11)  NOT NULL COMMENT '班级编号',
  `teacher_name` varchar(30) NOT NULL COMMENT '老师姓名',
  PRIMARY KEY (`class_num`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
