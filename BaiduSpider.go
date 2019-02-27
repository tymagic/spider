package main
//并发爬虫读取到文件，棒棒哒
import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)
var ch = make(chan int)
func spider(pagenum int,ch chan int){
	num:=(pagenum-1)*50
	url:="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa(num)
	resp,err:=http.Get(url)
	if err!=nil{
		return
	}
	defer resp.Body.Close()
	var tmp string
	buf:=make([]byte,4*1024)
	for {
		n,err :=resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read",err)
			break
		}
		tmp += string(buf[:n])
	}
	//fmt.Println("tmp = ", tmp)
	WriteFile(pagenum,tmp)
	ch <- pagenum
}
//获取拼接url
//爬取

//写入文件
func WriteFile(Page int,result string) {
	filename:=strconv.Itoa(Page)+".html"
	file,err1:=os.Create(filename)
	if err1!=nil{
		fmt.Println("创建文件错误",err1)
	}
	defer file.Close()
	file.WriteString(result)

}

func main() {

	//url:=
	var start,end int
	fmt.Print("请输入起始页码:")
	fmt.Scanf("%d",&start)
	fmt.Print("请输入结束页码:")
	fmt.Scanf("%d",&end)
	fmt.Println(start)
	fmt.Println(end)
	//result,err1:=spider(url)
	//fmt.Println(result,err1)
	//按照百度贴吧规则生成
	for i:=start;i<=end;i++{
		go spider(i,ch)

	}
	for i:=start;i<=end;i++{
		fmt.Printf("%d页爬完了",<-ch)
	}


}


