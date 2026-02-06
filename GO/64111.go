package main

import "fmt"

func main() {
	var arsenid float64
	fmt.Scanln(&arsenid)
	var gosha float64
	fmt.Scanln(&gosha)
	var irinka float64
	fmt.Scanln(&irinka)

	if        arsenid > gosha && gosha > irinka {
		fmt.Println(irinka)
	} else if gosha > arsenid && arsenid > irinka {
		fmt.Println(irinka)
	} else if arsenid > irinka && irinka > gosha {
		fmt.Println(gosha)
	} else if gosha > irinka && irinka > arsenid {
		fmt.Println(arsenid)
	} else if irinka > arsenid && arsenid > gosha {
		fmt.Println(gosha)									
	} else if irinka > gosha && gosha > arsenid {
		fmt.Println(arsenid)					
	} else if irinka == gosha && gosha == arsenid {
			fmt.Println(arsenid)				
	} else {
		fmt.Println(0)
	}
}
