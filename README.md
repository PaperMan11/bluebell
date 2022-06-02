# bluebell

## 项目部署
```
一、docker方式部署
1、运行容器 docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v ~/tan/docker/mysql:/var/lib/mysql -d mysql:8.0.19（根据models.create_tables自行创建表）
2、运行容器 docker run --name some-redis -d redis
3、修改 bluebell/conf/conf.yaml 文件中的 host 为容器的 NAMES
4、构建镜像 docker build . -t goweb_app
5、运行容器 docker run --link=mysql8019:mysql8019 --link=some-redis:some-redis -p 8080:8080 bubble_app
```