package lessons

import (
	"fmt"
	"math/rand"
	"time"
)

func bankHeist() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time")
	}

	openedVault := rand.Intn(100)
	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault can't be opened")
	}

	fmt.Println("isHeistOn is currently:", isHeistOn)
}
