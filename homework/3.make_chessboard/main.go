package main

import (
    "fmt"
    "strings"
)

func makeChessboard(rows, cols int, darkChar, lightChar rune) string {
    if rows <= 0 || cols <= 0 {
        return ""
    }

    var sb strings.Builder

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if (r+c) % 2 == 0 {
                sb.WriteRune(darkChar)
            } else {
                sb.WriteRune(lightChar)
            }
        }
        if r != rows - 1 {
            sb.WriteByte('\n')
        }
    }

    return sb.String()
}

func main() {
    board := makeChessboard(8, 8, '#', ' ')
    fmt.Println(board)
    fmt.Println()

    board := fmt.Println(makeChessboard(5, 12, 'X', '.'))
    fmt.Println(board)
    fmt.Println()

    board := fmt.Println(makeChessboard(15, 15, '■', '□'))
    fmt.Println(board)
    fmt.Println()

    board := fmt.Println(makeChessboard(8, 8, '■', '□'))
    fmt.Println(board)
    fmt.Println()
}