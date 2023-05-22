// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package main

import (
	"fmt"
	"os"

	"tiago.com/parser"
)

func run() error {
	const templateFile = "template/template.yaml"
	const dataFile = "template/values.yaml"
	const outputFile = "parsed/parsed.yaml"
	if err := parser.Parse(templateFile, dataFile, outputFile); err != nil {
		return err
	}
	fmt.Printf("file %s was generated.\n", outputFile)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
