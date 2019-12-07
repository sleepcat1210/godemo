package main
import (
	"strconv"
	"fmt"
	"time"
	"math/rand"
)
func main() {
  strings :=CreateNew()
  fmt.Println(strings)
  Shuffle(strings)
  fmt.Println(strings)
  fmt.Println("农民：0",Dispacther(0,strings))
  fmt.Println("农民：1",Dispacther(1,strings))
  fmt.Println("地主：2",Dispacther(2,strings))
  fmt.Println("底牌：3",Dispacther(3,strings))
  
}

//创建棋牌
func CreateNew()[]string{
	numbers :=make([]string,54)//构造一个大小为54的数组
	start :=0//造牌标记
	for i:=3;i<=16;i++{
		if i==16 {
			numbers[start] ="Q88"
			numbers[start+1] ="K99"//构造大小王
		} else {
			numbers[start] ="A" +strconv.Itoa(i)
			numbers[start+1] ="B" +strconv.Itoa(i)
			numbers[start+2] ="C" +strconv.Itoa(i)
			numbers[start+3] ="D" +strconv.Itoa(i)
			start +=4
		}
	}
	return numbers
}
//洗牌
func Shuffle(vals []string){
	r :=rand.New(rand.NewSource(time.Now().Unix()))//根据时间戳初始化random
	for len(vals) >0{
		n:=len(vals)//数组长度
		randIndex :=r.Intn(n)
		vals[n-1],vals[randIndex]=vals[randIndex],vals[n-1]
		vals=vals[:n-1]
	}
	
}
/**发牌
*order==0 玩家1次序
*order==1 玩家2次序
*order==2 玩家3次序
*order==3 底牌次序
*/
func Dispacther(order int,vals []string) []string{
	var playCards []string
	if order < 0 || order >3{//判断玩家次序
		return []string{}
	} else {
		size :=17 //默认总长度17
		if order == 3{
			size=3 //次序为3(底牌次序)时总长度为3
		}		
		for i:=0;i<size;i++{				
			 playCards = append(playCards,vals[order*17+i])//根据次序发牌
		}
	}	
	return playCards
}