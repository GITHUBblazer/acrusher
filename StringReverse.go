package  main

import "fmt"

func swap(ch []rune,i,j int){
	temp:=ch[i]
	ch[i] = ch[j]
	ch[j] = temp
}
func main(){
	str:= "Hello, world!"

	strArr := []rune(str)
	strArrLen := len(strArr)
	for i:=0;i<strArrLen/2;i++{ //前后对半 交换
		swap(strArr,i,strArrLen-1-i)
	}
    fmt.Println(string(strArr))
}
