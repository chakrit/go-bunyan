package main

import "fmt"
import "github.com/kylelemons/go-gypsy/yaml"

var yamlDoc = yaml.Map{"Hello": yaml.Scalar("World")}

func main() {
	fmt.Printf("%s\n", yaml.Render(yamlDoc))
}
