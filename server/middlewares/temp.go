package middlewares

import "fmt"

func notify(a, b string) {
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	subscription2 := fmt.Sprintf("%s", "bbb")
	subscription1 := fmt.Sprintf("%s", "aaa")

	notify(subscription1, subscription2)
}
