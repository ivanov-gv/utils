package internal

import "fmt"

func Mapper[In, Out any](inputs []In, mapper func(In) Out) []Out {
	result := make([]Out, 0, len(inputs))
	for _, in := range inputs {
		result = append(result, mapper(in))
	}
	return result
}

func SliceToMap[KeyType comparable, ValueType any](src []ValueType,
	generateKey func(index int, value ValueType) KeyType) map[KeyType]ValueType {
	var result = make(map[KeyType]ValueType, len(src))
	for i, value := range src {
		result[generateKey(i, value)] = value
	}
	return result
}

func SliceToMapWithUniquenessError[Key comparable, Value any](src []Value, getKey func(Value) Key) (map[Key]Value, error) {
	var result = make(map[Key]Value, len(src))
	for _, v := range src {
		key := getKey(v)
		if _, ok := result[key]; ok { // check uniqueness
			return nil, fmt.Errorf("key %v is not unique", key)
		}
		result[key] = v
	}
	return result, nil
}

func MapToSlice[KeyType comparable, ValueType any, Out any](mapper func(key KeyType, value ValueType) Out,
	_map map[KeyType]ValueType,
) []Out {
	result := make([]Out, 0, len(_map))
	for key, value := range _map {
		result = append(result, mapper(key, value))
	}
	return result
}
