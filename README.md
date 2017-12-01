# AOC - Advent Of Code
This tool will automatically download the input file for [AOC's challenge](http://adventofcode.com/) and provide boilerplate code to open the input file so you can immediately start coding. 



### AOC Session Cookie is REQUIRED
[Instructions for retreving](https://blog.ajn.me/aoc-session/)

Add this session to an environment variable named AOC_SESSION. (If you cannot set an environment variable, or don't want to, you can run the script with the flag --session <your session token> instead.

## Usage

This will download today's challenge with go boilerplate

```./aoc get```

This will download the challenge for the 15th

```./aoc get 15```

This will download today's challenge with python boilerplate

```./aoc get --lang python```


#### All Custom Options:

```./aoc get [dayNum] [--lang (go|python)] [--year (2015|2016|2017)] [--session session-here]```


## Development

Glide is needed for dep management

```glide update```

To produce the binaries needed run

```
make binaries
```

To add a new boilerplate language
1. Add boilerplate.(language) to the languages folder
1. Update Makefile with ```	echo "const (language) = \"$$(cat languages/boilerplate.(language) | base64)\"" >> langs.go```
1. Update main.go with a new case 
    ```		
    case "(language)":
        f, err := os.Create(directory + "main.(language)")
        check(err)
        decoded, err := base64.StdEncoding.DecodeString((language))
        check(err)

        f.WriteString(string(decoded))
        f.Close()
    ```
1. Build and test 
1. Make a Pull Request with the changes
