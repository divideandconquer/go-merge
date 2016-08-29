package merge

import (
	"reflect"
	"testing"

	"github.com/divideandconquer/go-merge/merge"
)

func makeStringPtr(v string) *string {
	return &v
}

func TestUnit_Merge_BasePath(t *testing.T) {
	a := ""
	b := "foo"
	expected := "foo"
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StringNotSetB(t *testing.T) {
	a := "bar"
	var b string
	expected := "bar"
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StringNotSetA(t *testing.T) {
	var a string
	b := "bar"
	expected := "bar"
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StringEmptyB(t *testing.T) {
	a := "bar"
	b := ""
	expected := "bar"
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StringBothSet(t *testing.T) {
	a := "bar"
	b := "foo"
	expected := "foo"
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_IntEmptyA(t *testing.T) {
	a := 0
	b := 10
	expected := 10
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_IntEmptyB(t *testing.T) {
	a := 10
	b := 0
	expected := 10
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_IntNotSetA(t *testing.T) {
	var a int
	b := 10
	expected := 10
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_IntNotSetB(t *testing.T) {
	a := 10
	var b int
	expected := 10
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_IntBothSet(t *testing.T) {
	a := 5
	b := 10
	expected := 10
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_UIntEmptyA(t *testing.T) {
	a := uint64(0)
	b := uint64(10)
	expected := uint64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_UIntEmptyB(t *testing.T) {
	a := uint64(10)
	b := uint64(0)
	expected := uint64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_UIntNotSetA(t *testing.T) {
	var a uint64
	b := uint64(10)
	expected := uint64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_UIntNotSetB(t *testing.T) {
	a := uint64(10)
	var b uint64
	expected := uint64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_UIntBothSet(t *testing.T) {
	a := uint64(5)
	b := uint64(10)
	expected := uint64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_FloatEmptyA(t *testing.T) {
	a := float64(0)
	b := float64(10)
	expected := float64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_FloatEmptyB(t *testing.T) {
	a := float64(10)
	b := float64(0)
	expected := float64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_FloatNotSetA(t *testing.T) {
	var a float64
	b := float64(10)
	expected := float64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_FloatNotSetB(t *testing.T) {
	a := float64(10)
	var b float64
	expected := float64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_FloatBothSet(t *testing.T) {
	a := float64(5)
	b := float64(10)
	expected := float64(10)
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_BoolNotSetA(t *testing.T) {
	var a bool
	b := true
	expected := true
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_BoolNotSetB(t *testing.T) {
	a := true
	var b bool
	expected := true
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

// Note that because an empty bool is the same as false, true always wins out
func TestUnit_Merge_AlternatePath_BoolBothSet(t *testing.T) {
	a := true
	b := false
	expected := true
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_ArrayEmptyA(t *testing.T) {
	a := []int{}
	b := []int{10, 11, 12}
	expected := []int{10, 11, 12}
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_ArrayEmptyB(t *testing.T) {
	a := []int{10, 11, 12}
	b := []int{}
	expected := []int{10, 11, 12}
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_ArrayNotSetA(t *testing.T) {
	var a []int
	b := []int{10, 11, 12}
	expected := []int{10, 11, 12}
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_ArrayNotSetB(t *testing.T) {
	a := []int{10, 11, 12}
	var b []int
	expected := []int{10, 11, 12}
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_ArrayBothSet(t *testing.T) {
	a := []int{5, 6, 7}
	b := []int{10, 11, 12}
	expected := []int{10, 11, 12}
	ret := Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_Pointers(t *testing.T) {
	a := 10
	b := 5
	expected := &b
	ret := Merge(&a, &b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_Map(t *testing.T) {
	baseData := make(map[string]string)
	baseData["test"] = "foo"
	baseData["test2"] = "bar"

	overrideData := make(map[string]string)
	overrideData["test2"] = "override"

	expected := make(map[string]string)
	expected["test"] = "foo"
	expected["test2"] = "override"

	ret := merge.Merge(baseData, overrideData)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_MapEmptyValue(t *testing.T) {
	baseData := make(map[string]string)
	baseData["test"] = "foo"
	baseData["test2"] = ""

	overrideData := make(map[string]string)
	overrideData["test3"] = ""

	expected := make(map[string]string)
	expected["test"] = "foo"
	expected["test2"] = ""
	expected["test3"] = ""

	ret := merge.Merge(baseData, overrideData)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_MapEmptyValuePtr(t *testing.T) {
	baseData := make(map[string]*string)
	baseData["test"] = makeStringPtr("foo")
	baseData["test2"] = makeStringPtr("")

	overrideData := make(map[string]*string)
	overrideData["test3"] = makeStringPtr("")

	expected := make(map[string]*string)
	expected["test"] = makeStringPtr("foo")
	expected["test2"] = makeStringPtr("")
	expected["test3"] = makeStringPtr("")

	ret := merge.Merge(baseData, overrideData)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_MapInterface(t *testing.T) {
	base := make(map[string]map[string]interface{})
	base["Config"] = make(map[string]interface{})
	base["Config"]["Test"] = "this is a test value"
	base["Config"]["Test2"] = "this is also a test value"

	override := make(map[string]map[string]interface{})
	override["Config"] = make(map[string]interface{})
	override["Config"]["foo"] = "bar"
	override["Config"]["Test2"] = "baz"

	expected := make(map[string]map[string]interface{})
	expected["Config"] = make(map[string]interface{})
	expected["Config"]["foo"] = "bar"
	expected["Config"]["Test"] = "this is a test value"
	expected["Config"]["Test2"] = "baz"

	ret := merge.Merge(base, override)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructMapStruct(t *testing.T) {
	type testStruct struct {
		Foo string
		Bar string
	}

	type mStruct struct {
		M map[string]testStruct
	}

	b := make(map[string]testStruct)
	o := make(map[string]testStruct)
	e := make(map[string]testStruct)

	base := testStruct{}
	base.Foo = "foo"
	base.Bar = "hello"
	b["test"] = base
	bs := mStruct{b}

	override := testStruct{}
	override.Bar = "bar"
	o["test"] = override
	os := mStruct{o}

	expected := testStruct{}
	expected.Foo = "foo"
	expected.Bar = "bar"
	e["test"] = expected
	es := &mStruct{e}

	ret := Merge(bs, os)

	if !reflect.DeepEqual(ret, es) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, es)
	}
}

func TestUnit_Merge_AlternatePath_StructPointer(t *testing.T) {

	type innerStruct struct {
		InnerVal string
	}
	type testStruct struct {
		Foo string
		Bar *innerStruct
	}

	i1 := &innerStruct{"inner"}
	base := testStruct{"bar", i1}

	i2 := &innerStruct{"inner2"}
	override := testStruct{"", i2}

	expected := testStruct{"bar", i2}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).Bar.InnerVal != expected.Bar.InnerVal {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructPrivate(t *testing.T) {

	type innerStruct struct {
		InnerVal string
	}
	type testStruct struct {
		Foo string
		bar innerStruct
	}

	i1 := innerStruct{"inner"}
	base := testStruct{"bar", i1}

	i2 := innerStruct{"inner2"}
	override := testStruct{"", i2}

	expected := testStruct{"bar", innerStruct{}}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).bar.InnerVal != expected.bar.InnerVal {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructWithStringPointers(t *testing.T) {
	type testStruct struct {
		Foo string
		Bar *string
	}

	bs := "bar"
	base := testStruct{Foo: "foo"}

	override := testStruct{"foo2", &bs}
	expected := testStruct{"foo2", &bs}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).Bar != expected.Bar {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructWithStringNilPointers(t *testing.T) {
	type testStruct struct {
		Foo string
		Bar *string
	}

	bs := "bar"
	base := testStruct{"foo", &bs}

	override := testStruct{"foo2", nil}
	expected := testStruct{"foo2", &bs}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).Bar != expected.Bar {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructWithStructNilPointers(t *testing.T) {

	type innerStruct struct {
		InnerVal string
	}
	type testStruct struct {
		Foo string
		Bar *innerStruct
	}

	bs := innerStruct{"bar"}
	base := testStruct{"foo", &bs}

	override := testStruct{"foo2", nil}
	expected := testStruct{"foo2", &bs}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).Bar != expected.Bar {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

func TestUnit_Merge_AlternatePath_StructWithStructNilPointersBase(t *testing.T) {

	type innerStruct struct {
		InnerVal string
	}
	type testStruct struct {
		Foo string
		Bar *innerStruct
	}

	bs := innerStruct{"bar"}
	base := testStruct{"foo", nil}

	override := testStruct{"foo2", &bs}
	expected := testStruct{"foo2", &bs}

	ret := Merge(base, override)
	if ret.(*testStruct).Foo != expected.Foo || ret.(*testStruct).Bar != expected.Bar {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}
