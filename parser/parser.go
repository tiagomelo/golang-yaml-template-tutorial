// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package parser

import (
	"html/template"
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// For ease of unit testing.
var (
	parseFile           = template.ParseFiles
	openFile            = os.Open
	createFile          = os.Create
	ioReadAll           = io.ReadAll
	yamlUnmarshal       = yaml.Unmarshal
	executeTemplateFile = func(templateFile *template.Template, wr io.Writer, data any) error {
		return templateFile.Execute(wr, data)
	}
)

// valuesFromYamlFile extracts values from yaml file.
func valuesFromYamlFile(dataFile string) (map[string]interface{}, error) {
	data, err := openFile(dataFile)
	if err != nil {
		return nil, errors.Wrap(err, "opening data file")
	}
	defer data.Close()
	s, err := ioReadAll(data)
	if err != nil {
		return nil, errors.Wrap(err, "reading data file")
	}
	var values map[string]interface{}
	err = yamlUnmarshal(s, &values)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling yaml file")
	}
	return values, nil
}

// Parse replaces values present in the template file
// with values defined in the data file, saving the result
// as an output file.
func Parse(templateFile, dataFile, outputFile string) error {
	tmpl, err := parseFile(templateFile)
	if err != nil {
		return errors.Wrap(err, "parsing template file")
	}
	values, err := valuesFromYamlFile(dataFile)
	if err != nil {
		return err
	}
	output, err := createFile(outputFile)
	if err != nil {
		return errors.Wrap(err, "creating output file")
	}
	defer output.Close()
	err = executeTemplateFile(tmpl, output, values)
	if err != nil {
		return errors.Wrap(err, "executing template file")
	}
	return nil
}
