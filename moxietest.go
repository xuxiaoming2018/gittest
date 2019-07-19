package main 
	import "fmt"
	// 把要默写的英文单词与对应的汉字放入二维数组中
   var n int=3
var a3 [3]string

	var a2 = [3][3]string{

{"nodes","[nəuds]","n. 节点（node的复数形式）"},
{"snake","[sneɪk]","蛇"},
{"render","['rendə]","渲染着色"},


}          
	func main (){
		var word string 
		var name string
		fmt.Println("请输入您的姓名,输入完请按一次回车键")
				fmt.Scanln(&name)
 
		for i := 0; i <=n-1; i++ {

			

				fmt.Println(name,"跟根据中文提示输入英文单词：",a2[i][2])
				fmt.Scanln(&word)
				var value string
				value=a2[i][0]

				// 第1次开始判断：否
					if word != value{
					fmt.Println("你写错了,请你再好好想想！" )
					fmt.Println("请重新输入",a2[i][2],":的英文单词")
					fmt.Scanln(&word)

					// 第2次开始判断：否
					if word != value{
                    // 将写错的单词放入数组a3中
                    
                 
                 a3[i]= value
                  fmt.Println(word,value)
     



					fmt.Println(name,"又写错了,正确答案：",a2[i][0],a2[i][2],a2[i][1])
					
				   }else{
					fmt.Println(name,"回答正确！",a2[i][2],a2[i][1])
				
                    }







                    if i == n-1 {
                        fmt.Println("不对的单词： ",a3)
						fmt.Println("单词默写完毕，自动退出程序！")
						return
					}


			    fmt.Println( )

			}
 }
}    

 