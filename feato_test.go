package feato_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/imantung/feato"
	"github.com/stretchr/testify/require"
)

func ExampleIsEnabled() {
	feato.RegisterGlobal([]*feato.Feature{
		{Name: "group01", Childs: []*feato.Feature{
			{Name: "name03", Flag: feato.Enabled},
			{Name: "name04", Flag: feato.Disabled},
		}},
	})

	if feato.IsEnabled("group01.name03") {
		fmt.Println("group01.name03 is enabled")
	}

	// Output: group01.name03 is enabled
}

func ExampleRegister() {
	flagMap := make(feato.FlagMap, 0)
	feato.Register(flagMap, []*feato.Feature{
		{Name: "name01", Flag: feato.Enabled},
		{Name: "name02", Flag: feato.Disabled},
		{Name: "group01", Childs: []*feato.Feature{
			{Name: "name03", Flag: feato.Enabled},
			{Name: "name04", Flag: feato.Disabled},
			{Name: "sub01", Flag: feato.Enabled, Childs: []*feato.Feature{
				{Name: "name05", Flag: feato.Enabled},
				{Name: "name06", Flag: feato.Disabled},
			}},
		}},
	})

	b, _ := json.MarshalIndent(flagMap, "", "   ")
	fmt.Println(string(b))

	// Output: {
	//    "group01.name03": true,
	//    "group01.name04": false,
	//    "group01.sub01": true,
	//    "group01.sub01.name05": true,
	//    "group01.sub01.name06": false,
	//    "name01": true,
	//    "name02": false
	// }
}

func TestRegisterGlobal(t *testing.T) {
	feato.RegisterGlobal([]*feato.Feature{
		{Name: "name01", Flag: feato.Enabled},
		{Name: "name02", Flag: feato.Disabled},
		{Name: "group01", Childs: []*feato.Feature{
			{Name: "name03", Flag: feato.Enabled},
			{Name: "name04", Flag: feato.Disabled},
			{Name: "sub01", Flag: feato.Enabled, Childs: []*feato.Feature{
				{Name: "name05", Flag: feato.Enabled},
				{Name: "name06", Flag: feato.Disabled},
			}},
		}},
	})
	defer feato.Instance.ClearAll()

	require.Equal(t, feato.FlagMap{
		"name01":               feato.Enabled,
		"name02":               feato.Disabled,
		"group01.name03":       feato.Enabled,
		"group01.name04":       feato.Disabled,
		"group01.sub01":        feato.Enabled,
		"group01.sub01.name05": feato.Enabled,
		"group01.sub01.name06": feato.Disabled,
	}, feato.Instance)

	b, _ := json.MarshalIndent(feato.Instance, "", "   ")
	fmt.Println(string(b))

	require.True(t, feato.IsEnabled("name01"))
	require.False(t, feato.IsEnabled("group01.sub01.name06"))
	require.False(t, feato.IsEnabled("not-found"))

}
