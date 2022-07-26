package datahelper

import (
	"fmt"

	"github.com/romeq/testaustime-cli/apiengine"
	"github.com/romeq/testaustime-cli/logger"
)

func showList(list apiengine.TopStatsList) {
	for i, item := range list {
		if i >= 5 {
			break
		}

		if item.Name == "" {
			item.Name = "<none>"
		}
		color := 37
		if i == 0 {
			color = 32
		}

		printField(item.Name, rawTimeToHumanReadable(item.Time), color)
	}
}

func rawTimeToHumanReadable(minutesCoded float32) string {
	var daysCoded, hoursCoded int
	remainderMinutes := int(minutesCoded)

	for remainderMinutes >= 60 {
		if remainderMinutes >= 60*24 {
			remainderMinutes -= 60 * 24
			daysCoded += 1
			continue
		}
		remainderMinutes -= 60
		hoursCoded += 1
	}

	if daysCoded == 0 && hoursCoded > 0 {
		return fmt.Sprintf("%dh, %dm", hoursCoded, remainderMinutes)
	} else if daysCoded == 0 && hoursCoded == 0 {
		return fmt.Sprintf("%dm", remainderMinutes)
	}
	return fmt.Sprintf("%dd, %dh, %dm", daysCoded, hoursCoded, remainderMinutes)
}

func printField(key string, value any, color int) {
	if logger.ColorsEnabled {
		fmt.Printf("\033[%dm%s\033[0m: %s\n", color, key, value)
	} else {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func printBold(a ...any) {
	printColored(1, a...)
}

func printColored(color int, a ...any) {
	if logger.ColorsEnabled {
		fmt.Printf("\033[%dm", color)
		fmt.Print(a...)
		fmt.Printf("\033[0m\n")
		return
	}
	fmt.Print(a...)
	fmt.Printf("\n")
}
