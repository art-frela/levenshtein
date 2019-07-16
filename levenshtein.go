// Package levenshtein - simple lib for calculate levenshtein number for two words
// Source - https://github.com/ataklychev/levenshtein
package levenshtein

import "strings"

// Distance - Вычисление расстояния Дамерау-Левенштейна
func Distance(inputStr1, inputStr2 string) int {
	// преведение строк к нижнему регистру
	inputStr1 = strings.ToLower(inputStr1)
	inputStr2 = strings.ToLower(inputStr2)

	// преобразуем слайс байт в слайс рун
	str1 := []rune(inputStr1)
	str2 := []rune(inputStr2)

	// определяем количество символов(рун)
	str1Len := len(str1)
	str2Len := len(str2)

	// если хотя бы одна строка пустая, возвращается длина другой строки
	if str1Len == 0 && str2Len == 0 {
		return 0
	}

	if str1Len == 0 {
		return str2Len
	}

	if str2Len == 0 {
		return str1Len
	}

	// объявление матрицы
	matrix := make([][]int, str1Len+1)
	for i := range matrix {
		matrix[i] = make([]int, str2Len+1)
	}

	// инициализация нулевой строки и нулевого столбца матрицы
	for i := 0; i <= str1Len; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= str2Len; j++ {
		matrix[0][j] = j
	}

	// вычисление расстояния Дамерау-Левенштейна
	for i := 1; i <= str1Len; i++ {
		for j := 1; j <= str2Len; j++ {

			// эквивалентность символов
			symbEqual := map[bool]int{true: 0, false: 1}[(str1[i-1] == str2[j-1])]

			// определяем стоимость каждой операции
			ins := matrix[i][j-1] + 1             // добавление
			del := matrix[i-1][j] + 1             // удаление
			subst := matrix[i-1][j-1] + symbEqual // замена

			// элемент матрицы вычисляется как минимальный из трех случаев
			matrix[i][j] = min(min(ins, del), subst)

			// дополнение Дамерау по перестановке соседних символов
			if (i > 1) && (j > 1) && (str1[i-1] == str2[j-2]) && (str1[i-2] == str2[j-1]) {
				matrix[i][j] = min(matrix[i][j], matrix[i-2][j-2]+symbEqual)
			}
		}
	}
	return matrix[str1Len][str2Len]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
