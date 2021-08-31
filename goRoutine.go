package main

import (
	"fmt"
	"time"
)

func wrap(){
	
	start := time.Now()
	ninja := []string{"tommy", "bobby", "kenneth", "nnamdi"}
	for _, evil  := range ninja{
		go attack(evil)
	}
	time.Sleep(time.Second*2)
	defer func(){
		fmt.Println(time.Since(start))
	}()
	
}
// Attack method called inside rap
func attack(target string){
 fmt.Printf("throwing ninja strikes on %v\n", target)
 time.Sleep(time.Second)
}

func main(){
   wrap()
}