package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/jsnfwlr/bcrypt-cli/internal/feedback"
)

// ANP - All None Prompt - the overwrite setting for a field
// All - do not ask the user if they want to overwrite subsequent values, just ask the user for new values and overwrite the existing values
// None - do not ask the user if they want to overwrite subsequent values, just keep the original value, unless it is invalid
// Prompt - ask the user if they want to overwrite subsequent valid values, if they say yes, ask them for a new value, if they say no, keep the original value. Invalid existing values should skip the prompt and just ask the user for a new value.
type ANP int

const (
	All ANP = iota
	None
	Prompt
)

type ANPOption struct {
	Label       string
	Description string
	Icon        string
	Value       ANP
}

// Overwrite - prompt the user to select an overwrite setting and return their answer
// The value returned should be used in subsequent calls to OverwriteBool, OverwriteSelect
// and OverwriteText for fields valid existing values related to the settings this overwrite
// setting represents.
//
// Params:
//   - question: the question to ask the user about overwriting values
//   - hideHelp: hide the help text that explains how to answer the prompt
func Overwrite(question string, hideHelp bool) ANP {
	options := []ANPOption{
		{
			Label:       "All",
			Description: "Ignore all existing values and ask for new values for all fields",
			Icon:        "▸ ",
			Value:       All,
		},
		{
			Label:       "None",
			Description: "Only ask for values for new fields",
			Icon:        "▸ ",
			Value:       None,
		},
		{
			Label:       "Prompt",
			Description: "Ask if fields should be overwritten on a field by field basis",
			Icon:        "▸ ",
			Value:       Prompt,
		},
	}
	check := promptui.Select{
		Label: question,
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Label }}?",
			Active:   "{{ .Icon | bold | cyan }}{{ .Label }}",
			Inactive: "  {{ .Label }}",
			Details:  "{{ .Description | faint }}",
			Selected: fmt.Sprintf("%s {{ .Label | cyan }}", question),
		},
		HideHelp: hideHelp,
	}

	choice, _, err := check.Run()
	feedback.HandleFatalErr(err)

	return options[choice].Value
}

// OverwriteBool - prompt the user to answer a yes/no question and return their answer,
// with special consideration for how existing values should be overwritten
//
// Params:
//   - question:  the question to ask the user
//   - overwrite: the pre-defined overwrite setting.
//   - hideHelp:  hide the help text that explains how to answer the overwrite prompt
//   - original:  the original value of the field
//   - prompt:    the function to call to prompt the user for a new value for the field - should be prompt.Bool()
func OverwriteBool(question string, overwrite ANP, hideHelp, original bool, prompt func() bool) bool {
	switch overwrite {
	case None:
		return original
	case All:
		return prompt()

	default:
		check := promptui.Select{
			Label: question,
			Items: []stringOption{
				{
					Label:       "Yes",
					Description: "Enter a new value, overwriting the existing value",
					Icon:        "▸ ",
					Value:       "Yes",
				},
				{
					Label:       "No",
					Description: "Keep the existing value",
					Icon:        "▸ ",
					Value:       "No",
				},
			},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ .Label }}?",
				Active:   "{{ .Icon | bold | cyan }}{{ .Label }}",
				Inactive: "  {{ .Label }}",
				Details:  "{{ .Description | faint }}",
				Selected: fmt.Sprintf("%s {{ .Label | cyan }}", question),
			},
			HideHelp: hideHelp,
		}

		result, _, err := check.Run()
		feedback.HandleFatalErr(err)

		if result == 0 {
			return prompt()
		}

		return original
	}
}

