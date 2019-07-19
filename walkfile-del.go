import (
package main
"fmt"
"os"
// "io/ioutil"
"path/filepath"
"regexp"
)
var count int =0
func walkFn(path string ,info os.FileInfo,err error) error {
count++
//  字符是否以 ._ 开头？
	reg, err := regexp.Compile(`^\._`)
	if nil != err {
		fmt.Println(err)
		
	}
if reg.MatchString(filepath.Base(path))==true{
	fmt.Println(reg.MatchString(filepath.Base(path)),filepath.Base(path))

	//目录删除
	err1 := os.Remove(path)
if err1 != nil{
	fmt.Println("删除目录失败！",err1)
}else{

	fmt.Println("目录删除成功！",info.Name(),err1,count)
}

}


return err
}

func main(){

path1:="E:/"
// 遍历
filepath.Walk(path1,walkFn)

}


