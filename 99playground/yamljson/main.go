package main

import (
    "fmt"

    "github.com/ghodss/yaml"
)

type Person struct {
    Name string `json:"name"` // Affects YAML field names too.
    Age  int    `json:"age"`
}

func main() {
    // Marshal a Person struct to YAML.
    p := Person{"John", 30}
    y, err := yaml.Marshal(p)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }
    fmt.Println(string(y))
    /* Output:
	age: 30
	name: John
	*/

    // Unmarshal the YAML back into a Person struct.
    var p2 Person
    err = yaml.Unmarshal(y, &p2)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }
    fmt.Println(p2)
    /* Output:
	{John 30}
	*/
}

func Json2Yaml() {
    j := []byte(`{"name": "John", "age": 30}`)
    y, err := yaml.JSONToYAML(j)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }
    fmt.Println(string(y))
    /* Output:
	name: John
	age: 30
	*/
    j2, err := yaml.YAMLToJSON(y)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }
    fmt.Println(string(j2))
    /* Output:
	{"age":30,"name":"John"}
	*/
}
