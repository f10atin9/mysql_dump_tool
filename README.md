# mysql_dump_tool
A tool that periodically executes mysqldump operations and uploads backup files to QingStor object storage

## Quick Start
1. Set up connection and upload configuration

   In order to be able to connect to mysql-server and QingStor bucket correctly, please modify the [dump-configmap.yaml](K8s/dump-configmap.yaml) file according to your own configuration.

2. Configure cronJob

   Set the cronJob's timing schedule as needed, and configure the volume mount path

3. Deployment configMap and cronJob

    `kubectl apply -f K8s/dump-configmap.yaml K8s/dump-cm.yaml`

## Troubleshoot

1. ERROR 2059 (HY000): Authentication plugin 'caching_sha2_password' cannot be loaded: /usr/lib64/mysql/plugin/caching_sha2_password.so: cannot open shared object file: No such file or directory


```shell
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'PASSWORD'; FLUSH PRIVILEGES;
```
