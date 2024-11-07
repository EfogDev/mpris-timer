package ui

import (
	"fmt"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"log"
	"mpris-timer/internal/util"
	"slices"
	"strconv"
	"strings"
)

const (
	prefsMinWidth      = 300
	prefsMinHeight     = 200
	prefsDefaultWidth  = 420
	prefsDefaultHeight = 650
	sliderWidth        = 150
)

var prefsWin *adw.PreferencesWindow

func NewPrefsWindow() {
	prefsWin = adw.NewPreferencesWindow()
	prefsWin.SetTitle("Preferences")
	prefsWin.SetSizeRequest(prefsMinWidth, prefsMinHeight)
	prefsWin.SetDefaultSize(prefsDefaultWidth, prefsDefaultHeight)

	header := gtk.NewHeaderBar()
	box := gtk.NewBox(gtk.OrientationVertical, 8)
	box.SetVExpand(true)
	box.AddCSSClass("prefs-inner")

	NewPrefsWidgets(box)

	content := gtk.NewBox(gtk.OrientationVertical, 0)
	content.Append(box)

	scrolledWindow := gtk.NewScrolledWindow()
	scrolledWindow.SetVExpand(true)
	scrolledWindow.SetHExpand(true)
	scrolledWindow.SetOverlayScrolling(true)
	scrolledWindow.SetChild(content)

	outerBox := gtk.NewBox(gtk.OrientationVertical, 0)
	outerBox.Append(header)
	outerBox.Append(scrolledWindow)

	prefsWin.SetContent(outerBox)
	prefsWin.SetVisible(true)
	prefsWin.Activate()
	prefsWin.GrabFocus()
}

var (
	presetsOnRightSwitch *adw.SwitchRow
	presetsBox           *gtk.ListBox
)

func NewPrefsWidgets(parent *gtk.Box) {
	parent.SetSpacing(24)

	timerGroup := adw.NewPreferencesGroup()
	timerGroup.SetTitle("Timer")

	presetsGroup := adw.NewPreferencesGroup()
	presetsGroup.SetTitle("Interface")

	PopulateTimerGroup(timerGroup)
	PopulatePresetsGroup(presetsGroup)

	parent.Append(timerGroup)
	parent.Append(presetsGroup)
}

