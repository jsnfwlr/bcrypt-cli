package prompt

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/jsnfwlr/bcrypt-cli/internal/feedback"
)

// Bool - prompt the user to answer a yes/no question and return their answer
//
// Params:
//   - question: the question to ask the user
func Bool(question string) bool {
	chooser := promptui.Prompt{
		Label:     question,
		IsConfirm: true,
		Validate: func(i string) error {
			if strings.ToLower(i) == "y" || strings.ToLower(i) == "n" {
				return nil
			}
			return errors.New("please entire either 'y' or 'n'")
		},
	}

	result, err := chooser.Run()
	feedback.HandleFatalErr(err)

	return strings.ToLower(result) == "y"
}
