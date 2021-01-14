/*
   Each day a plant is growing by upSpeed meters. Each night that plant's height decreases by downSpeed meters due to the lack of sun heat. Initially, plant is 0 meters tall. We plant the seed at the beginning of a day. We want to know when the height of the plant will reach a certain level.
*/

package main

import (
	"fmt"
)

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	currentHeight := upSpeed
	days := 1
	
	for {
		if days > 1 {
			currentHeight += upSpeed
		}
		
		if currentHeight >= desiredHeight {
			break 
		}
		
		currentHeight -= downSpeed
		days += 1
	}
	
	return days
}

func main() {
	fmt.Println(GrowingPlant(100, 10, 910)) // should be 10
	fmt.Println(GrowingPlant(10, 9, 4)) // should be 1
}