func PopulateTimerGroup(group *adw.PreferencesGroup) {
	textEntry := adw.NewEntryRow()
	volumeRow := adw.NewActionRow()
	volumeSlider := gtk.NewScaleWithRange(gtk.OrientationHorizontal, 0, 100, 1)

	soundSwitch := adw.NewSwitchRow()
	soundSwitch.SetTitle("Enable sound")
	soundSwitch.SetActive(util.UserPrefs.EnableSound)
	soundSwitch.Connect("notify::active", func() {
		util.SetEnableSound(soundSwitch.Active())
		volumeRow.SetSensitive(util.UserPrefs.EnableSound)
	})

	volumePreviewCtrl := gtk.NewGestureClick()
	volumePreviewCtrl.SetPropagationPhase(gtk.PhaseCapture)
	volumePreviewCtrl.ConnectReleased(func(_ int, _ float64, _ float64) {
		go util.PlaySound()
	})

	volumeRow.SetTitle("Sound volume")
	volumeRow.SetSubtitle(fmt.Sprintf("%v%%", int(util.Volume*100)))
	volumeRow.SetSensitive(util.UserPrefs.EnableSound)
	volumeRow.AddSuffix(volumeSlider)
	volumeRow.AddController(volumePreviewCtrl)

	volumeSlider.SetValue(util.Volume * 100)
	volumeSlider.SetSizeRequest(sliderWidth, 0)
	volumeSlider.ConnectChangeValue(func(scroll gtk.ScrollType, value float64) (ok bool) {
		util.SetVolume(value / 100)
		volumeRow.SetSubtitle(fmt.Sprintf("%v%%", int(util.Volume*100)))
		return false
	})

	notificationSwitch := adw.NewSwitchRow()
	notificationSwitch.SetTitle("Enable notification")
	notificationSwitch.SetActive(util.UserPrefs.EnableNotification)
	notificationSwitch.Connect("notify::active", func() {
		util.SetEnableNotification(notificationSwitch.Active())
		textEntry.SetSensitive(notificationSwitch.Active())
	})

	titleEntry := adw.NewEntryRow()
	titleEntry.SetTitle("Default title")
	titleEntry.SetText(util.UserPrefs.DefaultTitle)
	titleEntry.ConnectChanged(func() {
		util.SetDefaultTitle(titleEntry.Text())
	})

	textEntry.SetTitle("Default text")
	textEntry.SetText(util.UserPrefs.DefaultText)
	textEntry.SetSensitive(util.UserPrefs.EnableNotification)
	textEntry.ConnectChanged(func() {
		util.SetDefaultText(textEntry.Text())
	})

	color, err := util.RGBAFromHex(util.Color)
	if err != nil {
		log.Fatalf("unexpected: nil color, %v (%s)", err, util.UserPrefs.ProgressColor)
	}

	dialog := gtk.NewColorDialog()
	dialog.SetWithAlpha(false)
	colorSwitch := gtk.NewColorDialogButton(dialog)
	colorSwitch.AddCSSClass("color-picker-btn")
	colorSwitch.SetRGBA(color)
	colorSwitch.SetVExpand(false)
	colorRow := adw.NewActionRow()
	colorRow.AddSuffix(colorSwitch)
	colorRow.SetTitle("Progress color")

	colorSwitch.Connect("notify", func() {
		util.SetProgressColor(util.HexFromRGBA(colorSwitch.RGBA()))
	})

	group.Add(soundSwitch)
	group.Add(volumeRow)
	group.Add(notificationSwitch)
	group.Add(titleEntry)
	group.Add(textEntry)
	group.Add(colorRow)
}

func PopulatePresetsGroup(group *adw.PreferencesGroup) {
	newPresetBtn := gtk.NewButton()
	defaultPresetSelect := adw.NewComboRow()

	presetsOnRightSwitch = adw.NewSwitchRow()
	presetsOnRightSwitch.SetTitle("Presets on right side")
	presetsOnRightSwitch.SetSubtitle("Requires restart")
	presetsOnRightSwitch.SetSensitive(util.UserPrefs.ShowPresets)
	presetsOnRightSwitch.SetActive(util.UserPrefs.PresetsOnRight)
	presetsOnRightSwitch.Connect("notify::active", func() {
		util.SetPresetsOnRight(presetsOnRightSwitch.Active())
	})

	showPresetsSwitch := adw.NewSwitchRow()
	showPresetsSwitch.SetTitle("Show presets")
	showPresetsSwitch.SetSubtitle("Requires restart")
	showPresetsSwitch.SetActive(util.UserPrefs.ShowPresets)
	showPresetsSwitch.Connect("notify::active", func() {
		util.SetShowPresets(showPresetsSwitch.Active())
		presetsOnRightSwitch.SetSensitive(showPresetsSwitch.Active())
		defaultPresetSelect.SetSensitive(showPresetsSwitch.Active())
		presetsBox.SetVisible(showPresetsSwitch.Active())
		newPresetBtn.SetVisible(showPresetsSwitch.Active())
	})

	defaultPresetSelect.SetTitle("Default preset")
	defaultPresetSelect.SetModel(gtk.NewStringList(util.UserPrefs.Presets))
	selectedPos := slices.Index(util.UserPrefs.Presets, util.UserPrefs.DefaultPreset)
	defaultPresetSelect.SetSelected(uint(selectedPos))
	defaultPresetSelect.SetActivatable(true)
	defaultPresetSelect.SetSensitive(showPresetsSwitch.Active())
	defaultPresetSelect.Connect("notify::selected", func() {
		preset := util.UserPrefs.Presets[defaultPresetSelect.Selected()]
		util.SetDefaultPreset(preset)
	})

	presetsBox = gtk.NewListBox()
	presetsBox.SetVisible(util.UserPrefs.ShowPresets)
	presetsBox.AddCSSClass("presets-list")
	presetsBox.SetVExpand(true)

	RenderPresets([]string{})

	btnContent := adw.NewButtonContent()
	btnContent.SetLabel("")
	btnContent.SetIconName("insert-text")

	newPresetBtn.SetChild(btnContent)
	newPresetBtn.AddCSSClass("add-preset-btn")
	newPresetBtn.SetVisible(showPresetsSwitch.Active())
	newPresetBtn.ConnectClicked(func() {
		presets := append(util.UserPrefs.Presets, "00:00")
		util.SetPresets(presets)
		RenderPresets([]string{"00:00"})
	})

	footer := gtk.NewBox(gtk.OrientationHorizontal, 0)
	footer.SetHAlign(gtk.AlignCenter)
	footer.Append(newPresetBtn)

	group.Add(showPresetsSwitch)
	group.Add(presetsOnRightSwitch)
	group.Add(defaultPresetSelect)
	group.Add(presetsBox)
	group.Add(footer)
}

