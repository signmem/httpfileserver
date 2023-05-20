# httpfileserver

> use to set golang read cfg.json template.
> log info into logfile.


# log vendor

>  确保 vendor/github.com/coreos/go-log/log/fields.go 文件被修改
> 1 修改 full_time 满足格式要求

```
"full_time":  time.Now().Format("2006-01-02 15:04:05.999"),  // time of log entry

```

> 2 修改 logger.verbose = true  由于属于内部变量无法外部修改

```
logger.verbose = true
```
> 3 logger example 

```
[2022-04-26 18:12:46.65] [DEBUG] [28062] [commands.go:33] >>> [main] msg=debug: yes
[2022-04-26 18:13:30.858] [DEBUG] [28224] [commands.go:33] >>> [main] msg=debug: yes
[2022-04-26 18:13:30.858] [INFO] [28224] [commands.go:33] >>> [main] msg=info: yes
[2022-04-26 18:13:30.858] [WARNING] [28224] [commands.go:33] >>> [main] msg=warning: yes
```

# download example 

> /api/v1/fileget   

```
curl http://10.189.20.49:8088/api/v1/fileget  -H 'Content-Type: application/json' -d '{"fspath":"20230519/binlog/backup_1.1.1.1_3306/mega.log","fstype":"temp","gfscluster":"gd15-ceph-mon-dbbackup"}'
```

> /api/v1/download  

```
curl http://10.189.20.49:8088/api/v1/download  -H 'Content-Type: application/json' -d '{"fspath":"20230519/binlog/backup_1.1.1.1_3306/mega.log","fstype":"temp","gfscluster":"gd15-ceph-mon-dbbackup"}'
```

## 说明  

> 经过验证, 两个接口下载文件内容 md5sum 一致   
> 区别在于 header set, 可以根据需要自己进行调整  
