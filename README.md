### Mysql DB Initialization

```
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=url_shortener_db -e MYSQL_USER=url_shortener_app -e MYSQL_PASSWORD=url_shortener_pwd mysql
```

## Create Table

```
CREATE TABLE IF NOT EXISTS URL_SHORTENER
(
	ID BIGINT NOT NULL AUTO_INCREMENT,
	LONG_URL varchar(2048),
	PRIMARY KEY (ID)
)
```

### Run App

```
go run cmd/url-shortener/start.go
```
