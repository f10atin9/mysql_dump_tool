# mysql_dump_tool
A tool to backup mysql

## Backup Database

#### Specify database
```shell
docker run --rm -v $PWD:/mysqldump -e DB_PASS=PASSWORD -e DB_USER=USER -e DB_HOST=HOST -e DB_NAME=DB_NAME -e DB_PORT=3306  f10atin9/dump
```

#### All
```shell
docker run --rm -v $PWD:/mysqldump -e DB_PASS=PASSWORD -e DB_USER=USER -e DB_HOST=HOST -e ALL_DATABASES=true -e DB_PORT=3306  f10atin9/import
```

mysqldump -e DB_PASS=P@2s50rd -e DB_USER=f10atin9 -e DB_HOST=192.168.0.5 -e ALL_DATABASES=true -e DB_PORT=3306  f10atin9/dump_dir

## upload sql to Qing-stor bucket