package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordWrap(t *testing.T) {
	tests := []struct {
		description string
		input       struct {
			text   string
			column int
		}
		want string
	}{
		{
			input: struct {
				text   string
				column int
			}{
				text:   "The quick brown fox jumped over the lazy dog.",
				column: 10,
			},
			want: "The quick\nbrown fox\njumped\nover the\nlazy dog.",
		},
		{
			input: struct {
				text   string
				column int
			}{
				text:   "No wraps.",
				column: 20,
			},
			want: "No wraps.",
		},
		{
			input: struct {
				text   string
				column int
			}{
				text:   "geiyequau9pa5OoNg9poomoh6daungoh4Shooteexaeyiovie7eesh2tee3thitheeT3ierij9einisoad9iavaquoh6ga5dooghooYiesi0om4quohf1tei9vaiShohfae7cohY9Fiesheedooso1eiraiw2bah",
				column: 80,
			},
			want: "geiyequau9pa5OoNg9poomoh6daungoh4Shooteexaeyiovie7eesh2tee3thitheeT3ierij9einisoad9iavaquoh6ga5dooghooYiesi0om4quohf1tei9vaiShohfae7cohY9Fiesheedooso1eiraiw2bah",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := WordWrap(test.input.text, test.input.column)

			if !cmp.Equal(got, test.want) {
				t.Errorf("%v", cmp.Diff(got, test.want))
			}
		})
	}
}
