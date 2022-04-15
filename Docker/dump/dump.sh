#!/bin/bash

DB_USER=${Mysql_User:-${MYSQL_ENV_DB_USER}}
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

DIR_PATH="`date +%Y/%m/%d/%H`"
TOTAL_PATH=${Local_Path}/mysqldump/${DIR_PATH}
if [ ! -d "${TOTAL_PATH}" ]; then
    mkdir -p ${TOTAL_PATH}
fi

databases=`mysql -u${DB_USER} -p${DB_PASS} -h${DB_HOST} -e "SHOW DATABASES;" | tr -d "| " | grep -v Database`
for db in $databases; do
    if [[ "$db" != "information_schema" ]] && [[ "$db" != "performance_schema" ]] && [[ "$db" != "mysql" ]] && [[ "$db" != _* ]] && [[ "$db" != "$IGNORE_DATABASE" ]]; then
        echo "Dumping database: $db"
        mysqldump --user="${DB_USER}" --password="${DB_PASS}" --host="${DB_HOST}" --port="${DB_PORT}" --databases $db > ${TOTAL_PATH}/$db.sql
    fi
done