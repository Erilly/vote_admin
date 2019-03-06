# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.41)
# Database: vote
# Generation Time: 2019-03-06 08:14:02 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table vt_option
# ------------------------------------------------------------

DROP TABLE IF EXISTS `vt_option`;

CREATE TABLE `vt_option` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `selector_id` char(32) NOT NULL DEFAULT '' COMMENT '问题id',
  `options_id` char(32) NOT NULL DEFAULT '' COMMENT '选项id',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除状态：0在线，1删除',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='选项';



# Dump of table vt_question
# ------------------------------------------------------------

DROP TABLE IF EXISTS `vt_question`;

CREATE TABLE `vt_question` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question_id` char(8) NOT NULL DEFAULT '' COMMENT '问卷key',
  `page_cut` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否分页',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '问卷标题',
  `description` text NOT NULL COMMENT '问卷描述',
  `usertoken` varchar(32) NOT NULL COMMENT '所属用户',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='问卷';



# Dump of table vt_answer_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `vt_answer_log`;

CREATE TABLE `vt_answer_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `qustion_id` char(8) NOT NULL DEFAULT '' COMMENT '答卷id',
  `content` text NOT NULL COMMENT '答案结果',
  `token` varchar(32) NOT NULL DEFAULT '' COMMENT '答题客户端标识',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '答题用户id',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='回答结果';



# Dump of table vt_selector
# ------------------------------------------------------------

DROP TABLE IF EXISTS `vt_selector`;

CREATE TABLE `vt_selector` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question_id` char(8) NOT NULL DEFAULT '' COMMENT '问卷id',
  `selector_id` char(32) NOT NULL DEFAULT '' COMMENT '问题id',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `template_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '模板类型',
  `page` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '所属页码',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='问题';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
