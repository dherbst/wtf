// Package sentry calls the sentry api for the configured projects to display the issues.
package sentry

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// HelpText is displayed when the user presses the `/` key when this widget has focus.
const HelpText = `
 Keyboard commands for Sentry:
`

// Widget is the display view for this widget.
type Widget struct {
	wtf.KeyboardWidget
	wtf.TextWidget

	results  []Issue
	settings *Settings
}

// NewWidget is called to return an instance of the sentry widget.
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: wtf.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.init()

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// HelpText returns the text to display when the user asks for help.
func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

// Refresh fetches the information from the sentry api and displays it in the widget view.
func (widget *Widget) Refresh() {
	results, err := widget.GetIssues()
	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}

	widget.results = results

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(widget.CommonSettings.Title, widget.renderIssues(widget.results), false)
}

func (widget *Widget) init() {

}

// renderIssues draws the rows of issues
func (widget *Widget) renderIssues(issues []Issue) string {
	view := "[red]Issues[white]\n"

	if len(issues) == 0 {
		view += "No Issues\n"
		return view
	}

	view += "[red]XY[white] - title here\n"

	return view
}
