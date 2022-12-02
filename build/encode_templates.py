#!/usr/bin/env python3
import os
import base64

directory = 'templates'

filesToEncode = []
for filename in os.listdir(directory):
    if filename == "README.md":
        continue
    file = os.path.join(directory, filename)
    # checking if it is a file
    if os.path.isfile(file):
        print(file)
        language, trailing = filename.split("_", 1)
        _, extension = trailing.split(".", 1)

        print(f'{language} - {extension}')
        filesToEncode.append((language, extension, file))

content = """\
package templating

func Initialize() {
    templates = make(map[string]string)
    fileExtensions = make(map[string]string)
"""
for (language, extension, file) in filesToEncode:
    with open(file, "rb") as encode_file:
        encoded_string = base64.b64encode(encode_file.read())
    print(f'{language} - {extension} - {file} -> {encoded_string.decode("utf-8")}')
    content = content + f'\n    templates["{language}"] = "{encoded_string.decode("utf-8")}"'
    content = content + f'\n    fileExtensions["{language}"] = "{extension}"'
content = content + "\n}"

with open("aoc-boilerplate/templating/init.go","w") as f:
    f.writelines(content)



