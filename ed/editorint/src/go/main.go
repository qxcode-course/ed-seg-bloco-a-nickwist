package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	text   *List[*List[rune]] // a lista de linhas
	itLine *Node[*List[rune]] // iterador para a linha corrente
	itChar *Node[rune]        // iterador para o caracter do cursor
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.itChar = e.itLine.Value.Insert(e.itChar, r) // insere antes do elemento apontado pelo cursor
	e.itChar = e.itChar.Next()                    // move o cursor para próxima posição
}

func (e *Editor) KeyLeft() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itChar.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		e.itLine = e.itLine.Prev()      // Atualiza iterador de linha para linha anterior
		e.itChar = e.itLine.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	if e.itChar != e.itLine.Value.End() { // Se o cursor não está no fim da linha
		linhanova := NewList[rune]()
		linhaantiga := e.itLine.Value
		e.text.Insert(e.itLine.Next(), linhanova)
		
		for char := e.itChar; char != linhaantiga.End(); char = char.Next() {
			linhanova.Insert(linhanova.End(), char.Value)	
		}

		for e.itChar != linhaantiga.End() {
			e.itChar = linhaantiga.Erase(e.itChar)
		}

		e.itLine = e.itLine.Next()
		e.itChar = e.itLine.Value.Front()
		return
	}

	e.text.Insert(e.itLine.Next(), NewList[rune]()) // cria uma nova linha e insere abaixo da linha corrente
	e.itLine = e.itLine.Next() // vai pra próxima linha
	e.itChar = e.itLine.Value.Front() // move o cursor para o início da linha		

}

func (e *Editor) KeyRight() {
	if e.itChar != e.itLine.Value.End() {
		e.itChar = e.itChar.Next()
		return
	}

	if e.itLine != e.text.End() {
		e.itLine = e.itLine.Next()
		e.itChar = e.itLine.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	if e.itLine == e.text.Front() {
		return
	}	
		pos := 0
		for char := e.itLine.Value.Front(); char != e.itChar; char = char.Next() {
			pos++
		}

		e.itLine = e.itLine.Prev()

		e.itChar = e.itLine.Value.Front()
		
		for i := 0; i < pos && e.itChar != e.itLine.Value.End(); i++ {
			e.itChar = e.itChar.Next()
		}

	
}

func (e *Editor) KeyDown() {
	if e.itLine == e.text.End() {
		return
	}	
		pos := 0
		for char := e.itLine.Value.Front(); char != e.itChar; char = char.Next() {
			pos++
		}

		e.itLine = e.itLine.Next()

		e.itChar = e.itLine.Value.Front()
		
		for i := 0; i < pos && e.itChar != e.itLine.Value.End(); i++ {
			e.itChar = e.itChar.Next()
		}

	
}

func (e *Editor) KeyBackspace() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itLine.Value.Erase(e.itChar.prev)
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		linhaexcluida := e.itLine
		line := e.itLine.Value
		e.itLine = e.itLine.Prev()
		e.itChar = e.itLine.Value.End()
		for char := line.Front(); char != line.End(); char = char.Next() {
			letra := char.Value
			e.InsertChar(letra)
		}
		e.text.Erase(linhaexcluida)
	}
}

func (e *Editor) KeyDelete() {
	if e.itChar != e.itLine.Value.End() { // Se o cursor não está no fim da linha
		e.itLine.Value.Erase(e.itLine.Value.End().Prev())
		return
	}

	linhaexcluida := e.itLine.Next()
		line := e.itLine.Next().Value
		for char := line.Front(); char != line.End(); char = char.Next() {
			letra := char.Value
			e.InsertChar(letra)
		}
		e.text.Erase(linhaexcluida)
}

func main() { // Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.text = NewList[*List[rune]]()
	e.text.PushBack(NewList[rune]())
	e.itLine = e.text.Front()
	e.itChar = e.itLine.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.text.Front(); line != e.text.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '↲'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.itChar {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
