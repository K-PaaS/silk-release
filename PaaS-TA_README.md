## 01. Check the commit contents
> [enable cce](https://github.com/PaaS-TA/silk-release/commit/60db690221b2570cab9721a3e6a18fab7ce92a72)

## 02. Submodule modify : lib/pq (src/github.com/lib/pq)
> pull v1.10.0 lib/pq
``` 
$ rm -rf src/github.com/lib/pq
$ mkdir -p src/github.com/lib/pq
$ git clone https://github.com/lib/pq.git -b v1.10.0 src/github.com/lib/pq

$ rm -rf src/code.cloudfoundry.org/silk/vendor/github.com/lib/pq
$ mkdir -p src/code.cloudfoundry.org/silk/vendor/github.com/lib/pq
$ git clone https://github.com/lib/pq.git -b v1.10.0 src/code.cloudfoundry.org/silk/vendor/github.com/lib/pq
```

## 03. Submodule modify : go-sql-driver/mysql (src/github.com/go-sql-driver/mysql)
> pull v1.4.1 go-sql-driver/mysql
``` 
$ rm -rf src/github.com/go-sql-driver/mysql
$ mkdir -p src/github.com/go-sql-driver/mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.4.1 src/github.com/go-sql-driver/mysql

$ rm -rf src/code.cloudfoundry.org/silk/vendor/github.com/go-sql-driver/mysql
$ mkdir -p src/code.cloudfoundry.org/silk/vendor/github.com/go-sql-driver/mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.4.1 src/code.cloudfoundry.org/silk/vendor/github.com/go-sql-driver/mysql
```