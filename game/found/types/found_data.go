package types

//找回资源数据
type FoundData struct {
	FoundSilver   int32
	FoundGold     int32
	FoundBindgold int32
	FoundExp      int32
	FoundExpPoint int32
	FoundItemMap  map[int32]int32
}

func CreateFoundData() FoundData {
	found := FoundData{
		FoundItemMap: make(map[int32]int32),
	}
	return found
}
