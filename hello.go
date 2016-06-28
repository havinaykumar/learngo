package main

import "fmt"

type LOC struct {
   id string
   buyer string
   seller string
   amount int
}


func main() {
   var myloc LOC

   locs := make(map[int]LOC)

   var loclist [50]LOC

   myloc.id = "LOC001"
   myloc.buyer = "Alex"
   myloc.seller = "Berney"
   myloc.amount = 6495



   fmt.Printf( "LOC id : %s\n", myloc.id)
   fmt.Printf( "LOC buyer : %s\n", myloc.buyer)
   fmt.Printf( "LOC seller : %s\n", myloc.seller)
   fmt.Printf( "LOC amount : %d\n", myloc.amount)

   loclist[0] = myloc;
   fmt.Println("LocList Array:", loclist[0])

   locs[1] = myloc;
   


   fmt.Println("LocList MAP:", locs)
}
