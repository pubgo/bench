package encode

// rlp目的是可以将常用的数据结构,uint,string,[]byte,struct,slice,array,big.int等序列化以及反序列化

func GetStu() *Student {
	return &Student{
		//BirthDay: time.Date(2013, 1, 2, 3, 4, 5, 6, time.UTC),
		Phone:    "5551234567",
		Siblings: 2,
		//Spouse:   false,
		//Money:    100.0,
		//Tags:     map[string]string{"key": "value"},
		Aliases: []string{"Bobby", "Robert"},
		Name:    "张三",
		//Payload:  []byte("hi"),
		Ssid:    []uint32{1, 2, 3},
		Age:     100,
		Address: []string{"张三", "张三", "张三", "张三", "张三"},
		//Data: map[string]string{
		//	"张三":  "1",
		//	"张三2": "nil",
		//	"张三3": "sss",
		//},
		Data: "hello",
		Stu: &Student1{
			Name:    "张三",
			Age:     100,
			Address: []string{"张三", "张三", "张三", "张三", "张三"},
			//Data: map[string]interface{}{
			//	"张三":  1,
			//	"张三2": nil,
			//	"张三3": "sss",
			//},
		},
		//Stu: Student1{
		//	Name:    "张三",
		//	Age:     100,
		//	Address: []string{"张三", "张三", "张三", "张三", "张三"},
		//	//Data: map[string]interface{}{
		//	//	"张三":  1,
		//	//	"张三2": nil,
		//	//	"张三3": "sss",
		//	//},
		//},
	}
}

//go:generate msgp

type Student struct {
	//BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
	Tags     map[string]string
	Payload  []byte
	Stu      *Student1
	//Spouse   bool
	//Money    float64
	//Tags     map[string]string
	Aliases []string
	Name    string
	Age     uint8
	Address []string
	Data    string
	//Payload  []byte
	Ssid []uint32
	//Stu      Student1
}

type Student1 struct {
	Name    string
	Age     uint8
	Address []string
}
