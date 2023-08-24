package internal

func MapKeysDifference[Key comparable, ValueA any, ValueB any](a map[Key]ValueA, b map[Key]ValueB) (aMinusB, bMinusA []Key) {
	// create set keys of A
	aKeysSet := make(map[Key]struct{}, len(a))
	for aKey, _ := range a {
		aKeysSet[aKey] = struct{}{}
	}
	// calculate A minus B difference
	for bKey, _ := range b {
		if _, ok := aKeysSet[bKey]; !ok { // we have founa s key that is present in B but not in A - add to B - A set
			bMinusA = append(bMinusA, bKey)
		}
		delete(aKeysSet, bKey) // delete every B element from A to get A - B set
	}
	// now we have A - B set in aKeysSet. map it to aMinusB slice
	for aKey, _ := range aKeysSet {
		aMinusB = append(aMinusB, aKey)
	}
	return aMinusB, bMinusA
}
