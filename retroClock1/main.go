package main

import (
	"fmt"
	"time"
)

func main() {
	type location [5]string

	zero := location{
		"###",
		"# #",
		"# #",
		"# #",
		"###",
	}
	one := location{
		"## ",
		" # ",
		" # ",
		" # ",
		"###",
	}
	two := location{
		"###",
		"  #",
		"###",
		"#  ",
		"###",
	}
	three := location{
		"###",
		"  #",
		"###",
		"  #",
		"###",
	}
	four := location{
		"# #",
		"# #",
		"###",
		"  #",
		"  #",
	}
	five := location{
		"###",
		"#  ",
		"###",
		"  #",
		"###",
	}
	six := location{
		"###",
		"#  ",
		"###",
		"# #",
		"###",
	}
	seven := location{
		"###",
		"  #",
		"  #",
		"  #",
		"  #",
	}
	eight := location{
		"###",
		"# #",
		"###",
		"# #",
		"###",
	}
	nine := location{
		"###",
		"# #",
		"###",
		"  #",
		"###",
	}

	digits := [...]location{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	seperator := location{
		"   ",
		" # ",
		"   ",
		" # ",
		"   ",
	}

	fmt.Print("\033[2J") // Clearing the console

	for {

		fmt.Print("\033[H") // Moving the cursor to the top-left corner

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]location{
			digits[hour/10], digits[hour%10],
			seperator,
			digits[min/10], digits[min%10],
			seperator,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
