package virtualmachine

import (
	"../parser"
	"errors"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

type VirtualMachine struct {
	stack              []interface{}
	registers          map[string]interface{}
	function_registers map[string]interface{}
	instructions       []string
	library            map[string]func(*VirtualMachine)
}

func (self *VirtualMachine) GetStack() []interface{} {
	return self.stack
}

func (self *VirtualMachine) Load(instructions []string) {
	self.instructions = instructions
}

func MakeVM(instructions []string, library map[string]func(*VirtualMachine)) VirtualMachine {
	var empty_arr []interface{}
	return VirtualMachine{empty_arr, make(map[string]interface{}), make(map[string]interface{}), instructions, library}
}

func (self *VirtualMachine) Run() {
	for _, token := range self.instructions {
		if token == "+" {
			self.Op_add()
		} else if token == "-" {
			self.Op_sub()
		} else if token == "*" {
			self.Op_mul()
		} else if token == "/" {
			self.Op_div()
		} else if token == "|" {
			self.Op_print()
		} else if token == "<" {
			self.Op_store(false)
		} else if token == ">" {
			self.Op_load(false)
		} else if token == "$" {
	      self.Op_getln()
	    } else if token == "!" {
			self.Op_call(true)
		} else if token == "%" {
			self.Op_real_call()
		} else if token == "&" {
			self.Op_loop(false)
		} else if token == "@" {
			self.Op_create(false)
		} else if token == "." {
			self.Op_read(false)
		} else if token == "," {
			self.Op_write(false)
		} else if token == "^" {
			self.Op_top()
		} else if token == "=" {
			self.Op_eq()
		} else if token == ">>" {
			self.Op_greater()
		} else if token == "<<" {
			self.Op_less()
		} else {
			self.Push(token)
		}
		// self.function_registers = make(map[string]interface{})
	}
}

func (self *VirtualMachine) Local_Run() {
	for _, token := range self.instructions {
		if token == "+" {
			self.Op_add()
		} else if token == "-" {
			self.Op_sub()
		} else if token == "*" {
			self.Op_mul()
		} else if token == "/" {
			self.Op_div()
		} else if token == "|" {
			self.Op_print()
		} else if token == "<" {
			self.Op_store(true)
		} else if token == ">" {
			self.Op_load(true)
		} else if token == "$" {
	      self.Op_getln()
	    } else if token == "!" {
			self.Op_call(true)
		} else if token == "%" {
			self.Op_real_call()
		} else if token == "&" {
			self.Op_loop(true)
		} else if token == "@" {
			self.Op_create(false)
		} else if token == "." {
			self.Op_read(false)
		} else if token == "," {
			self.Op_write(false)
		} else if token == "^" {
			self.Op_top()
		} else if token == "=" {
			self.Op_eq()
		} else if token == ">>" {
			self.Op_greater()
		} else if token == "<<" {
			self.Op_less()
		} else {
			self.Push(token)
		}
	}
}

func (self *VirtualMachine) Push(item interface{}) {
	switch item.(type) {
	case string:
		if IsNumeric(item.(string)) {
			f, _ := strconv.ParseFloat(item.(string), 64)
			self.stack = append(self.stack, f)
		} else {
			self.stack = append(self.stack, item)
		}
	default:
		self.stack = append(self.stack, item)
	}
}

func (self *VirtualMachine) Pop() interface{} {
	var Pop interface{}
	Pop, self.stack = self.stack[len(self.stack)-1], self.stack[:len(self.stack)-1]
	return Pop
}

func (self *VirtualMachine) Op_add() {
	operand1 := self.Pop()
	operand2 := self.Pop()
	switch operand1.(type) {
	case int:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(int) + operand2.(int))
		case float64:
			self.Push(float64(operand1.(int)) + operand2.(float64))
		default:
			fmt.Println(errors.New("ERROR: Cannot add Non-Number to Number"))
			os.Exit(1)
		}
	case float64:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(float64) + float64(operand2.(int)))
		case float64:
			self.Push(operand1.(float64) + operand2.(float64))
		default:
			fmt.Println(errors.New("ERROR: Cannot add Non-Number to Number"))
			os.Exit(1)
		}
	case string:
		switch operand2.(type) {
		case string:
			self.Push(operand1.(string) + operand2.(string))
		default:
			fmt.Println(errors.New("ERROR: Cannot add Non-String to String"))
			os.Exit(1)
		}
	default:
		fmt.Println(errors.New("ERROR: Cannot add Non-Number or Non-Strings"))
		os.Exit(1)
	}
}

func (self *VirtualMachine) Op_sub() {
	operand1 := self.Pop()
	operand2 := self.Pop()
	switch operand1.(type) {
	case int:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(int) - operand2.(int))
		case float64:
			self.Push(float64(operand1.(int)) - operand2.(float64))
		default:
			fmt.Println(errors.New("ERROR: Cannot subtract Non-Number from Number"))
			os.Exit(1)
		}
	case float64:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(float64) - float64(operand2.(int)))
		case float64:
			self.Push(operand1.(float64) - operand2.(float64))
		default:
			fmt.Println(errors.New("ERROR: Cannot subtract Non-Number from Number"))
			os.Exit(1)
		}
	default:
		fmt.Println(errors.New("ERROR: Cannot subtract Non-Number"))
		os.Exit(1)
	}
}

