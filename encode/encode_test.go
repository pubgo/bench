package encode

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	jsonitor "github.com/json-iterator/go"
	"github.com/pubgo/xerror"
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

func init() {
	std = jsonitor.ConfigCompatibleWithStandardLibrary
	cdc.RegisterConcrete(&Student{}, "main/Student", nil)
}

func TestAmino(t *testing.T) {
	s1 := Student{"张三", 18, "江苏省"}
	var dt = xerror.PanicBytes(cdc.MarshalJSON(s1))
	fmt.Printf("%s\n", dt)

	var val Student
	xerror.Panic(cdc.UnmarshalJSON(dt, &val))
	fmt.Printf("%#v\n", val)
}

func TestMsgpack(t *testing.T) {
	s1 := Student{"张三", 18, "江苏省"}
	var dt, err = msgpack.Marshal(s1)
	xerror.Panic(err)
	fmt.Printf("%s\n", dt)

	var val interface{}
	xerror.Panic(msgpack.Unmarshal(dt, &val))
	fmt.Printf("%#v\n", val)
}

func BenchmarkAminoEncode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	for i := 0; i < b.N; i++ {
		var _, _ = cdc.MarshalBinaryBare(s1)
	}
}

func BenchmarkAminoDecode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	var dt, err = cdc.MarshalBinaryBare(s1)
	xerror.Panic(err)

	var val Student
	for i := 0; i < b.N; i++ {
		_ = cdc.UnmarshalBinaryBare(dt, &val)
	}
}

func BenchmarkMsgEncode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	for i := 0; i < b.N; i++ {
		var _, _ = msgpack.Marshal(s1)
	}
}

func BenchmarkMsgDecode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	var dt, err = msgpack.Marshal(s1)
	xerror.Panic(err)

	var val Student
	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(dt, &val)
	}
}

func TestGobEncode(t *testing.T) {
	s1 := Student{"张三", 18, "江苏省"}
	_, err := serialize(s1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGobDecode(t *testing.T) {
	s1 := Student{"张三", 18, "江苏省"}
	dt, _ := serialize(s1)
	fmt.Printf("%s\n", dt)

	var val, err = deserialize(dt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", val)

	_, ok := val.(Student)
	if !ok {
		t.Fatal("not ok")
	}
}

func BenchmarkGobEncode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	for i := 0; i < b.N; i++ {
		_, _ = serialize(s1)
	}
}

func BenchmarkGobDecode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	var dt, _ = serialize(s1)
	for i := 0; i < b.N; i++ {
		_, _ = deserialize(dt)
	}
}

func BenchmarkJsonEncode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(s1)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	var dd interface{}
	var dt, _ = json.Marshal(s1)
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(dt, &dd)
	}
}

func BenchmarkJsonxEncode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	for i := 0; i < b.N; i++ {
		_, _ = std.Marshal(s1)

	}
}

func BenchmarkJsonxDecode(b *testing.B) {
	s1 := Student{"张三", 18, "江苏省"}
	var dd interface{}
	var dt, _ = std.Marshal(s1)
	for i := 0; i < b.N; i++ {
		_ = std.Unmarshal(dt, &dd)
	}
}

type Student struct {
	Name    string
	Age     uint8
	Address string
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
