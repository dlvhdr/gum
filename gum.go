package main

import (
	"github.com/charmbracelet/gum/input"
	"github.com/charmbracelet/gum/join"
	"github.com/charmbracelet/gum/progress"
	"github.com/charmbracelet/gum/search"
	"github.com/charmbracelet/gum/spin"
	"github.com/charmbracelet/gum/style"
	"github.com/charmbracelet/gum/write"
)

// Gum is the command-line interface for Soda Gum.
type Gum struct {
	// Input provides a shell script interface for the text input bubble.
	// https://github.com/charmbracelet/bubbles/textinput
	//
	// It can be used to prompt the user for some input. The text the user
	// entered will be sent to stdout.
	//
	//   $ gum input --placeholder "What's your favorite gum?" > answer.text
	//
	Input input.Options `cmd:"" help:"Prompt for input."`

	// Write provides a shell script interface for the text area bubble.
	// https://github.com/charmbracelet/bubbles/textarea
	//
	// It can be used to ask the user to write some long form of text
	// (multi-line) input. The text the user entered will be sent to stdout.
	//
	//   $ gum write > output.text
	//
	Write write.Options `cmd:"" help:"Prompt for text"`

	// Search provides a fuzzy searching text input to allow filtering a list of
	// options to select one option.
	//
	// By default it will list all the files (recursively) in the current directory
	// for the user to choose one, but the script (or user) can provide different
	// new-line separated options to choose from.
	//
	// I.e. let's pick from a list of gum flavors:
	//
	//   $ cat flavors.text | gum search
	//
	Search search.Options `cmd:"" help:"Fuzzy search options."`

	// Spin provides a shell script interface for the spinner bubble.
	// https://github.com/charmbracelet/bubbles/spinner
	//
	// It is useful for displaying that some task is running in the background
	// while consuming it's output so that it is not shown to the user.
	//
	// For example, let's do a long running task:
	//   $ sleep 5
	//
	// We can simply prepend a spinner to this task to show it to the user,
	// while performing the task / command in the background.
	//
	//   $ gum spin -t "Taking a nap..." -- sleep 5
	//
	// The spinner will automatically exit when the task is complete.
	Spin spin.Options `cmd:"" help:"Show spinner while executing a command."`

	// Progress provides a shell script interface for the progress bubble.
	// https://github.com/charmbracelet/bubbles/progress
	//
	// It's useful for indicating that something is happening in the background
	// that will end after some set time.
	//
	Progress progress.Options `cmd:"" help:"Show a progress bar for some amount of time."`

	// Style provides a shell script interface for Lip Gloss.
	// https://github.com/charmbracelet/lipgloss
	//
	// It allows you to use Lip Gloss to style text without needing to use Go.
	// All of the styling options are available as flags.
	//
	// Let's make some text glamorous using bash:
	//
	//   $ gum style \
	//  	--foreground "#FF06B7" --border "double" --align "center" \
	//  	--width 50 --margin 2 --padding "2 4" \
	//  	"Bubble Gum (1¢)" "So sweet and so fresh\!"
	//
	//
	//    ╔══════════════════════════════════════════════════╗
	//    ║                                                  ║
	//    ║                                                  ║
	//    ║                 Bubble Gum (1¢)                  ║
	//    ║              So sweet and so fresh!              ║
	//    ║                                                  ║
	//    ║                                                  ║
	//    ╚══════════════════════════════════════════════════╝
	//
	Style style.Options `cmd:"" help:"Style some text."`

	// Join provides a shell script interface for the lipgloss JoinHorizontal
	// and JoinVertical commands.
	//
	Join join.Options `cmd:"" help:"Join text horizontally or vertically"`
}
