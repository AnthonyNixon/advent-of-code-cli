# AOC Templates

These are the template files used by the program. They are compiled and added to the built binary as base64 encoded strings. Feel free to add additional template files to this directory! When compiled, the name of the template file is important. What comes before the `_template` is what the `language` attribute will be to access your template. The extension which is used will be the extension for the solution file generated. 

For example, the 'go' template is named `go_template.go` and will generate `main.go` each day. The 'python' template is named `python_template.py` and will generate `python.py` each day.