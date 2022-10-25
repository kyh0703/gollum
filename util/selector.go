package util

import "github.com/manifoldco/promptui"

type Selector struct{}

func (selector *Selector) SelectBox(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}
