// Copyright (C) Egor
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
// GNU Lesser General Public Licence is available at
// http://www.gnu.org/copyleft/lesser.html
package gtuiqm

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func Clear() {
	fmt.Printf("\033c")
}
func ShowMenu(menu []string, title string, fullline bool, selected_bg_color int, selected_color int) {
	elem := 0
	CreateMenu(elem, menu, title, fullline, selected_bg_color, selected_color)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if elem != 0 {
				elem--
			} else {
				elem = len(menu) - 1
			}
			CreateMenu(elem, menu, title, fullline, selected_bg_color, selected_color)

		case keys.Down:

			if elem != len(menu)-1 {
				elem++
			} else {
				elem = 0
			}

			CreateMenu(elem, menu, title, fullline, selected_bg_color, selected_color)

		case keys.Enter:
			return true, nil
		}

		return false, nil // Return false to continue listening
	})

}
func CreateMenu(elem int, menu []string, title string, fullline bool, selected_bg_color int, selected_color int) {

	Clear()
	fmt.Printf("\x1b[35m" + title + "\x1b[0m\n")
	for i := 0; i < len(menu); i++ {
		if elem == i {
			if fullline {
				fmt.Printf("\x1b[4" + fmt.Sprint(selected_bg_color) + ";3" + fmt.Sprint(selected_color) + ";4" + fmt.Sprint(selected_bg_color) + "m" + menu[i] + "\x1b[K\x1b[0m\n")
			} else {
				fmt.Printf("\x1b[4" + fmt.Sprint(selected_bg_color) + ";3" + fmt.Sprint(selected_color) + ";4" + fmt.Sprint(selected_bg_color) + "m" + menu[i] + "\x1b[0m\n")
			}
		} else {
			fmt.Println(menu[i])
		}
	}
}
