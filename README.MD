# Go Text User Interface Question Manager | GTUIQM

## Contents<a name="contents"></a> 

* [Contents](#contents)
* [Docs](#docs)


## Docs <a name="docs"></a> 

```go
Clear() //clear screen
ShowMenu(menu []string, title string, fullline bool, selected_bg_color int, selected_color int) //showmenu, returns the number of the element in the array and handles keyboard presses
CreateMenu(elem int, menu []string, title string, fullline, selected_bg_color int, selected_color int)//base of the ShowMenu

```
