# go-liziwei01-appui
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a school project written by liziwei from CUHK(SZ)

## Clone & Install Hooks

git clone https://github.com/liziwei01/go-liziwei01-appui.git &&\
wget https://github.com/liziwei01/hooks/archive/refs/tags/1.0.tar.gz &&\
tar -xzvf 1.0.tar.gz &&\
mv hooks-1.0/commit-msg go-liziwei01-appui/.git/hooks &&\
rm -rf hooks-1.0 &&\
rm 1.0.tar.gz &&\
cd go-liziwei01-appui

## Run

go run main.go

## Use

域名和接口设置在/go-liziwei01-appui/httpapi package下\
目前测试版本ip:port设定为\
localhost:8086\
\
For Project ERG3020\
论文访问接口为/paperList\
启动APP后使用以下url访问服务\
http://localhost:8086/paperList?\

|get参数支持|说明|是否必传|
| --------- | --------- | --------- |
|pageIndex|默认1|否|
|pageLength|默认10|否|
|authors|默认空|否|
|title|标题|否|
|journal|期刊名|否|
|startTime|默认当天0:00，时间戳形式。1618761600表示2021-04-19_0:00|否|
|endTime|默认当天24:00|否|
|type|?type=author&key=he和authors=he是等价的|否|
|key||否\

返回示例：

```bash
{
    "data": {
        "count":1,
        "list":[
            {
            "authors":"liziwei01",
            "journal":"nature",
            "point":96,
            "publish_time":"2021-04-30_23:59:59",
            "references":"maybe should be sth here",
            "title":"research gesture teleoperation"
            }
        ]
    },
    "errno": 0,
    "errmsg": "success"
}
```

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
password      用户密码     是

## Contributing

library内容参考了gdp

## Database

使用\
mysql -uwork -pliziwei01 -h10.30.202.213 -P3306 db_liziwei01\
远程登录\
\
CSC3170数据表已创建\
\
CREATE TABLE \`tb_star_user_info\` (\
        \`user_id\` int unsigned NOT NULL DEFAULT '0' COMMENT 'index_number, primary key',\
        \`name\` varchar(256) NOT NULL DEFAULT '' COMMENT 'user name',\
        \`password\` varchar(256) NOT NULL DEFAULT '' COMMENT 'user password',\
        PRIMARY KEY (\`user_id\`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='明星查询平台用户信息表';\
\
ERG3020数据表已创建\
\
建表语句：\
CREATE TABLE \`tb_gesture_teleoperation_paper_info\` (\
        \`index_number\` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'auto increment primary key',\
        \`title\` varchar(256) NOT NULL DEFAULT '' COMMENT 'title',\
        \`author\` varchar(1024) NOT NULL DEFAULT '' COMMENT 'separated by comma',\
        \`publish_time\` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'time stamp',\
        \`journal\` varchar(256) NOT NULL DEFAULT '' COMMENT 'publisher',\
        \`ref\` varchar(10240) NOT NULL DEFAULT '' COMMENT 'separated by comma',\
        \`total_cites\` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'total cites',\
        \`score\` int unsigned NOT NULL DEFAULT '0' COMMENT 'score',\
        PRIMARY KEY (\`index_number\`),\
        KEY \`idx_title\` (\`title\`),\
        KEY \`idx_author\` (\`author\`),\
        KEY \`idx_journal\` (\`journal\`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tb_gesture_teleoperation_paper_info';

## License

MIT License
