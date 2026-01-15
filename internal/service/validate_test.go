package service

import "testing"

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      string
		wantError bool
	}{
		{"ok_trim", "  Дима  ", "Дима", false},
		{"empty", "", "", true},
		{"spaces_only", "   ", "", true},
		{"too_short_2", "аб", "", true},
		// 3 символа — ок
		{"min_len_ok", "abc", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateName(tt.in)
			if tt.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%q)", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestValidateDist(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      float64
		wantError bool
	}{
		{"ok_int", "5", 5, false},
		{"ok_float_dot", "5.5", 5.5, false},
		{"ok_float_comma", "5,5", 5.5, false},
		{"ok_trim", " 10 ", 10, false},

		{"not_number", "abc", 0, true},
		{"negative", "-1", 0, true},
		{"zero", "0", 0, true},
		{"too_big", "201", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateDist(tt.in)
			if tt.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil (got=%v)", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
