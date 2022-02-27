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

ports are unnder /go-liziwei01-appui/httpapi package\
now default ip:port is \
localhost:8080\
\

For crawler\
use\
http://localhost:8080/crawlPaper?title=SRCNN
to crawl papers related with SRCNN

use\
http://localhost:8080/paperList?startTime=1&ref=SRCNN
to access the SRCNN papers

use\
http://localhost:8080/paper?title=2111.15185.pdf
to access exact paper


For Project ERG3020\
paper acccess port is /paperList\
use the url\
http://localhost:8080/paperList?
to access

|get params|comment|require|
| --------- | --------- | --------- |
|pageIndex|1|no|
|pageLength|10|no|
|authors|emp|no|
|title|title|no|
|journal|journal name|no|
|startTime|timestamp: 1618761600 means 2021-04-19_0:00|no|
|endTime|24:00|no|
|type|?type=author&key=he and authors=he is equal|no|
|key||no|

eg：

```bash
{
    "data": {
        "count":1,
        "list":[
            {
            "authors":"liziwei01",
            "journal":"ieee",
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
/userList\
|postParams|comment|require|
|name|user name|yes|
|pageIndex|1|no|
|pageLength|10|no|

/insertUser\
|postParams|comment|require|
|name|user name|yes|
|ID|user id|yes|
|password|user pwd|yes|

## Contributing

go-liziwei01-library refers baidu_gdp

## Database

use\
mysql -uwork -pliziwei01 -h10.30.202.213 -P3306 db_liziwei01\
to log in mysql\
\
CSC3170 table:\
CREATE TABLE \`tb_star_user_info\` (\
        \`user_id\` int unsigned NOT NULL DEFAULT '0' COMMENT 'index_number, primary key',\
        \`name\` varchar(256) NOT NULL DEFAULT '' COMMENT 'user name',\
        \`password\` varchar(256) NOT NULL DEFAULT '' COMMENT 'user password',\
        PRIMARY KEY (\`user_id\`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='明星查询平台用户信息表';\
\
ERG3020 table：\
CREATE TABLE \`tb_gesture_teleoperation_paper_info\` (\
        \`index_number\` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'auto increment primary key',\
        \`title\` varchar(256) NOT NULL DEFAULT '' COMMENT 'title',\
        \`author\` varchar(1024) NOT NULL DEFAULT '' COMMENT 'separated by comma',\
        \`publish_time\` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'time stamp',\
        \`journal\` varchar(256) NOT NULL DEFAULT '' COMMENT 'publisher',\
        \`ref\` varchar(10240) NOT NULL DEFAULT '' COMMENT 'separated by comma',\
        \`total_cites\` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'total cites',\
        \`score\` int unsigned NOT NULL DEFAULT '0' COMMENT 'score',\
        \`content\` text  COMMENT 'content',\
        PRIMARY KEY (\`index_number\`),\
        KEY \`idx_title\` (\`title\`),\
        KEY \`idx_author\` (\`author\`),\
        KEY \`idx_journal\` (\`journal\`)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tb_gesture_teleoperation_paper_info';

## License

MIT License
