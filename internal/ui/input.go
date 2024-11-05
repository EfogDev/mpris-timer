package ui

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"log"
	"mpris-timer/internal/util"
	"slices"
	"strconv"
)

func setupTimeEntry(entry *gtk.Entry, prev *gtk.Widget, next *gtk.Widget, maxVal int, finish func()) {
	if maxVal <= 0 {
		maxVal = 59
	}

	entry.AddCSSClass("monospace")
	entry.AddCSSClass("entry")
	entry.AddCSSClass("timer-entry")
	entry.SetSensitive(true)
	entry.SetCanFocus(true)
	entry.SetCanTarget(true)
	entry.SetMaxWidthChars(2)
	entry.SetWidthChars(2)
	entry.SetOverflow(gtk.OverflowHidden)
	entry.SetHExpand(false)
	entry.SetHAlign(gtk.AlignCenter)
	entry.SetVAlign(gtk.AlignCenter)
	entry.SetAlignment(.5)
	entry.SetText("00")

	formatValue := func() {
		val := entry.Text()

		if len(val) == 0 {
			entry.SetText("00")
		}

		if len(val) == 1 {
			entry.SetText("0" + val)
		}
	}

	clickCtrl := gtk.NewGestureClick()
	clickCtrl.SetPropagationPhase(gtk.PhaseCapture)
	clickCtrl.ConnectReleased(func(nPress int, x, y float64) {
		_, _, ok := entry.SelectionBounds()

		if !ok {
			entry.SelectRegion(0, -1)
		}
	})

	focusCtrl := gtk.NewEventControllerFocus()
	focusCtrl.SetPropagationPhase(gtk.PhaseTarget)
	focusCtrl.ConnectLeave(func() {
		formatValue()

		if entry.HasFocus() {
			entry.SelectRegion(0, 0)
		}
	})

	kbCtrl := gtk.NewEventControllerKey()
	kbCtrl.SetPropagationPhase(gtk.PhaseCapture)
	kbCtrl.SetPropagationLimit(gtk.LimitNone)
	kbCtrl.ConnectKeyPressed(func(keyval, keycode uint, state gdk.ModifierType) (ok bool) {
		// allow some basic keys
		allowedKeyvals := []uint{
			gdk.KEY_Tab,
			gdk.KEY_ISO_Left_Tab,
			gdk.KEY_3270_BackTab,
			gdk.KEY_Return,
			gdk.KEY_RockerEnter,
			gdk.KEY_ISO_Enter,
			gdk.KEY_3270_Enter,
			gdk.KEY_KP_Enter,
			gdk.KEY_BackSpace,
			gdk.KEY_Delete,
			gdk.KEY_KP_Delete,
			gdk.KEY_Left,
			gdk.KEY_Right,
			gdk.KEY_Up,
			gdk.KEY_Down,
			gdk.KEY_Home,
			gdk.KEY_KP_Home,
			gdk.KEY_End,
			gdk.KEY_KP_End,
		}

		type shortcut struct {
			keyval []uint
			mask   gdk.ModifierType
			fn     func() bool
		}

		// allow some (unhandled) shortcuts
		allowedShortcuts := []shortcut{
			{
				// ^A = select all
				keyval: []uint{gdk.KEY_a},
				mask:   gdk.ControlMask,
			},
			{
				// space = focus next
				keyval: []uint{gdk.KEY_space, gdk.KEY_KP_Space},
				mask:   gdk.NoModifierMask,
				fn: func() bool {
					formatValue()
					next.GrabFocus()
					return true
				},
			},
			{
				// enter = start timer
				keyval: []uint{gdk.KEY_Return, gdk.KEY_KP_Enter, gdk.KEY_ISO_Enter, gdk.KEY_3270_Enter, gdk.KEY_RockerEnter},
				mask:   gdk.NoModifierMask,
				fn: func() bool {
					if finish == nil {
						return false
					}

					formatValue()
					finish()
					return true
				},
			},
			{
				// left = focus prev
				keyval: []uint{gdk.KEY_Left},
				mask:   gdk.NoModifierMask,
				fn: func() bool {
					if prev == nil {
						return false
					}

					_, _, selection := entry.SelectionBounds()
					if entry.Position() == 0 && !selection {
						formatValue()
						prev.GrabFocus()
						return true
					}

					return false
				},
			},
			{
				// right = focus next
				keyval: []uint{gdk.KEY_Right},
				mask:   gdk.NoModifierMask,
				fn: func() bool {
					_, _, selection := entry.SelectionBounds()
					if !selection && entry.Position() == len(entry.Text()) {
						formatValue()
						next.GrabFocus()
						return true
					}

					return false
				},
			},
		}

		for _, cfg := range allowedShortcuts {
			if slices.Contains(cfg.keyval, keyval) && cfg.mask == state {
				if cfg.fn != nil {
					return cfg.fn()
				}

				return false
			}
		}

		isNumber := util.IsGdkKeyvalNumber(keyval)
		if !isNumber && !slices.Contains(allowedKeyvals, keyval) {
			return true
		}

		// now we are interested only in numbers
		if !isNumber {
			return false
		}

		val := entry.Text()
		_, _, selectionPresent := entry.SelectionBounds()
		if len(val) >= 2 && !selectionPresent {
			return true
		}

		if len(val) != 1 {
			return false
		}

		// section finished
		newVal, err := strconv.Atoi(val + util.ParseKeyval(keyval))
		if err != nil {
			log.Printf("Error converting keyval to int: %v", err)
			return true
		}

		if newVal > maxVal {
			entry.SetText(util.NumToLabelText(maxVal))
			return true
		}

		return false
	})

	entry.ConnectChanged(func() {
		val := entry.Text()

		if len(val) == 2 && entry.HasFocus() {
			next.GrabFocus()
		}
	})

	entry.AddController(kbCtrl)
	entry.AddController(clickCtrl)
	entry.AddController(focusCtrl)
}
