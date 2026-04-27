package abstract_factory

import "testing"

func TestUIFactory(t *testing.T) {
	tests := []struct {
		name             string
		factory          UIFactory
		label            string
		expectedButton   string
		expectedCheckbox string
	}{
		{
			name:             "light theme",
			factory:          &LightThemeFactory{},
			label:            "Submit",
			expectedButton:   "[Light Button: Submit]",
			expectedCheckbox: "[Light Checkbox: Submit]",
		},
		{
			name:             "dark theme",
			factory:          &DarkThemeFactory{},
			label:            "Accept",
			expectedButton:   "[Dark Button: Accept]",
			expectedCheckbox: "[Dark Checkbox: Accept]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := tt.factory.CreateButton(tt.label)
			if got := btn.Render(); got != tt.expectedButton {
				t.Errorf("Button.Render() = %q, want %q", got, tt.expectedButton)
			}

			cb := tt.factory.CreateCheckbox(tt.label)
			if got := cb.Render(); got != tt.expectedCheckbox {
				t.Errorf("Checkbox.Render() = %q, want %q", got, tt.expectedCheckbox)
			}
		})
	}
}
