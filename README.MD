# config 使用go读取json配置文件
- 使用JSON文件作为GO程序的配置文件
- 使用字符串和.对配置项深度获取


### 使用
> 简单获取
```
config := NewConfig("config_test.json")
d := config.Get("name")
```

> 深度获取值
```
config := NewConfig("config_test.json")
c := config.Get("database.host")
```

> 获取配置集合
```
config := NewConfig("config_test.json")
m := config.Get("database")
host := m.(map[string]interface{})["host"]
```
