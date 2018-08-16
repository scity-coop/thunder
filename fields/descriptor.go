package fields

import "reflect"

// Descriptor is a cache object that holds onto relevant information about our struct field and
// allows us not to worry about dealing with pointers during the coercion process.
type Descriptor struct {
	Tags TagSet
	Type reflect.Type
	Kind reflect.Kind
	Ptr  bool
}

// New creates a new FieldDescriptor from a type and tags.
func New(t reflect.Type, tags []string) *Descriptor {
	it := &Descriptor{Tags: newTagSet(tags...), Type: t}

	// If the type is a pointer, dereference the type on iType and continue
	// analysis with dereference.
	it.Ptr = t.Kind() == reflect.Ptr
	if it.Ptr {
		it.Type = it.Type.Elem()
	}
	it.Kind = it.Type.Kind()

	return it
}

// Valuer creates a sql/driver.Valuer from the type and value.
func (d Descriptor) Valuer(val reflect.Value) Valuer {
	// Ideally we would de-reference pointers here in order to simplify how we work with the value.
	// However, some interfaces (I'm looking at you, gogo/protobuf) implement their methods as
	// pointer methods.
	return Valuer{Descriptor: &d, value: val}
}
