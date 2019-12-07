package main
import (
	"github.com/wqtapp/poker"
	"fmt"
	"sync"
	"time"
	"math/rand"
)
type User struct{
	sync.RWMutex                         //操作playNum以及player时加锁
	pokerCards poker.PokerSet        //当前游戏中的所有的牌
	card []poker.PokerSet
	bottomCards poker.PokerSet
}
func main() {
	var user User
	user.pokerCards =poker.CreateDeck().ToPokerSet()
	rand.Seed(time.Now().Unix())
	for i := len(user.pokerCards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		user.pokerCards[i], user.pokerCards[num] = user.pokerCards[num], user.pokerCards[i]
	}
	fmt.Println(user.pokerCards)	
	user.Lock()
	for i,cards := range user.pokerCards{
		if i > 50{
			break
		}
		 shang := i/17
		// fmt.Println(cards)
		// fmt.Println(shang)
		user.card[shang] = append(user.card[shang],cards)
	}
	user.Unlock()
	// user.Card=card
	// fmt.Println(user.Card)
	// card.SortAsc()
	// for _,v :=range user.Card{
	// 	// for _,val :=range v.Card{
	// 		fmt.Println(v)
	// 	// }		
	// }
}