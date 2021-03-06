package shortquery

import (
	"syscall/js"
)

// WasmQueryOuterWidthMargin is a function from shortQuery module.
// This function sets or returns the sum of width, padding, border and margin of an html element. Can work only on a single element.
// Never returns undefined.
func WasmQueryOuterWidthMargin(this js.Value, args []js.Value) interface{} {
	width := js.Global().Get("window").Call("getComputedStyle", this).Get("width").String()
	paddingLeft := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-left").String()
	paddingRight := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-right").String()
	borderLeft := js.Global().Get("window").Call("getComputedStyle", this).Get("border-left-width").String()
	borderRight := js.Global().Get("window").Call("getComputedStyle", this).Get("border-right-width").String()
	marginLeft := js.Global().Get("window").Call("getComputedStyle", this).Get("margin-left").String()
	marginRight := js.Global().Get("window").Call("getComputedStyle", this).Get("margin-right").String()
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject && value.Length() > 0 {
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeString {
					this.Get("style").Set("width", "calc("+value.Index(i).String()+" - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+" - "+marginLeft+" - "+marginRight+")")
				} else if value.Index(i).Type() == js.TypeNumber {
					this.Get("style").Set("width", "calc("+value.Index(i).String()+"px - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+" - "+marginLeft+" - "+marginRight+")")
				}
			}
			return WasmQueryOuterWidthMargin(this, []js.Value{})
		} else if value.Type() == js.TypeString {
			this.Get("style").Set("width", "calc("+value.String()+" - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+" - "+marginLeft+" - "+marginRight+")")
			return WasmQueryOuterWidthMargin(this, []js.Value{})
		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("width", "calc("+value.String()+"px - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+" - "+marginLeft+" - "+marginRight+")")
			return WasmQueryOuterWidthMargin(this, []js.Value{})
		}
	}
	fWidth := js.Global().Call("parseFloat", js.ValueOf(width[:len(width)-2])).Float()
	fPaddingLeft := js.Global().Call("parseFloat", js.ValueOf(paddingLeft[:len(paddingLeft)-2])).Float()
	fPaddingRight := js.Global().Call("parseFloat", js.ValueOf(paddingRight[:len(paddingRight)-2])).Float()
	fBorderLeft := js.Global().Call("parseFloat", js.ValueOf(borderLeft[:len(borderLeft)-2])).Float()
	fBorderRight := js.Global().Call("parseFloat", js.ValueOf(borderRight[:len(borderRight)-2])).Float()
	fMarginLeft := js.Global().Call("parseFloat", js.ValueOf(marginLeft[:len(marginLeft)-2])).Float()
	fMarginRight := js.Global().Call("parseFloat", js.ValueOf(marginRight[:len(marginRight)-2])).Float()
	return js.Global().Call("String", js.ValueOf(fWidth+fPaddingLeft+fPaddingRight+fBorderRight+fBorderLeft+fMarginRight+fMarginLeft)).String() + "px"
}
