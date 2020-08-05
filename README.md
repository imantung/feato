[![Go Report Card](https://goreportcard.com/badge/github.com/imantung/feato)](https://goreportcard.com/report/github.com/imantung/feato)

# feato

[Simplest possible **fea**ture **to**ggle](https://www.thoughtworks.com/radar/techniques?blipid=202005002) Library in Go

## Usage

Convenience function using global instance
```go
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
```

Parseable to JSON or other format
```go
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
```


## References

- <https://medium.com/@dehora/feature-flags-smaller-better-faster-software-development-f2eab58df0f9>
- <https://martinfowler.com/articles/feature-toggles.html>
