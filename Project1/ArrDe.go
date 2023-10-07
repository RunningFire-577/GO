package main

func Del[T any](array []T, index int) []T {
	//a1 = append(a1[:1], a1[1+1:]...)
	return append(array[:index], array[index+1:]...)
}
