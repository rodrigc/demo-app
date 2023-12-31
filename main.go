package main

import (
	"fmt"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("rodrigc is Awesome!!!", "larry3d", "yellow", true)
	myFigure.Print()

	if secret := os.Getenv("SECRET"); secret != "" {
		mySecretFigure := figure.NewColorFigure(fmt.Sprintf("Secret value is: %s", secret), "larry3d", "red", true)
		mySecretFigure.Print()
	}
	fmt.Printf("This is a BUG")
	fmt.Printf("This is a BUG 2")
	fmt.Printf("This is a BUG 3")

	time.Sleep(10 * time.Hour)
}
