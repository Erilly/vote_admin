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



CREATE TABLE `vt_question` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question_id` char(8) NOT NULL DEFAULT '' COMMENT '问卷key',
  `page_cut` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否分页',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '问卷标题',
  `description` text NOT NULL COMMENT '问卷描述',
  `usertoken` varchar(32) NOT NULL COMMENT '所属用户',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0正常 1删除',
  `publish_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '发布状态：0上线 1下线',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='问卷';



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