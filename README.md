# go-liziwei01-appui

This is a school project written by liziwei from CUHK(SZ)

## Install

Just git pull and go run main.go\
在/go-liziwei01-appui/目录下使用\
go run main.go\
编译并运行APP

## Usage


域名和接口设置在/go-liziwei01-appui/httpapi package下\
目前测试版本ip:port设定为\
localhost:8086\

\
For Project ERG3020\
论文访问接口为/paperList\
启动APP后使用以下url访问服务\
http://localhost:8086/paperList?\
\
get参数支持       说明                                                  是否必传\
pageIndex       默认1                                                    否\
pageLength      默认10                                                   否\
authors         默认空                                                   否\
startTime       默认当天0:00，时间戳形式。1618761600表示2021-04-19_0:00      否\
endTime         默认当天24:00                                             否\
\
搜索示例：\
http://10.30.202.213.:8086/paperList?authors=liziwei&startTime=1618761600

For Project CSC3170\
用户访问接口为/userList\
post参数支持     说明      是否必传\
name            用户名称     是\
pageIndex       默认1       否\
pageLength      默认10      否\
\
用户插入接口为/insertUser\
post参数支持     说明      是否必传\
name          用户名称     是\
ID            用户id      是\
password      用户密码     是\

## Contributing

library内容参考了gdp

## Database

使用\
mysql -uwork -pliziwei01 -h10.30.202.213 -P3306 db_liziwei01\
远程登录\
\
CSC3170数据表已创建\
CREATE TABLE `tb_star_user_info` (\
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT 'index_number, primary key',\
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT 'user name',\
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT 'user password',\
  PRIMARY KEY (`user_id`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='明星查询平台用户信息表';\
\
ERG3020数据表已创建\
\
建表语句：\
CREATE TABLE `tb_gesture_teleoperation_paper_info` (\
  `index_number` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'index_number, primary key',\
  `title` varchar(256) NOT NULL DEFAULT '' COMMENT '',\
  `author` varchar(256) NOT NULL DEFAULT '' COMMENT 'json',\
  `publish_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '',\
  `journal` varchar(256) NOT NULL DEFAULT '' COMMENT 'journal',\
  `references` varchar(10240) NOT NULL DEFAULT '' COMMENT 'references',\
  `point` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'point',]\
  PRIMARY KEY (`index_number`),\
  KEY `idx_title` (`title`),\
  KEY `idx_author` (`author`),\
  KEY `idx_publish_time` (`publish_time`),\
  KEY `idx_journal` (`journal`),\
\
  KEY `idx_title_point` (`title`,`point`),\
  KEY `idx_author_point` (`author`,`point`),\
  KEY `idx_journal_point` (`journal`,`point`),\
\
  KEY `idx_title_publish_time_point` (`title`,`publish_time`,`point`),\
  KEY `idx_author_publish_time_point` (`title`,`publish_time`,`point`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tb_gesture_teleoperation_paper_info';

## License

```