func (self *VirtualMachine) Op_mul() {
	operand1 := self.Pop()
	operand2 := self.Pop()
	switch operand1.(type) {
	case int:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(int) * operand2.(int))
		case string:
			self.Push(strings.Repeat(operand2.(string), operand1.(int)))
		default:
			fmt.Println(errors.New("ERROR: Cannot multiply Number by Non-Number or Non-String"))
			os.Exit(1)
		}
	case float64:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(float64) * float64(operand2.(int)))
		case float64:
			self.Push(operand1.(float64) * operand2.(float64))
		case string:
			self.Push(strings.Repeat(operand2.(string), int(operand1.(float64))))
		default:
			fmt.Println(errors.New("ERROR: Cannot multiply Number by Non-Number or Non-String"))
			os.Exit(1)
		}
	case string:
		switch operand2.(type) {
		case int:
			self.Push(strings.Repeat(operand1.(string), operand2.(int)))
		case float64:
			self.Push(strings.Repeat(operand1.(string), int(operand2.(float64))))
		default:
			fmt.Println(errors.New("ERROR: Cannot multiply String by Non-Number"))
			os.Exit(1)
		}
	default:
		fmt.Println(errors.New("ERROR: Cannot multiply Non-Number or Non-String"))
		os.Exit(1)
	}
}

func (self *VirtualMachine) Op_div() {
	operand1 := self.Pop()
	operand2 := self.Pop()
	switch operand1.(type) {
	case int:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(int) / operand2.(int))
		default:
			fmt.Println(errors.New("ERROR: Cannot divide Number by Non-Number"))
			os.Exit(1)
		}
	case float64:
		switch operand2.(type) {
		case int:
			self.Push(operand1.(float64) / float64(operand2.(int)))
		case float64:
			self.Push(operand1.(float64) / operand2.(float64))
		default:
			fmt.Println(errors.New("ERROR: Cannot divide Number by Non-Number"))
			os.Exit(1)
		}
	default:
		fmt.Println(errors.New("ERROR: Cannot divide Non-Numbers"))
		os.Exit(1)
	}
}

func (self *VirtualMachine) Op_store(local bool) {
	operand1 := self.Pop()
	operand2 := self.Pop()
	if local {
		self.function_registers[operand1.(string)] = operand2
	} else {
		self.registers[operand1.(string)] = operand2
	}
}

func (self *VirtualMachine) Op_load(local bool) {
	operand := self.Pop()
	// fmt.Print("Loading: ")
	// fmt.Println(operand)
	if local {
    if _, ok := self.function_registers[operand.(string)]; ok {
      self.Push(self.function_registers[operand.(string)])
    } else {
      self.Push(self.registers[operand.(string)])
    }
	} else {
		self.Push(self.registers[operand.(string)])
	}
}

func (self *VirtualMachine) Op_call(local bool) {
	function := self.Pop().(string)
	parser := parser.MakeParser(function[1 : len(function)-1])
	self.Load(parser.Parse())
	if local {
		self.Local_Run()
	} else {
		self.Run()
	}
}

func (self *VirtualMachine) Op_top() {
	operand := self.Pop()
	self.Push(operand)
	self.Push(operand)
}

func (self *VirtualMachine) Op_ws() {
	self.Push(" ")
}

func (self *VirtualMachine) Op_loop(local bool) {
	condition := self.Pop()
	function := self.Pop()
	var value interface{}
	for {
		self.Push(condition)
		self.Op_call(local)
		value = self.Pop()
		if value.(float64) > 0 {
			self.Push(function)
			self.Op_call(local)
		} else {
			break
		}
	}
}

func (self *VirtualMachine) Op_read(local bool) {
  var index interface{} = self.Pop()
  var this interface{} = self.Pop()
  self.Push(this.(map[string]interface{})[index.(string)])
}

func (self *VirtualMachine) Op_write(local bool) {
    var index interface{} = self.Pop()
    var this interface{} = self.Pop()
    var value interface{} = self.Pop()
    this.(map[string]interface{})[index.(string)] = value
    self.Push(this)
}

func (self *VirtualMachine) Op_create(local bool) {
  self.Push(make(map[string]interface{}))
}
// def create(self, f=False):
//     self.Push({})

func (self *VirtualMachine) Op_getln() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  self.Push(text[:len(text)-1])
}

func (self *VirtualMachine) Op_greater() {
  var operand1 interface{} = self.Pop()
  var operand2 interface{} = self.Pop()
  if operand1.(float64) > operand2.(float64) {
    self.Push("1")
  } else {
    self.Push("0")
  }
}

func (self *VirtualMachine) Op_less() {
  var operand1 interface{} = self.Pop()
  var operand2 interface{} = self.Pop()
  if operand1.(float64) < operand2.(float64) {
    self.Push("1")
  } else {
    self.Push("0")
  }
}

func (self *VirtualMachine) Op_eq() {
  var operand1 interface{} = self.Pop()
  var operand2 interface{} = self.Pop()
  if operand1.(float64) == operand2.(float64) {
    self.Push("1")
  } else {
    self.Push("0")
  }
}

func (self *VirtualMachine) Op_print() {
	// fmt.Print("Printing...")
	operand := self.Pop()
	fmt.Print(operand)
}

func (self *VirtualMachine) Op_real_call() {
	f := self.Pop().(string)
	self.library[f](self)
}
