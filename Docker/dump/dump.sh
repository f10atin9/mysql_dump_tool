#!/bin/bash

DB_USER=${Mysql-User:-${MYSQL_ENV_DB_USER}}
DB_PASS=${Mysql_Password:-${MYSQL_ENV_DB_PASS}}
DB_NAME=${Mysql_DBName:-${MYSQL_ENV_DB_NAME}}
DB_HOST=${Mysql_Host:-${MYSQL_ENV_DB_HOST}}
DB_PORT=${Mysql_Port:-${MYSQL_ENV_DB_PORT}}
ALL_DATABASES=${Mysql_AllDatabase}
IGNORE_DATABASE=${IGNORE_DATABASE}
Local_Path=${VolumeDump_Path}

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

DIR_NAME="`date +%Y%m%d%H`"
mkdir /mysqldump/${DIR_NAME}

if [[ ${ALL_DATABASES} == "" ]]; then
        if [[ ${DB_NAME} == "" ]]; then
                echo "Missing DB_NAME env variable"
                exit 1
        fi
        mysqldump --user="${DB_USER}" --password="${DB_PASS}" --host="${DB_HOST}" --port="${DB_PORT}" "$@" "${DB_NAME}" > /mysqldump/${DIR_NAME}/"${DB_NAME}".sql
else
        databases=`mysql --user="${DB_USER}" --password="${DB_PASS}" --host="${DB_HOST}" -e "SHOW DATABASES;" | tr -d "| " | grep -v Database`
for db in $databases; do
    if [[ "$db" != "information_schema" ]] && [[ "$db" != "performance_schema" ]] && [[ "$db" != "mysql" ]] && [[ "$db" != _* ]] && [[ "$db" != "$IGNORE_DATABASE" ]]; then
        echo "Dumping database: $db"
        mysqldump --user="${DB_USER}" --password="${DB_PASS}" --host="${DB_HOST}" --port="${DB_PORT}" --databases $db > /mysqldump/${DIR_NAME}/$db.sql
    fi
done
fi