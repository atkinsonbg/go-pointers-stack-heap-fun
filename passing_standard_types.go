package main

import (
	"fmt"
	"strings"
)

func PassByValueReturnValueConcatenation(byValueConcat string) string {
	return byValueConcat + "-updated"
}

func PassByValueReturnValueFmt(byValueFmt string) string {
	return fmt.Sprintf("%v-updated", byValueFmt)
}

func PassByValueReturnValueBuilder(byValueBuilder string) string {
	var sb strings.Builder
	sb.WriteString(byValueBuilder)
	sb.WriteString("-updated")
	return sb.String()
}

func PassByPointerReturnValueConcatenation(byPointerConcat *string) string {
	return *byPointerConcat + "-updated"
}

func PassByPointerReturnPointerConcatenation(byPointerConcat *string) *string {
	s := *byPointerConcat + "-updated"
	return &s
}

////go:noinline
//func PassByPointerReturnValue(originalByPointer *string) string {
//	return fmt.Sprintf("%v-updated", originalByPointer)
//}
//
////go:noinline
//func PassByPointerReturnPointer(originalByPointer *string) *string {
//	val := fmt.Sprintf("%v-updated", originalByPointer)
//	return &val
//}
