package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"
)

var display = st7789.New(machine.SPI0,
	machine.GP0,  // ResetPin
	machine.GP16, // LCD_DC
	machine.GP17, // LCD_CS
	machine.GP20, // Backlight (I think)
)

var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 0}
var green = color.RGBA{0, 255, 0, 255}
var yellow = color.RGBA{255, 255, 0, 255}
var darkred = color.RGBA{128, 0, 0, 255}
var darkblue = color.RGBA{0, 0, 128, 0}
var darkgreen = color.RGBA{0, 128, 0, 255}
var darkyellow = color.RGBA{128, 128, 0, 255}
var white = color.RGBA{255, 255, 255, 255}
var black = color.RGBA{0, 0, 0, 255}

var queue []int
var level = 0
var alive = true

func main() {
	//st7789 driver https://pkg.go.dev/tinygo.org/x/drivers/st7789
	//pin layout https://cdn.shopify.com/s/files/1/0174/1800/files/pico_display_pack_schematic.pdf?v=1639565857

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		SCK:       machine.GP18,
		SDO:       machine.GP19,
		SDI:       machine.GP19,
		Mode:      0,
	})

	display.Configure(st7789.Config{
		RowOffset:    40,
		ColumnOffset: 53,
		Width:        135,
		Height:       240,
	})

	//width, height := display.Size()

	display.EnableBacklight(true)
	display.FillScreen(black)

	buttona := machine.GP12
	buttona.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})
	buttonb := machine.GP13
	buttonb.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})
	buttonx := machine.GP14
	buttonx.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})
	buttony := machine.GP15
	buttony.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})

	buttonb.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				//Release B
			} else {
				//Press B
			}
		})
	buttona.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				//Release A
			} else {
				//Press A
			}
		})
	buttony.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				//Release Y
			} else {
				//Press Y
			}
		})
	buttonx.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				//Release X
			} else {
				//Press X
			}
		})

}
