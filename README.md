# AOC - Advent Of Code
This tool will automatically download the input file for [AOC's challenge](http://adventofcode.com/) and provide boilerplate code to open the input file so you can immediately start coding. 


### AOC Session Cookie is REQUIRED
[Instructions for retreving](https://github.com/AnthonyNixon/advent-of-code-boilerplate/blob/main/docs/setup/session.md)

Add this session to an environment variable named AOC_SESSION. (If you cannot set an environment variable, or don't want to, you can run the script with the flag --session <your session token> instead.

## Usage

#### Get
| Command                    | Description                                                                                      |
|----------------------------|--------------------------------------------------------------------------------------------------|
| `./aoc get`                | Bootstraps the challenge with your custom input and templated file, defaults to the current day. |
| `./aoc get <day_num>`      | Bootstraps the challenge for the <day_num> day of the current year                               |

Get Flags

| Flag        | Description                                                                   | Environment Variable |
|-------------|-------------------------------------------------------------------------------|----------------------|
| `--session` | Sets the required session token value | `AOC_SESSION` |
| `--lang`    | Sets the Default Language                                                     | `AOC_LANG` |
| `--year`    | Sets the year to pull input from                                              | |

#### Templates
| Command               | Description                                               |
|-----------------------|-----------------------------------------------------------|
| `./aoc templates`     | Shows all current templates configured in the application |

#### Languages
| Command           | Description                                               |
|-------------------|-----------------------------------------------------------|
| `./aoc languages` | Shows all current languages configured in the application |

#### Version
| Command         | Description                                               |
|-----------------|-----------------------------------------------------------|
| `./aoc version` | Shows the current version information for the application |

#### Update
| Command                | Description                                                           |
|------------------------|-----------------------------------------------------------------------|
| `./aoc update`         | Automatically updates the application to the latest release on Github |
| `./aoc update <X.Y.Z>` | Updates the application to version X.Y.Z                              |


## Configuration

There are several configuration options for `aoc`.

| Environment Variable | Description                                                                                                          |
|----------------------|----------------------------------------------------------------------------------------------------------------------|
| `AOC_LANG`           | Sets the default language to bootstrap. Must be a valid option from `./aoc languages`                                |
| `AOC_SESSION`         | The session cookie contents to download unique puzzle inputs. This is required to be set or a flag must be utilized. |

## Contributing

`aoc` is **highly** customizable and extensible! Pull requests are always welcome. New templates and modifications to existing ones are also encouraged!

### Adding a Template

Adding a new built-in template to `aoc` is easy! The steps to create a new template are:
1. Create a new template file in the `templates/` directory following the [established pattern](templates/README.md).
2. Commit with a message containing `feat(YOUR_LANGUAGE): ...`.
   - This project is utilizing [Semantic Release](https://github.com/semantic-release/semantic-release) and commits must follow the [Commit Message format](https://semantic-release.gitbook.io/semantic-release/#commit-message-format).
3. Once merged and released your template will be packaged and available in the latest release!

## Development

To produce the binaries needed run:

```
make binaries
```