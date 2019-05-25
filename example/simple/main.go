package main

import (
	"fmt"
	"log"

	"github.com/typical-go/feato"
)

var (
	// feature with specific ID
	featureHelloWorld = feato.NewFeature("hello-world")

	// unique feature
	feature1 = feato.NewUniqueFeature()

	// feature is enable by default
	feature2 = feato.NewUniqueFeature().EnableByDefault()

	// feature have specific index toggle to route to
	feature3 = feato.NewUniqueFeature().SetDefaultIndexToggle(feato.DisableIndexToggle)
)

func init() {
	// Set ToggleRouter to define which is enable
	feato.SetToggleRouter(feato.NewSimpleToggleRouter().
		SetIndexToggle("hello-world", feato.EnableIndexToggle).
		Enable(feature1).
		Disable(feature2).
		SetFeature(feature3, feato.EnableIndexToggle))
}

func main() {
	feato.Run(featureHelloWorld, func() (err error) {
		fmt.Println("Hello World")
		return
	})
	feato.Run(feature1, feature1Func, feature1AlternativeFunc)

	// ok is identify if function run
	ok, _ := feato.Run(feature2, feature2Func)
	if !ok {
		log.Println("Feature2 is not running")
	}

	log.Println(runFeature3())
}

func runFeature3() error {
	_, err := feato.Run(feature3,
		func() error {
			// By pass error to caller
			return fmt.Errorf("something error on feature 3")
		},
		func() error {
			return fmt.Errorf("feature 3 alternateve also got error")
		})

	return err
}

func feature1Func() error {
	fmt.Println("Feature 1 ")
	return nil
}

func feature1AlternativeFunc() error {
	fmt.Println("Feature 1 - Alternative")
	return nil
}

func feature2Func() error {
	fmt.Println("Feature 2")
	return nil
}
