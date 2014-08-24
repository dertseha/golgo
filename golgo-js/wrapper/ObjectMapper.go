package wrapper

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type ObjectMapper interface {
	put(value js.Object) uint
	get(key uint) js.Object
	del(key uint) js.Object
}

type objectMap struct {
	objects map[uint]js.Object
	counter uint
}

func NewObjectMapper() ObjectMapper {
	result := &objectMap{
		objects: make(map[uint]js.Object),
		counter: 0}

	return result
}

func (omap *objectMap) put(value js.Object) uint {
	var key uint = 0

	for key == 0 {
		_, exists := omap.objects[omap.counter]

		if (omap.counter == 0) || exists {
			omap.counter++
		} else {
			key = omap.counter
		}
	}
	omap.objects[key] = value

	return key
}

func (omap *objectMap) get(key uint) js.Object {
	value, ok := omap.objects[key]

	if !ok {
		panic(fmt.Sprintf("Object with ID %u not known", key))
	}

	return value
}

func (omap *objectMap) del(key uint) js.Object {
	defer delete(omap.objects, key)

	return omap.get(key)
}
