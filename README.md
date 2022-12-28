### Technical Solution

This was implemented only for learning some of the golang ecosystem, such as syntax, manual dependency injection, relational database communication, etc.

##### Important

Although the system works and does what it should be done, it can not be used as a real implementation of a URL Shortener system:

- It gets an auto-generated id from the MySQL database and turns it into a base 62 representation using an ASCII table without randomization.
  - Example: Id 1 from database will get 1 as short url. Id 11 from database will get 1A as short url. Id 195 from database will get 39 as short url. Id 196 from database will get 3A as short url. Always following 1, 2, 3, 4, 5, 6, 7, 8, 9, 1A, 1B, ..., 1a, 1b, ..., 20, 21, ..., and so on.
- There is no verification for malicious URL.
- There is no implementation against BOT/SPAM.
- There is no checksum in the algorithm to convert decimal to base62 representation.
- There is no verification against SQL injection.

### Mysql DB Container Initialization

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
