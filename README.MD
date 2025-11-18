

进入postgresql docker

docker exec -it postgresql sh

再进入 postgresql

psql -U postgres


\l                # 查看有哪些数据库
\c simple_admin   # 切换到 simple_admin
\dt               # 列出数据库中的所有表
\d users          # 查看表结构
\q                # 退出 psql


-------------------------
 CREATE DATABASE aqua;

 