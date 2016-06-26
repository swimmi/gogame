package game 

import (
 	"fmt"
 	
	"consts"
 )

func Help(args []string) int {
	fmt.Println(consts.HELP)
	return 0
}