func RenderPresets(toAdd []string) {
	newPresets := toAdd
	if len(toAdd) == 0 {
		newPresets = util.UserPrefs.Presets
	}

	for _, preset := range newPresets {
		row := adw.NewActionRow()
		row.SetTitle(preset)

		container := gtk.NewListBoxRow()
		container.SetChild(row)

		box := gtk.NewBox(gtk.OrientationHorizontal, 16)
		box.SetVAlign(gtk.AlignCenter)
		box.AddCSSClass("presets-list-item")

		title := gtk.NewEditableLabel(preset)
		title.AddCSSClass("presets-list-title")
		title.SetVAlign(gtk.AlignStart)
		title.SetAlignment(0)
		title.SetHExpand(true)

		cleanTitle := func() {
			title.Activate()
			title.GrabFocus()
			title.SetText("00:01")
		}

		focusCtrl := gtk.NewEventControllerFocus()
		title.AddController(focusCtrl)

		focusCtrl.ConnectLeave(func() {
			octets := strings.Split(title.Text(), ":")
			if len(octets) < 1 && len(octets) > 3 {
				cleanTitle()
				return
			}

			var vals []int
			for _, v := range octets {
				val, err := strconv.Atoi(v)
				if err != nil || val < 0 || val > 59 {
					cleanTitle()
					return
				}

				vals = append(vals, val)
			}

			var newText string
			if len(vals) == 1 {
				newText = fmt.Sprintf("00:%02d", vals[0])
			} else if len(vals) == 2 {
				newText = fmt.Sprintf("%02d:%02d", vals[0], vals[1])
			} else if len(vals) == 3 {
				newText = fmt.Sprintf("%02d:%02d:%02d", vals[0], vals[1], vals[2])
			}

			if newText == "" {
				cleanTitle()
				return
			}

			var presets []string
			for _idx, p := range util.UserPrefs.Presets {
				if container.Index() == _idx {
					presets = append(presets, newText)
					continue
				}

				presets = append(presets, p)
			}

			title.SetText(newText)
			util.SetPresets(presets)
			prefsWin.GrabFocus()
		})

		btnContent := adw.NewButtonContent()
		btnContent.SetHExpand(false)
		btnContent.SetLabel("")
		btnContent.SetIconName("edit-delete")

		btn := gtk.NewButton()
		btn.SetChild(btnContent)
		btn.AddCSSClass("list-btn")
		btn.ConnectClicked(func() {
			var presets []string
			for _idx, p := range util.UserPrefs.Presets {
				if container.Index() == _idx {
					continue
				}

				presets = append(presets, p)
			}

			util.SetPresets(presets)
			presetsBox.Remove(container)
		})

		box.Append(title)
		box.Append(btn)

		row.SetChild(box)
		row.SetActivatableWidget(title)
		presetsBox.Append(container)
	}
}
