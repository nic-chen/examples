
```shell
curl http://127.0.0.1:9180/apisix/admin/upstreams/up1  \
-H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -i -X PUT -d '
{
    "type":"roundrobin",
    "nodes":{
        "127.0.0.1:1980": 1
    }
}'
```


```shell
curl http://127.0.0.1:9180/apisix/admin/routes/1 \
-H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -i -d '
{
    "uri": "/products/*",
    "hosts": ["www.shop.com"],
    "plugins": {
      "limit-count": {
        "count": 2,
        "time_window": 60,
        "rejected_code": 503,
        "key": "remote_addr"
      },
      "basic-auth": {}
    }
    "upstream_id": "up1"
}'
```

```shell
curl http://127.0.0.1:9180/apisix/admin/consumers  \
-H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -i -d '
{
    "username": "jack",
    "plugins": {
        "basic-auth": {
            "username": "jack",
            "password": "123456"
        }
    }
}'
```


```shell
curl http://127.0.0.1:9180/apisix/admin/global_rules/1  \
-H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -i -X PUT -d '
{
    "plugins": {
      "prometheus": {}
    }
}'
```

