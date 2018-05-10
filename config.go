package lconfig

import (
	"bytes"
	"sync"
	"os"
	"bufio"
	"io"
	"fmt"
	"strings"
	"strconv"
)

var(
	DEFAULT_SECTION = "default"
	DEFAULT_COMMENT = []byte{'#'}
	SECTION_BEGIN = []byte{'['}
	SECTION_END = []byte{']'}
)


type LConfigInterface interface{
	String(key string) string
	Strings(key string,split ...string) ([]string,error)
	Int(key string)(int,error)
	Int64(key string)(int64,error)
	Bool(key string)(bool,error)
	Float64(key string)(float64,error)
	Sections(key string)(map[string]string,error)
}

type LConfig struct{
	sync.RWMutex
	data map[string]map[string]string
}

func NewConfig(confFile string)(LConfigInterface,error){
	//初始化
	c := &LConfig{
		data : make(map[string]map[string]string),
	}
	//读取指定的配置文件,并解析
	err := c.parse(confFile)
	return c,err
}

//解析指定的配置文件
func (c *LConfig) parse(confFile string)error{
	c.Lock()
	defer c.Unlock()
	f,err := os.OpenFile(confFile,os.O_RDONLY,0666)
	if err != nil{
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var section string
	for{
		line,_,err := reader.ReadLine()
		if err == io.EOF{
			break
		}else if bytes.Equal(line,[]byte{}){
			continue
		}else if err != nil{
			return err
		}
		line = bytes.TrimSpace(line)
		switch  {
		case bytes.HasPrefix(line,DEFAULT_COMMENT):
			//注释
			continue
		case bytes.HasPrefix(line,SECTION_BEGIN) && bytes.HasSuffix(line,SECTION_END):
			//section
			section = string(line[1:len(line)-1])
		default:
			//key=value值
			values := bytes.Split(line,[]byte{'='})
			if len(values) != 2{
				return fmt.Errorf("parse %s content has error.line:%s",confFile,string(line))
			}
			key := bytes.TrimSpace(values[0])
			value := bytes.TrimSpace(values[1])
			c.addConfig(section,string(key),string(value))
		}
	}

	return nil
}

func (c *LConfig) addConfig(section,key,value string){
	if section == ""{
		section = "default"
	}
	if _,ok := c.data[section];!ok{
		c.data[section] = make(map[string]string)
	}
	//_,ok := c.data[section][key]
	c.data[section][key] = value
}



func (c *LConfig) String(key string)string{
	return c.get(key)
}

func (c *LConfig) Strings(key string,split ...string)([]string,error){
	var sp = ","
	value := c.get(key)
	if value == ""{
		return nil,fmt.Errorf("can't find key:%s",key)
	}
	if len(split) > 0{
		sp = split[0]
	}
	return strings.Split(value,sp),nil
}

func (c *LConfig) Int(key string)(int,error){

	value := c.get(key)
	if value != ""{
		ret,err := strconv.Atoi(value)
		if err == nil{
			return ret,nil
		}
		return 0,err
	}

	return 0,fmt.Errorf("can't find key:%s",key)

}

func (c *LConfig) Int64(key string)(int64,error){

	value := c.get(key)
	if value != ""{
		ret,err := strconv.ParseInt(value,10,64)
		if err == nil{
			return ret,nil
		}
		return 0,err
	}

	return 0,fmt.Errorf("can't find key:%s",key)
}

func (c *LConfig) Bool(key string)(bool,error){

	value := c.get(key)
	if value != ""{
		ret,err := strconv.ParseBool(value)
		if err == nil{
			return ret,nil
		}
		return false,err
	}

	return false,fmt.Errorf("can't find key:%s",key)
}

func (c *LConfig) Float64(key string)(float64,error){

	value := c.get(key)
	if value != ""{
		ret,err := strconv.ParseFloat(value,64)
		if err == nil{
			return ret,nil
		}
		return 0.0,err
	}
	return 0,fmt.Errorf("can't find key:%s",key)
}

func (c *LConfig) Sections(section string)(map[string]string,error){
	ret,ok := c.data[section]
	if !ok{
		return nil,fmt.Errorf("can't find section:%s",section)
	}
	return ret,nil
}


func (c *LConfig) get(key string) string{
	var(
		section string
		option string
	)
	keys := strings.Split(strings.TrimSpace(key),"::")
	if len(keys) == 2{
		section = keys[0]
		option = keys[1]
	}else{
		section = DEFAULT_SECTION
		option = key
	}
	if value,ok := c.data[section][option];ok{
		return value
	}
	return ""
}