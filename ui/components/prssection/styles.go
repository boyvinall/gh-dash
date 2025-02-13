package prssection

import "github.com/charmbracelet/lipgloss"

var (
	ciCellWidth        = lipgloss.Width(" CI ")
	linesCellWidth     = lipgloss.Width(" 123450 / -123450 ")
	updatedAtCellWidth = lipgloss.Width("2mo ago")
	prRepoCellWidth    = 30
	prAuthorCellWidth  = 15
	ContainerPadding   = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)
)
