package main

func removeEvenNumbers(list []int) (oddNumbers []int) { //parametro: um array/slice - retorno um array/slice de nome oddNumbers do tipo int
	for i := 0; i < len(list); i++ {
		if list[i]%2 != 0 {
			oddNumbers = append(oddNumbers, list[i])
		}
	}
	return
}
