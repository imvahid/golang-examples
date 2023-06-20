## :man_technologist: Examples of my practice projects with the Go

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)

## :blue_book: List of practices
- Hello World
  ```go
  fmt.Println("Hello World!")
  ```
- Ponters
  ```go
  i := 52

  var p *int
  fmt.Println("p:", p)
  p = &i

  fmt.Println("i:", i)
  fmt.Println("p:", p)
  fmt.Println("&i:", &i)
  fmt.Println("*p:", *p)

  i = 92

  fmt.Println("*p:", *p, "p:", p)
  ```
  
- Variables
  ```go
  var integer1 int
  fmt.Println(integer1)

  var string1 string
  string1 = "This is a string variable"
  fmt.Println(string1)

  var float1 = -1.2
  fmt.Println(float1)

  integer2 := 52
  fmt.Println(integer2)

  string2 := "This is a string"
  fmt.Println(string2)

  var var1, var2, var3 = -1.2, 1, false
  fmt.Println(var1, var2, var3)

  var (
    var4 float64 = -1.2
    var5 int     = 1
    var6 bool    = true
    var7 string
  )
  fmt.Println(var4, var5, var6, var7)
  ```

- Data Structures
  ```go
  // array
  integerArray := [5]int{10, 20, 30, 40, 50}
  fmt.Println("integerArray:", integerArray)
  stringArray := [4]string{"first", "second", "third", "fourth"}
  fmt.Println("stringArray:", stringArray)

  // slice
  integerSlice = []int{10, 20, 30, 40}
  fmt.Println("integerSlice:", integerSlice)
  stringSlice = []string{"first", "second", "third"}
  fmt.Println("stringSlice:", stringSlice)

  // map
  n := map[string]int{"foo": 1, "bar": 2, "test": 3, "go": 4}
  fmt.Println("map n:", n)

  // struct
  type person struct {
    name   string
    family string
    age    int
  }
  ```
  
...
