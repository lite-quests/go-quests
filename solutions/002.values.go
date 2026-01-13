package solutions

type User struct {
	Name string
	Age  int
}

// example
// str:="Hello"
// num:=42
// return Result{Str:str, Int:num}

type Result struct {
	Str     string             // str:="hello"
	Int     int                // Figure out yourself
	Float   float64            // Figure out yourself
	Bool    bool               // flag:=true
	Array   [3]int             // arr:=[3]int{6,7,8}
	Slice   []int              // Figure out yourself
	Map     map[string]int     // mp:=map[string]int{"apple":2,"banana":5}
	User    User               // user:=User{name:"Bob", age:20}
	Ptr     *int               // p:=&num
	AddFn   func(int, int) int // add:=func(a int,b int)int{return a-b}
	Any     interface{}        // var data interface{}="sample data"
	ZeroMap map[string]int
}

func BuildValues() Result {
	// TODO: implement
	// Read README.md for the instructions
	var strVal string = "go"
	var intVal int = 42
	var floatVal float64 = 3.14
	var boolVal bool = true
	var arrayVal [3]int
	arrayVal = [3]int{1, 2, 3}
	var sliceVal []int
	sliceVal = []int{4, 5, 6, 7}
	var mapVal map[string]int
	mapVal = map[string]int{"apple": 2, "banana": 5}
	var userVal User = User{Name: "Alice", Age: 20}
	var num int = 10
	var ptrVal *int = &num
	var addFnVal func(int, int) int = func(a int, b int) int { return a + b }
	var interfaceVal interface{} = 100
	var zeroMapVal map[string]int
	return Result{Str: strVal, Int: intVal, Float: floatVal, Bool: boolVal, Array: arrayVal, Slice: sliceVal, Map: mapVal, User: userVal, Ptr: ptrVal, AddFn: addFnVal, Any: interfaceVal, ZeroMap: zeroMapVal}
}
