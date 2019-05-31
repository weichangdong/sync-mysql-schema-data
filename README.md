## sync-mysql-schema-data
基于 [mysql-schema-sync](github.com/hidu/mysql-schema-sync) 开发的同步数据的工具.
### 使用
#### 多数用法,参看[mysql-schema-sync](github.com/hidu/mysql-schema-sync)的说明
```
Usage of ./main:
  -conf string
    	json config file path (default "./config.json")
  -dest string
    	mysql dsn dest,eg test@(127.0.0.1:3306)/imis
  -drop
    	drop fields,index,foreign key
  -mail_to string
    	overwrite config's email.to
  -source string
    	mysql dsn source,eg: test@(10.10.0.1:3306)/test
    		when it is not empty ignore [-conf] param
  -sync
    	sync shcema change to dest db
  -sync_data
    	sync source db table data  to dest db table (default true)
  -sync_data_truncate
    	is need truncate  source db table data  to dest db table
  -tables string
    	table names to check
    		eg : product_base,order_*
  -tables_ignore string
    	table names to ignore check
    		eg : product_base,order_*

mysql schema && data sync tools 0.3
Base On https://github.com/hidu/mysql-schema-sync/
```
#### `sync_data` ./main -sync_data=true,则表示这个操作是同步数据.否则就是同步数据结构.
#### `sync_data_truncate` ./main -sync_data_truncate=true,表示同步源数据的时候,是否truncate本地的数据,没有备份哦,操作需谨慎. 如果不为true,则同步数据的时候,如果目标的数据表,有自增的属性,则id的值是null,否则还是保留原有的id插入.
#### 配置项里面的`sync_data_tables`,指定需要同步数据的数据表.

```
"sync_data_tables":["user_e_trans","staff_loan_data"],
```
#### 另外说明,这个工具自己随手写的,可能比较粗糙,最近也特忙,可能有一些问题.如果使用的时候,遇到了,可以留言,我会处理的.
#### QQ:1694669

