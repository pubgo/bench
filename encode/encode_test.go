package encode

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	_ "github.com/alecthomas/go_serialization_benchmarks"
	"github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/rlp"
	jsonitor "github.com/json-iterator/go"
	"github.com/pubgo/xerror"
	_ "github.com/smallnest/gosercomp/model"
	amino "github.com/tendermint/go-amino"
	msgpack "github.com/vmihailenco/msgpack/v5"
	"testing"
)

var std = jsonitor.Config{
	//EscapeHTML:             true,
	//UseNumber:              true,
	//ValidateJsonRawMessage: true,
}.Froze()
var cdc = amino.NewCodec()

type d interface {
}

func init() {
	std = jsonitor.ConfigCompatibleWithStandardLibrary
	std = jsonitor.ConfigDefault
	cdc.RegisterInterface((*d)(nil), nil)
	cdc.RegisterConcrete(&Student{}, "main/Student", nil)
	cdc.RegisterConcrete(&Student1{}, "main/Student1", nil)
}

//BenchmarkUnmarshalByColfer
//BenchmarkUnmarshalByZebrapack
//BenchmarkUnmarshalByGencode
//BenchmarkUnmarshalByMsgp

//func TestAmino(t *testing.T) {
//	s1 := getStu()
//	var dt = xerror.PanicBytes(cdc.MarshalJSON(s1))
//	fmt.Printf("%s\n", dt)
//
//	var val Student
//	xerror.Panic(cdc.UnmarshalJSON(dt, &val))
//	fmt.Printf("%#v\n", val)
//}

func TestRlp(t *testing.T) {
	s1 := getStu()
	var dt, err = rlp.EncodeToBytes(s1)
	xerror.Panic(err)

	fmt.Printf("%s\n", dt)

	var val Student
	xerror.Panic(rlp.DecodeBytes(dt, &val))
	fmt.Printf("%#v\n", val)
}

func TestMsgpack(t *testing.T) {
	s1 := getStu()
	var dt, err = msgpack.Marshal(s1)
	xerror.Panic(err)
	fmt.Printf("%s\n", dt)

	var val interface{}
	xerror.Panic(msgpack.Unmarshal(dt, &val))
	fmt.Printf("%#v\n", val)
}

func BenchmarkAminoEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		var _, _ = cdc.MarshalBinaryBare(s1)
	}
}

func BenchmarkAminoDecode(b *testing.B) {
	s1 := getStu()
	var dt, err = cdc.MarshalBinaryBare(s1)
	xerror.Panic(err)

	b.ResetTimer()
	var val Student
	for i := 0; i < b.N; i++ {
		_ = cdc.UnmarshalBinaryBare(dt, &val)
	}
}

func BenchmarkRlpEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		_, _ = rlp.EncodeToBytes(s1)
	}
}

func BenchmarkRlpDecode(b *testing.B) {
	s1 := getStu()
	var dt, err = rlp.EncodeToBytes(s1)
	xerror.Panic(err)

	//fmt.Printf("%s\n", dt)
	b.ResetTimer()

	var val Student
	for i := 0; i < b.N; i++ {
		_ = rlp.DecodeBytes(dt, &val)
	}
	//fmt.Printf("%#v\n", val)
}

func BenchmarkMsgPackEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		var _, _ = msgpack.Marshal(s1)
	}
}

func BenchmarkMsgPackDecode(b *testing.B) {
	s1 := getStu()
	var dt, err = msgpack.Marshal(s1)
	xerror.Panic(err)

	b.ResetTimer()
	var val Student
	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(dt, &val)
	}
}

func TestGobEncode(t *testing.T) {
	s1 := getStu()
	_, err := serialize(s1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGobDecode(t *testing.T) {
	s1 := getStu()
	dt, _ := serialize(s1)
	fmt.Printf("%s\n", dt)

	var val, err = deserialize(dt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", val)

	_, ok := val.(*Student)
	if !ok {
		t.Fatal("not ok")
	}
}

func BenchmarkGobEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		_, _ = serialize(s1)
	}
}

func BenchmarkGobDecode(b *testing.B) {
	s1 := getStu()
	var dt, _ = serialize(s1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = deserialize(dt)
	}
}

func BenchmarkJsonEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(s1)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	s1 := getStu()
	var dd interface{}
	var dt, _ = json.Marshal(s1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(dt, &dd)
	}
}

func BenchmarkJsonitorEncode(b *testing.B) {
	s1 := getStu()
	for i := 0; i < b.N; i++ {
		_, _ = std.Marshal(s1)

	}
}

func BenchmarkJsonitorDecode(b *testing.B) {
	s1 := getStu()
	var dd interface{}
	var dt, _ = std.Marshal(s1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = std.Unmarshal(dt, &dd)
	}
}

type Student struct {
	Name    string
	Age     uint8
	Address []string
	Data    string
	Stu     *Student1
}

type Student1 struct {
	Name    string
	Age     uint8
	Address []string
	//Data    map[string]interface{}
}

func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func getStu() *Student {
	return &Student{
		Name:    "张三",
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
	}
}
