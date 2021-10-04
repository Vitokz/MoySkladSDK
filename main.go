package main

import (
	"fmt"
	"github.com/Vitokz/MoySkladSDK/client"
)

func main() {
	nw := client.MoySkladClient("admin@vitkoz", "6c5f0b6c2a")
	//prd := new(entity.Product)
	//prd.Name = "sds"
	//for i := 0; i < 10; i++ {
	//	nw.Product().Create(prd)
	//}
	res, _ := nw.Product().TakeList(nil)
	for i := range *res {
		(*res)[i].Name = "llpp"
	}

	resp, err := nw.Product().CreateAndUpdateList(*res)
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = resp
}
