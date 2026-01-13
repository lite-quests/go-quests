package values

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
	return Result{}
}
