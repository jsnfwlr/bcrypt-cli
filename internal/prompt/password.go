package prompt

import (
	"errors"

	"github.com/manifoldco/promptui"

	"github.com/jsnfwlr/bcrypt-cli/internal/feedback"
)

// Password - prompt the user to enter a secret that answers your question and return their answer. The value entered will be masked with asterisks.
//
// Params:
//   - question: the question to ask the user
func Password(label string, allowBlank bool) string {
	prompter := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}
	if !allowBlank {
		prompter.Validate = func(i string) error {
			if i == "" {
				return errors.New("a non-blank value is required")
			}
			return nil
		}
	}
	result, err := prompter.Run()
	feedback.HandleFatalErr(err)

	return result
}
