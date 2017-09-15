package main

import (
    "fmt"
    "github.com/junhsieh/iojson"
    "strings"
)

func NewCar() *Car {
    return &Car{ItemArr: make([]Item, 0)}
}

type Car struct {
    Name    string
    ItemArr []Item
}

func (c *Car) GetName() string {
    return c.Name
}

func NewItem() *Item {
    return &Item{}
}

type Item struct {
    Name string
}

func (i *Item) GetName() string {
    return i.Name
}

func main()  {
    main3()
}

func main1() {
    item := NewItem()
    item.Name = "Bag"

    car := NewCar()
    car.Name = "My luxury car"
    car.ItemArr = append(car.ItemArr, *item)

    o := iojson.NewIOJSON()
    o.AddObjToArr(car) // add the car object to the slice.

    fmt.Printf("%s\n", o.EncodePretty()) // encode data with nice format or call o.Encode().
}

func main2() {
    item := NewItem()
    item.Name = "Bag"

    car := NewCar()
    car.Name = "My luxury car"
    car.ItemArr = append(car.ItemArr, *item)

    o := iojson.NewIOJSON()
    o.AddObjToMap("Car", car) // add the car object to the map.

    fmt.Printf("%s\n", o.EncodePretty()) // encode data with nice format or call o.Encode().
}

func main3() {
    jsonStr := `{"Status":true,"ErrArr":[],"ObjArr":[{"Name":"My luxury car","ItemArr":[{"Name":"Bag"},{"Name":"Pen"}]}],"ObjMap":{}}`

    car := NewCar()

    i := iojson.NewIOJSON()

    if err := i.Decode(strings.NewReader(jsonStr)); err != nil {
        fmt.Printf("err: %s\n", err.Error())
    }

    m1 := make(map[string]interface{})
    obj := "Status"
    s1, _ := i.GetObjFromMap(obj, m1)
    fmt.Println(s1)


    // populating data to a live car object.
    if v, err := i.GetObjFromArr(0, car); err != nil {
        fmt.Printf("err: %s\n", err.Error())
    } else {
        fmt.Printf("car (original): %s\n", car.GetName())
        fmt.Printf("car (returned): %s\n", v.(*Car).GetName())

        for k, item := range car.ItemArr {
            fmt.Printf("ItemArr[%d] of car (original): %s\n", k, item.GetName())
        }

        for k, item := range v.(*Car).ItemArr {
            fmt.Printf("ItemArr[%d] of car (returned): %s\n", k, item.GetName())
        }
    }
}