// OverwriteSelect - prompt the user to select an option from a list that answers your
// question and return their answer, with special consideration for how existing values
// should be overwritten
//
// Params:
//   - question:  the question to ask the user
//   - overwrite: the pre-defined overwrite setting.
//   - hideHelp:  hide the help text that explains how to answer the overwrite prompt
//   - original:  the original value of the field
//   - prompt:    the function to call to prompt the user for a new value for the field - should be prompt.Select()
func OverwriteSelect(question string, overwrite ANP, hideHelp bool, original string, prompt func() string) string {
	switch overwrite {
	case None:
		return original
	case All:
		return prompt()

	default:
		check := promptui.Select{
			Label: question,
			Items: []stringOption{
				{
					Label:       "Yes",
					Description: "Enter a new value, overwriting the existing value",
					Icon:        "▸ ",
					Value:       "Yes",
				},
				{
					Label:       "No",
					Description: "Keep the existing value",
					Icon:        "▸ ",
					Value:       "No",
				},
			},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ .Label }}?",
				Active:   "{{ .Icon | bold | cyan }}{{ .Label }}",
				Inactive: "  {{ .Label }}",
				Details:  "{{ .Description | faint }}",
				Selected: fmt.Sprintf("%s {{ .Label | cyan }}", question),
			},
			HideHelp: hideHelp,
		}

		result, _, err := check.Run()
		feedback.HandleFatalErr(err)

		if result == 0 {
			return prompt()
		}

		return original
	}
}

// OverwriteText -prompt the user to enter some text that answers your question and
// return their answer, with special consideration for how existing values should be
// overwritten
//
// Params:
//   - question:  the question to ask the user
//   - overwrite: the pre-defined overwrite setting.
//   - hideHelp:  hide the help text that explains how to answer the overwrite prompt
//   - original:  the original value of the field
//   - prompt:    the function to call to prompt the user for a new value for the field - should be prompt.Text()
func OverwriteText(question string, overwrite ANP, hideHelp bool, original string, prompt func() string) string {
	switch overwrite {
	case None:
		return original
	case All:
		return prompt()
	default:
		check := promptui.Select{
			Label: question,
			Items: []stringOption{
				{
					Label:       "Yes",
					Description: "Enter a new value, overwriting the existing value",
					Icon:        "▸ ",
					Value:       "Yes",
				},
				{
					Label:       "No",
					Description: "Keep the existing value",
					Icon:        "▸ ",
					Value:       "No",
				},
			},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ .Label }}?",
				Active:   "{{ .Icon | bold | cyan }}{{ .Label }}",
				Inactive: "  {{ .Label }}",
				Details:  "{{ .Description | faint }}",
				Selected: fmt.Sprintf("%s {{ .Label | cyan }}", question),
			},
			HideHelp: hideHelp,
		}

		result, _, err := check.Run()
		feedback.HandleFatalErr(err)

		if result == 0 {
			return prompt()
		}

		return original
	}
}

// OverwritePassword - prompt the user to enter a secret that answers your question and return their answer,
// with special consideration for how existing values should be overwritten. The value entered will be masked
// with asterisks.
//
// Params:
//   - question:  the question to ask the user
//   - overwrite: the pre-defined overwrite setting.
//   - hideHelp:  hide the help text that explains how to answer the overwrite prompt
//   - original:  the original value of the field
//   - prompt:    the function to call to prompt the user for a new value for the field - should be prompt.Password()
func OverwritePassword(question string, overwrite ANP, hideHelp bool, original string, prompt func() string) string {
	switch overwrite {
	case None:
		return original
	case All:
		return prompt()

	default:
		check := promptui.Select{
			Label: question,
			Items: []stringOption{
				{
					Label:       "Yes",
					Description: "Enter a new value, overwriting the existing value",
					Icon:        "▸ ",
					Value:       "Yes",
				},
				{
					Label:       "No",
					Description: "Keep the existing value",
					Icon:        "▸ ",
					Value:       "No",
				},
			},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ .Label }}?",
				Active:   "{{ .Icon | bold | cyan }}{{ .Label }}",
				Inactive: "  {{ .Label }}",
				Details:  "{{ .Description | faint }}",
				Selected: fmt.Sprintf("%s {{ .Label | cyan }}", question),
			},
			HideHelp: hideHelp,
		}

		result, _, err := check.Run()
		feedback.HandleFatalErr(err)

		if result == 0 {
			return prompt()
		}

		return original
	}
}
