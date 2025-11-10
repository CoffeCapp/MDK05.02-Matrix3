package main

import (
	"fmt"
)

func showMatr(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func matrMul(matrix [][]int, num int) [][]int {
	result := make([][]int, len(matrix))
	for i := range result {
		result[i] = make([]int, len(matrix[i]))
		for j := range matrix[i] {
			result[i][j] = matrix[i][j] * num
		}
	}
	return result
}

func matrMul2(matrix [][]int, num int) [][]int {
	return matrMul(matrix, num)
}

func multiplyMatrices(a, b [][]int) [][]int {
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
		for j := range result[i] {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func getMatr(size int) [][]int {
	matrix := make([][]int, size)
	fmt.Printf("Введите элементы матрицы %dx%d построчно, через пробел:\n", size, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			for {
				fmt.Printf("Введите элемент [%d][%d]: ", i+1, j+1)
				_, err := fmt.Scan(&matrix[i][j])
				if err != nil {
					fmt.Println("Ошибка: введите целое число.")
					var discard string
					fmt.Scan(&discard)
					continue
				}
				break
			}
		}
	}
	return matrix
}

func main() {
	var matrix1, matrix2 [][]int
	var size int
	var exists1, exists2 bool

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Ввести первую матрицу")
		fmt.Println("2. Ввести вторую матрицу")
		fmt.Println("3. Вывести матрицы")
		fmt.Println("4. Умножить первую матрицу на число")
		fmt.Println("5. Умножить вторую матрицу на число")
		fmt.Println("6. Умножить первую матрицу на вторую")
		fmt.Println("7. Выйти")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода. Введите число.")
			var discard string
			fmt.Scan(&discard)
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Выберите размер: 2, 3 или 4:")
			fmt.Scan(&size)
			if size < 2 || size > 4 {
				fmt.Println("Недопустимый размер.")
				continue
			}
			matrix1 = getMatr(size)
			exists1 = true
			fmt.Println("Первая матрица создана:")
			showMatr(matrix1)

		case 2:
			if !exists1 {
				fmt.Println("Сначала введите первую матрицу.")
				continue
			}
			matrix2 = getMatr(size)
			exists2 = true
			fmt.Println("Вторая матрица создана:")
			showMatr(matrix2)

		case 3:
			if exists1 {
				fmt.Println("Матрица 1:")
				showMatr(matrix1)
			} else {
				fmt.Println("Первая матрица не создана.")
			}
			if exists2 {
				fmt.Println("Матрица 2:")
				showMatr(matrix2)
			} else {
				fmt.Println("Вторая матрица не создана.")
			}

		case 4:
			if !exists1 {
				fmt.Println("Сначала введите первую матрицу.")
				continue
			}
			var num int
			fmt.Print("Введите число: ")
			_, err := fmt.Scan(&num)
			if err != nil {
				fmt.Println("Ошибка ввода. Введите целое число.")
				var discard string
				fmt.Scan(&discard)
				continue
			}
			result := matrMul(matrix1, num)
			fmt.Println("Результат умножения первой матрицы на", num, ":")
			showMatr(result)

		case 5:
			if !exists2 {
				fmt.Println("Сначала введите вторую матрицу.")
				continue
			}
			var num int
			fmt.Print("Введите число: ")
			_, err := fmt.Scan(&num)
			if err != nil {
				fmt.Println("Ошибка ввода. Введите целое число.")
				var discard string
				fmt.Scan(&discard)
				continue
			}
			result := matrMul2(matrix2, num)
			fmt.Println("Результат умножения второй матрицы на", num, ":")
			showMatr(result)

		case 6:
			if !exists1 || !exists2 {
				fmt.Println("Создайте обе матрицы.")
				continue
			}
			result := multiplyMatrices(matrix1, matrix2)
			fmt.Println("Результат умножения матриц:")
			showMatr(result)

		case 7:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
