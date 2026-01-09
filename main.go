package main

import (
	"fmt"
	"os"

	"github.com/PoukD/create-structure-controller/internal/generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  create-structure-controller <NameController> [subFolder]")
		fmt.Println("Example:")
		fmt.Println("  create-structure-controller cmsController")
		fmt.Println("  create-structure-controller cmsController internal")
		os.Exit(1)
	}

	controllerName := os.Args[1]

	subFolder := ""
	if len(os.Args) >= 3 {
		subFolder = os.Args[2]
	}

	if err := generator.CreateControllerStructure(controllerName, subFolder); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
