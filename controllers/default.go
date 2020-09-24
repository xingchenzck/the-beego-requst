package controller
import (
	"encoding/json"
	"fmt"
	"frist/models"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	//请求数据
	user := c.Ctx.Input.Query("user")
	psd := c.Ctx.Input.Query("psd")
	//访问数据库
	//eg:使用固定数据进行校验
	if user != "admin" || psd != "123456" {
		//处理错误
		c.Ctx.ResponseWriter.Write([]byte("===错误==="))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("成功"))
	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}


//
//func (c *MainController) Post(){
//	//解析json格式的文件
//	var person models.Person
//	dateBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
//	if err!= nil {
//		c.Ctx.WriteString("接收错误")
//		return
//	}
//	err = json.Unmarshal(dateBytes,&person)
//	if err!= nil {
//		c.Ctx.WriteString("解析错误")
//		return
//	}
//	fmt.Println("姓名",person.Name)
//	fmt.Println("年龄",person.Age)
//	fmt.Println("性别",person.Sex)
//	c.Ctx.WriteString("访问成功")
//}

func (c *MainController) Post(){
	var derson models.Derson
	 dataBtes ,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	 if err != nil {
	 	c.Ctx.WriteString("开始接受错误")
		 return
	 }
	 err =json.Unmarshal(dataBtes,&derson)
	if err !=nil {
		c.Ctx.WriteString("开始解析错误")
		return
	}
	fmt.Println("姓名",derson.Name)
	fmt.Println("生日",derson.Birthday)
	fmt.Println("地址",derson.Address)
	fmt.Println("尼克",derson.Nick)
	 c.Ctx.WriteString("你访问成功")

}

