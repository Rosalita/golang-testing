package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTellStory(t *testing.T) {
	emptyGopher := gopher{}
	emptyGopherExpected := "Once upon a time there was a gopher called . Their fur was  colour"

	gopherWithName := gopher{
		Name: "Rosie",
	}
	gopherWithNameExpected := "Once upon a time there was a gopher called Rosie. Their fur was  colour"

	gopherWithColour := gopher{
		Colour: "Red",
	}
	gopherWithColourExpected := "Once upon a time there was a gopher called . Their fur was Red colour"

	gopherWithNameAndColour := gopher{
		Name: "Rosalita",
		Colour: "Pink",
	}
	gopherWithNameAndColourExpected := "Once upon a time there was a gopher called Rosalita. Their fur was Pink colour"

	tests := []struct {
		input  gopher
		result string
	}{
		{emptyGopher, emptyGopherExpected},
		{gopherWithName, gopherWithNameExpected},
		{gopherWithColour, gopherWithColourExpected},
		{gopherWithNameAndColour, gopherWithNameAndColourExpected},
	}

	for _, test := range tests {
		result := tellStory(test.input)
		assert.Equal(t, test.result, result)
	}

}
