package utils

import (
	"github.com/manifoldco/promptui"
)

func PromptInputAndValidate(label string, validate func(string) error) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	result, err := prompt.Run()
	LoggerGetInstance().Log("\n")
	if err != nil {
		return ""
	}
	return result
}

func PromptInputAndConfirm(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	LoggerGetInstance().Log("\n")
	if err != nil {
		return ""
	}
	return result
}

func PromptInputWithTemplate(label string, normal string, valid string, invalid string, success string, validate func(string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  normal,
		Valid:   valid,
		Invalid: invalid,
		Success: success,
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		return ""
	}
	return result
}

func PromptPassword(label string, validate func(string) error) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}
	result, err := prompt.Run()
	LoggerGetInstance().Log("\n")
	if err != nil {
		return ""
	}
	return result
}

func PromptSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	LoggerGetInstance().Log("\033[H\033[2J")
	if err != nil {
		return ""
	}
	return result
}
