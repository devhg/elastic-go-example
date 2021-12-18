# elastic-go-example

### step1.

```shell
go generate ./...
go build -o esgo ./cmd && ./esgo
```

### step2.

you can see

```shell
[GIN-debug] POST   /api/user/create          --> github.com/devhg/es/internal/server/handler.Create (3 handlers)
[GIN-debug] PUT    /api/user/update          --> github.com/devhg/es/internal/server/handler.Update (3 handlers)
[GIN-debug] DELETE /api/user/delete          --> github.com/devhg/es/internal/server/handler.Delete (3 handlers)
[GIN-debug] GET    /api/user/info            --> github.com/devhg/es/internal/server/handler.MGet (3 handlers)
[GIN-debug] POST   /api/user/search          --> github.com/devhg/es/internal/server/handler.Search (3 handlers)
```

1. 创建客户端

./internal/data.go

3. 创建index和mapping

./internal/biz/useres.go

4. 批量添加

```shell
$ curl --location --request POST 'http://localhost:8080/api/user/create' \
    --header 'Content-Type: application/json' \
    --data-raw '[
    {
        "id": 1,
        "username": "a1",
        "nickname": "golang",
        "phone": "11111111111",
        "age": 20,
        "ancestral": "北京市",
        "identity": "工人"
    },
    {
        "id": 2,
        "username": "b2",
        "nickname": "golang",
        "phone": "22222222222",
        "age": 20,
        "ancestral": "北京市",
        "identity": "工人"
    },
    {
        "id": 3,
        "username": "c3",
        "nickname": "golang",
        "phone": "33333333333",
        "age": 20,
        "ancestral": "北京市",
        "identity": "工人"
    }
]'
```

验证结果

```shell
$ curl --location --request GET 'http://localhost:9200/devhg_elastic_go/_search'
```

5. 批量更新

```shell
$ curl --location --request PUT 'http://localhost:8080/api/user/update' \
    --header 'Content-Type: application/json' \
    --data-raw '[{
        "id": 1,
        "username": "a1",
        "nickname": "a1-update",
        "phone": "11111111111",
        "age": 20,
        "ancestral": "山东省",
        "identity": "工人"
    }]'
```

验证结果
```shell
$ curl --location --request GET 'http://localhost:8080/api/user/info?id=1,2,3'
```



6. 批量删除

```shell
$ curl --location --request DELETE 'http://localhost:8080/api/user/delete' \
    --header 'Content-Type: application/json' \
    --data-raw '[
    {
        "id": 1
    },
    {
        "id": 2
    },
    {
        "id": 3
    }]'
```

验证结果
```shell
$ curl --location --request GET 'http://localhost:9200/devhg_elastic_go/_search'
```


7. 查询
这里进行说明一下，使用json来选择不同的条件，需要那个条件就填写json就好了。这个测试的查询条件就是查找出籍贯是北京市的用户列表，通过num、size限制查询数据量，即第一页，数据量为10。
可以通过修改这三个字段来验证结果
```shell
$ curl --location --request POST 'http://localhost:8080/api/user/search' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "nickname": "",
        "phone": "",
        "identity": "",
        "ancestral": "北京市",
        "num": 1,
        "size":10
    }'
```

8. 批量查询

```shell
$ curl --location --request GET 'http://localhost:8080/api/user/info?id=1,2,3'
```