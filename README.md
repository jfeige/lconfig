# lconfig
	一款很简单的读取ini配置文件，后面会继续完善，支持其他配置文件格式

## 安装
	go get github.com/jfeige/lconfig

## 使用
	配置文件格式参考myconfig.ini

```
	config,err := NewConfig("myconfig.ini")
	if err != nil{
		t.Errorf("Load configfile has error:%v",err)
		return
	}
	//default::port
	if v,err := config.Int("port");err != nil{
		t.Errorf("Get value for default key port has error:%v\n",err)
	}else{
		fmt.Printf("Get value for default key port sucess value:%d\n",v)
	}

	//redis::port
	if v,err := config.Int("redis::port");err != nil{
		t.Errorf("Get value for redis key port has error:%v\n",err)
	}else{
		fmt.Printf("Get value for redis key port sucess value:%d\n",v)
	}
```

## API列表

```
	String(key string) string
	Strings(key string,split ...string) ([]string,error)
	Int(key string)(int,error)
	Int64(key string)(int64,error)
	Bool(key string)(bool,error)
	Float64(key string)(float64,error)
	Sections(key string)(map[string]string,error)
```
