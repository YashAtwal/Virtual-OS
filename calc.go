package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	output := ""
	input := widget.NewLabel(output)
	isHistory := false
	historyStr := ""
	history := widget.NewLabel(historyStr)
	var historyArr []string

	historyB := widget.NewButton("History", func() {
		if isHistory {
			historyStr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr += historyArr[i]
				historyStr += "\n"
			}
		}
		isHistory = !isHistory
		history.SetText(historyStr)
	})
	deleteB := widget.NewButton("Delete", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
		}
	})
	clearB := widget.NewButton("Clear", func() {
		output = " "
		input.SetText(output)
	})
	closebracB := widget.NewButton(")", func() {
		output += ")"
		input.SetText(output)
	})
	openbracB := widget.NewButton("(", func() {
		output += "("
		input.SetText(output)
	})
	divideB := widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	sevenB := widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	eightB := widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	nineB := widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	multiplyB := widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	fourB := widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	fiveB := widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	sixB := widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	minusB := widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	oneB := widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	twoB := widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	threeB := widget.NewButton("3", func() {
		output += "3"
		input.SetText(output)
	})
	plusB := widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	zeroB := widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	dotB := widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	equalB := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "Error"
			}
		} else {
			output = "Error"
		}
		input.SetText(output)
	})

	calcContainer := container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(2,
			historyB,
			deleteB,
		),
		container.NewGridWithColumns(4,
			clearB,
			openbracB,
			closebracB,
			divideB,
		),
		container.NewGridWithColumns(4,
			sevenB,
			eightB,
			nineB,
			multiplyB,
		),
		container.NewGridWithColumns(4,
			fourB,
			fiveB,
			sixB,
			minusB,
		),
		container.NewGridWithColumns(4,
			oneB,
			twoB,
			threeB,
			plusB,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				zeroB,
				dotB,
			),
			equalB,
		),
	)
	w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(150, 200))
	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, calcContainer))

	w.Show()
}
