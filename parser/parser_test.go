// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package parser

import (
	"errors"
	"html/template"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name                    string
		mockParseFile           func(filenames ...string) (*template.Template, error)
		mockOpenFile            func(name string) (*os.File, error)
		mockCreateFile          func(name string) (*os.File, error)
		mockIoReadAll           func(r io.Reader) ([]byte, error)
		mockYamlUnmarshal       func(in []byte, out interface{}) (err error)
		mockExecuteTemplateFile func(templateFile *template.Template, wr io.Writer, data any) error
		expectedError           error
	}{
		{
			name: "happy path",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockIoReadAll: func(r io.Reader) ([]byte, error) {
				return []byte(""), nil
			},
			mockYamlUnmarshal: func(in []byte, out interface{}) (err error) {
				return nil
			},
			mockCreateFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockExecuteTemplateFile: func(templateFile *template.Template, wr io.Writer, data any) error {
				return nil
			},
		},
		{
			name: "error when parsing template",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("parsing template file: random error"),
		},
		{
			name: "error when opening file",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("opening data file: random error"),
		},
		{
			name: "error when reading file",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockIoReadAll: func(r io.Reader) ([]byte, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("reading data file: random error"),
		},
		{
			name: "error when parsing yaml file",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockIoReadAll: func(r io.Reader) ([]byte, error) {
				return []byte(""), nil
			},
			mockYamlUnmarshal: func(in []byte, out interface{}) (err error) {
				return errors.New("random error")
			},
			expectedError: errors.New("unmarshalling yaml file: random error"),
		},
		{
			name: "error when creating file",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockIoReadAll: func(r io.Reader) ([]byte, error) {
				return []byte(""), nil
			},
			mockYamlUnmarshal: func(in []byte, out interface{}) (err error) {
				return nil
			},
			mockCreateFile: func(name string) (*os.File, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("creating output file: random error"),
		},
		{
			name: "error when executing template file",
			mockParseFile: func(filenames ...string) (*template.Template, error) {
				return new(template.Template), nil
			},
			mockOpenFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockIoReadAll: func(r io.Reader) ([]byte, error) {
				return []byte(""), nil
			},
			mockYamlUnmarshal: func(in []byte, out interface{}) (err error) {
				return nil
			},
			mockCreateFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockExecuteTemplateFile: func(templateFile *template.Template, wr io.Writer, data any) error {
				return errors.New("random error")
			},
			expectedError: errors.New("executing template file: random error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parseFile = tc.mockParseFile
			openFile = tc.mockOpenFile
			createFile = tc.mockCreateFile
			ioReadAll = tc.mockIoReadAll
			yamlUnmarshal = tc.mockYamlUnmarshal
			executeTemplateFile = tc.mockExecuteTemplateFile
			err := Parse("templateFile", "dataFile", "outputFile")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no error, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v", got nil`, tc.expectedError)
				}
				require.Nil(t, err)
			}
		})
	}
}
