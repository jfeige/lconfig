package lconfig

import (
	"testing"
	"fmt"
)

func TestGet(t *testing.T){
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
	//default::host
	if v := config.String("host");v == ""{
		t.Error("Can't find default key host\n")
	}else{
		fmt.Printf("Get value for default key host sucess value:%s\n",v)
	}

	//redis::port
	if v,err := config.Int("redis::port");err != nil{
		t.Errorf("Get value for redis key port has error:%v\n",err)
	}else{
		fmt.Printf("Get value for redis key port sucess value:%d\n",v)
	}

	//redis::host
	if v := config.String("redis::host");v == ""{
		t.Error("Can't find default key host\n")
	}else{
		fmt.Printf("Get value for redis key host sucess value:%s\n",v)
	}

	//redis::master bool
	if v,err := config.Bool("redis::master");err != nil{
		t.Errorf("Get value for redis key master has error:%v\n",err)
	}else{
		fmt.Printf("Get value for redis key master sucess value:%t\n",v)
	}

	//mysql slave_addr []string
	if v,err := config.Strings("mysql::slave_addr",";");err != nil{
		t.Errorf("Get value for mysql key slave_addr has error:%v\n",err)
	}else{
		fmt.Printf("Get value for mysql key slave_addr sucess value:%v\n",v)
	}

	if v,err := config.Sections("province");err != nil{
		t.Errorf("Get sections for province has error:%v\n",err)
	}else{
		fmt.Printf("Get sections for province sucess value:%v\n",v)
	}


}
