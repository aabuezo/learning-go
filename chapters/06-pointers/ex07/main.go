package main

import "fmt"

type Storage struct {
	value *int
}

func newValue() *int {
	value := 42
	return &value
}

func storeValue(storage *Storage, value *int) {
	storage.value = value
}

func main() {
	storage := Storage{}
	value := newValue()
	storeValue(&storage, value)

	fmt.Println(*storage.value)
	// TODO: ejecutar go build -gcflags='-m=2' . y observar los escapes.
}

// % cd chapters/06-pointers/ex07
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex07% go build -gcflags='-m=2' .
// # learning-go/chapters/06-pointers/ex07
// ./main.go:9:6: can inline newValue with cost 8 as: func() *int { value := 42; return &value }
// ./main.go:14:6: can inline storeValue with cost 4 as: func(*Storage, *int) { storage.value = value }
// ./main.go:18:6: cannot inline main: function too complex: cost 108 exceeds budget 80
// ./main.go:20:19: inlining call to newValue
// ./main.go:21:12: inlining call to storeValue
// ./main.go:23:13: inlining call to fmt.Println
// ./main.go:10:2: value escapes to heap in newValue:
// ./main.go:10:2:   flow: ~r0 ← &value:
// ./main.go:10:2:     from &value (address-of) at ./main.go:11:9
// ./main.go:10:2:     from return &value (return) at ./main.go:11:2
// ./main.go:10:2: moved to heap: value
// ./main.go:14:35: parameter value leaks to {heap} for storeValue with derefs=0:
// ./main.go:14:35:   flow: {heap} ← value:
// ./main.go:14:35:     from storage.value = value (assign) at ./main.go:15:16
// ./main.go:14:17: storage does not escape
// ./main.go:14:35: leaking param: value
// ./main.go:23:14: *storage.value escapes to heap in main:
// ./main.go:23:14:   flow: {storage for ... argument} ← &{storage for *storage.value}:
// ./main.go:23:14:     from *storage.value (spill) at ./main.go:23:14
// ./main.go:23:14:     from ... argument (slice-literal-element) at ./main.go:23:13
// ./main.go:23:14:   flow: fmt.a ← &{storage for ... argument}:
// ./main.go:23:14:     from ... argument (spill) at ./main.go:23:13
// ./main.go:23:14:     from fmt.a := ... argument (assign-pair) at ./main.go:23:13
// ./main.go:23:14:   flow: {heap} ← *fmt.a:
// ./main.go:23:14:     from fmt.Fprintln(os.Stdout, fmt.a...) (call parameter) at ./main.go:23:13
// ./main.go:20:19: value escapes to heap in main:
// ./main.go:20:19:   flow: ~r0 ← &value:
// ./main.go:20:19:     from &value (address-of) at ./main.go:20:19
// ./main.go:20:19:     from ~r0 = &value (assign-pair) at ./main.go:20:19
// ./main.go:20:19:   flow: value ← ~r0:
// ./main.go:20:19:     from value := ~r0 (assign) at ./main.go:20:8
// ./main.go:20:19:   flow: value ← value:
// ./main.go:20:19:     from storage, value := &storage, value (assign-pair) at ./main.go:21:12
// ./main.go:20:19:   flow: {heap} ← value:
// ./main.go:20:19:     from storage.value = value (assign) at ./main.go:21:12
// ./main.go:20:19: moved to heap: value
// ./main.go:23:13: ... argument does not escape
// ./main.go:23:14: *storage.value escapes to heap
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex07%
