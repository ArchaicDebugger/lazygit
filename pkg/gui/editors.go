package gui

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jesseduffield/gocui"
)

// we've just copy+pasted the editor from gocui to here so that we can also re-
// render the commit message length on each keypress
func (gui *Gui) commitMessageEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
	newlineKey, ok := gui.getKey(gui.Config.GetUserConfig().Keybinding.Universal.AppendNewline).(gocui.Key)
	if !ok {
		newlineKey = gocui.KeyEnter
	}
	newlineKey = gocui.KeyEnter

	gui.Log.Info(spew.Sdump(mod))
	gui.Log.Info(spew.Sdump(gocui.ModAlt))

	matched := true
	switch {
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyCtrlD || key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	case key == newlineKey && mod == gocui.ModAlt:
		v.EditNewLine()
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyCtrlU:
		v.EditDeleteToStartOfLine()
	case key == gocui.KeyCtrlA:
		v.EditGotoToStartOfLine()
	case key == gocui.KeyCtrlE:
		v.EditGotoToEndOfLine()
	default:
		if ch != 0 && mod == 0 {
			v.EditWrite(ch)
		} else {
			matched = false
		}
	}

	gui.RenderCommitLength()

	return matched
}

func (gui *Gui) defaultEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
	matched := true
	switch {
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyCtrlD || key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyCtrlU:
		v.EditDeleteToStartOfLine()
	case key == gocui.KeyCtrlA:
		v.EditGotoToStartOfLine()
	case key == gocui.KeyCtrlE:
		v.EditGotoToEndOfLine()
	default:
		if ch != 0 && mod == 0 {
			v.EditWrite(ch)
		} else {
			matched = false
		}
	}

	if gui.findSuggestions != nil {
		input := v.Buffer()
		suggestions := gui.findSuggestions(input)
		gui.setSuggestions(suggestions)
	}

	return matched
}
