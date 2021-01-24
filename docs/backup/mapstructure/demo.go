package mapstructure

import (
	"fmt"
	m "github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func DecodeExample() {
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := m.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}

func ExampleDecodeMetaData() {
	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md m.Metadata
	var result Person
	config := &m.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := m.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err := decoder.Decode(input); err != nil {
		panic(err)
	}

	fmt.Printf("Unused keys: %#v", md.Unused)
	// Output:
	// Unused keys: []string{"email"}
}
