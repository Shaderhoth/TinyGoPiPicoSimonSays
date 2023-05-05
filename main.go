package main

import (
	"image/color"
	"machine"
	"math/rand"

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
	for i := 0; i < 4; i++ {
		dull(i)
	}

	buttonb.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				release(0)
			} else {
				press(0)
			}
		})
	buttona.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() { // up
				release(1)
			} else {
				press(1) // down
			}
		})
	buttony.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				release(2)
			} else {
				press(2)
			}
		})
	buttonx.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			if p.Get() {
				release(3)
			} else {
				press(3)
			}
		})
	println("Start")
	makeQueue()

}

func makeQueue() {
	println("woo")
	println("making queue")
	println("beep boop")
	queue = make([]int, 0)

	for i := 0; i < level+1; i++ {
		item := rand.Intn(4)
		queue = append(queue, item)
	}
	println("~")
	for i, s := range queue {
		println(i, s)
	}
	println("~")
	for _, v := range queue {

		for i := 0; i < 10; i++ {
			flash(v)
		}
		//time.Sleep(time.Second)
		dull(v)
	}

}
func press(c int) {
	if alive && len(queue) > 0 {
		if c == queue[0] {
			queue = queue[1:]
			flash(c)
		} else {
			display.FillScreen(black)
			alive = false
		}
	} else if !alive {
		alive = true
		level = 0
		for i := 0; i < 4; i++ {
			dull(i)
		}
		makeQueue()
	}

}
func release(c int) {
	if alive {
		dull(c)
		if len(queue) == 0 {
			level += 1
			makeQueue()
		}
	}

}

func flash(c int) {
	if c == 0 {
		display.FillRectangle(68, 120, 67, 120, blue)
	} else if c == 1 {
		display.FillRectangle(0, 120, 68, 120, yellow)
	} else if c == 2 {
		display.FillRectangle(68, 0, 67, 120, red)
	} else if c == 3 {
		display.FillRectangle(0, 0, 68, 120, green)
	}

}
func dull(c int) {
	if c == 0 {
		display.FillRectangle(68, 120, 67, 120, darkblue)
	} else if c == 1 {
		display.FillRectangle(0, 120, 68, 120, darkyellow)
	} else if c == 2 {
		display.FillRectangle(68, 0, 67, 120, darkred)
	} else if c == 3 {
		display.FillRectangle(0, 0, 68, 120, darkgreen)
	}
}
