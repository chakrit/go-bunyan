package main

import "testing"
import "github.com/kylelemons/go-gypsy/yaml"
import a "github.com/stretchr/testify/assert"

func TestYamlRender(t *testing.T) {
	const EXPECTED = "Hello: World\n"
	a.Equal(t, yaml.Render(yamlDoc), EXPECTED, "incorrect rendering!")
}

func ExampleMainRun() {
	main()
	// Output: Hello: World
}
