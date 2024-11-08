

# kuboard-v4 部署

docker compose 
```yaml
services:
  mysql:
    image: mysql:8.0
    container_name: mysql_container
    environment:
      MYSQL_ROOT_PASSWORD: xxxx
      MYSQL_INITDB_SCRIPT: /docker-entrypoint-initdb.d/init.sql
      MYSQL_SQL_MODE: "STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION"
    ports:
      - "6606:3306"
    volumes:
      - ./data:/var/lib/mysql
      # - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    restart: always
  # kuboard-ldap-example:
  #   image: swr.cn-east-2.myhuaweicloud.com/kuboard/kuboard-v4-ldap-example:v4
  #   container_name: kuboard-ldap-example
  #   environment:
  #     - LDAP_URL=
  #     - LDAP_USERNAME=
  #     - LDAP_PASSWORD=
  #     - LDAP_BASE=ou=
  #   restart: always
  kuboard:
    image: swr.cn-east-2.myhuaweicloud.com/kuboard/kuboard:v4
    container_name: kuboard
    environment:
      TZ: "Asia/Shanghai"
      DB_DRIVER: com.mysql.cj.jdbc.Driver
      DB_URL: "jdbc:mysql://mysql_container:3306/kuboard?serverTimezone=Asia/Shanghai&useSSL=false&allowPublicKeyRetrieval=true"
      DB_USERNAME: xxxx
      DB_PASSWORD: xxxxx
    ports:
      - "8088:80"
    depends_on:
      - mysql
    restart: unless-stopped
```