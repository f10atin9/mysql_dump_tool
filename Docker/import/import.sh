#!/bin/bash

DB_USER=${Mysql_User:-${MYSQL_ENV_DB_USER}}
DB_PASS=${Mysql_Password:-${MYSQL_ENV_DB_PASS}}
DB_NAME=${Mysql_DBName:-${MYSQL_ENV_DB_NAME}}
DB_HOST=${Mysql_Host:-${MYSQL_ENV_DB_HOST}}
DB_PORT=${Mysql_Port:-${MYSQL_ENV_DB_PORT}}
ALL_DATABASES=${Mysql_AllDatabase}
SQL_Path=${VolumeDump_Path}/${Bucket_Path}

if [[ ${DB_USER} == "" ]]; then
	echo "Missing DB_USER env variable"
	exit 1
fi
if [[ ${DB_PASS} == "" ]]; then
	echo "Missing DB_PASS env variable"
	exit 1
fi
if [[ ${DB_HOST} == "" ]]; then
	echo "Missing DB_HOST env variable"
	exit 1
fi
if [[ ${ALL_DATABASES} == "" ]]; then
	if [[ ${DB_NAME} == "" ]]; then
		echo "Missing DB_NAME env variable"
		exit 1
	fi
	mysql --user="${DB_USER}" --password="${DB_PASS}" --port="${DB_PORT}"  --host="${DB_HOST}" "$@" "${DB_NAME}" < ${SQL_Path}/"${DB_NAME}".sql
else
	cd ${SQL_Path}
	databases=`for f in *.sql; do
    	printf '%s\n' "${f%.sql}"
	done`
for db in $databases; do
	  if [[ "$db" != "information_schema.sql" ]] && [[ "$db" != "performance_schema.sql" ]] && [[ "$db" != "mysql.sql" ]] && [[ "$db" != _* ]]; then
	      echo "Importing database: $db"
	      mysql --user="${DB_USER}" --password="${DB_PASS}" --port="${DB_PORT}"  --host="${DB_HOST}" "$@" "$db" < ${SQL_Path}/$db.sql
	  fi
done
fi