package main

type SmallStruct struct {
	ID          int
	Description string
}

func PassStructByValueReturnByValue(byValue SmallStruct) SmallStruct {
	byValue.Description = "the func updated this"
	return byValue
}

func PassStructByPointerReturnByPointer(byValue *SmallStruct) *SmallStruct {
	byValue.Description = "the func updated this"
	return byValue
}
