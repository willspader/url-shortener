### Mysql DB Initialization

```
docker run --name url_shortener_mysql_cluster -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=url_shortener_mysql_database  -e MYSQL_USER=url_shortener_app -e MYSQL_PASSWORD=url_shortener_pwd -d mysql:8.0
```
