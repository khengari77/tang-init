package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/khengari77/tang-init/generator"
	"github.com/khengari77/tang-init/profiles"
)

func main() {
	boards := profiles.All()

	var projName string
	var boardID string
	var projType string

	boardOpts := make([]huh.Option[string], len(boards))
	for i, b := range boards {
		boardOpts[i] = huh.NewOption(b.Name, boards[i].ID)
	}

	theme := huh.ThemeCatppuccin()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Prompt("? ").
				Value(&projName).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("project name cannot be empty")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("Which board are you using?").
				Options(boardOpts...).
				Value(&boardID),
		),
	).WithTheme(theme)

	err := form.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var selectedBoard profiles.BoardProfile
	for _, b := range boards {
		if b.ID == boardID {
			selectedBoard = b
			break
		}
	}

	projType = "blinky"
	if selectedBoard.HasHDMI {
		typeForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select project template").
					Options(
						huh.NewOption("Basic Blinky", "blinky"),
						huh.NewOption("HDMI Video Boilerplate", "hdmi"),
					).
					Value(&projType),
			),
		).WithTheme(theme)

		if err := typeForm.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Generating project...")
	if err := generator.Generate(projName, projType, selectedBoard); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nDone! cd %s && make flash\n", projName)
}
