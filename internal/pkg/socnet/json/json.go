package json

import (
	"unsafe"

	json "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func init() {
	json.RegisterExtension(nilSliceExtension{})
}

// nilSliceEncoder is a wrapper for default slice encoder except it marshall nil slice as empty array([])
// instead of null.
type nilSliceEncoder struct {
	sliceType *reflect2.UnsafeSliceType
	json.ValEncoder
}

func (encoder *nilSliceEncoder) Encode(ptr unsafe.Pointer, stream *json.Stream) {
	if encoder.sliceType.UnsafeIsNil(ptr) {
		stream.WriteEmptyArray()
		return
	}

	encoder.ValEncoder.Encode(ptr, stream)
}

// nilSliceExtension is a extension used to decorate default slice encoder.
type nilSliceExtension struct {
	*json.DummyExtension
}

func (receiver nilSliceExtension) DecorateEncoder(typ reflect2.Type, encoder json.ValEncoder) json.ValEncoder {
	sliceType, ok := typ.(*reflect2.UnsafeSliceType)
	if !ok {
		return encoder
	}

	return &nilSliceEncoder{sliceType: sliceType, ValEncoder: encoder}
}
