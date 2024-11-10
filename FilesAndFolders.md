# Files and Folders

| name                            | content                                          |
|---------------------------------|--------------------------------------------------|
| [./_test](./_test)              | automatic target code tests                      |
| `.code_snippets/`               | unused code                                      |
| `.github/`                      |                                                  |
| `.idea/`                        | GoLand settings                                  |
| `.vscode/`                      | vsCode settings                                  |
| [./cmd/cui](./cmd/_cui)         | (do not use) command user interface tryout code  |
| [./cmd/stim](./cmd/_stim)       | (do not use) target stimulation tool tryout code |
| [./cmd/trice](./cmd/trice)      | `trice` tool command Go sources                  |
| [./docs](./docs)                | documentation                                    |
| [./examples](./examples)        | example target projects                          |
| `./examples/*_inst/temp/`       | project binary logfiles                          |
| [./internal](./internal)        | `trice` tool internal Go packages                |
| [./pkg](./pkg)                  | `trice` tool common Go packages                  |
| [./src/](./src)                 | C sources for trice instrumentation              |
| `super-linter.report/`          |                                                  |
| [./third_party](./third_party)  | external components                              |
| `_config.yml`                   | unused                                           |
| `.clang-format`                 | See below                                        |
| `.clang-format-ignore`          | See below                                        |
| `.editorconfig`                 | See below                                        |
| `.git/`                         | version control data base                        |
| `.gitattributes`                | See below                                        |
| `.gitignore`                    |                                                  |
| `.goreleaser.yml`               | goreleaser configuration                         |
| `.travis.yml`                   |                                                  |
| `AUTHORS.md`                    | contributors                                     |
| `CHANGELOG.md`                  |                                                  |
| `CODE_OF_CONDUCT.md`            |                                                  |
| `CONTRIBUTING.md`               |                                                  |
| `FilesAndFolders.md`            | this file                                        |
| `go.mod`                        |                                                  |
| `go.sum`                        |                                                  |
| `GoInfos.txt`                   |                                                  |
| `LICENSE.md`                    |                                                  |
| `README.md`                     |                                                  |
| `branchesInfo.md`               |                                                  |
| `coverage.out`                  |                                                  |
| `dist/`                         | created by goreleaser                            |
| `fmtcoverage.html`              |                                                  |

## `.clang-format`

*Contributor: [Sazerac4](https://github.com/Sazerac4)*

Sazerac4 commented Aug 29, 2024:
I have a code formatter when I make changes to my application but I would like to keep the style of the library when modifying.
I couldn't find a code formatter, is there a tool used? If not, I propose this to provide one as an example by using clang-format.

```bash
# I have created a default style :
clang-format -style=llvm -dump-config > .clang-format
# Then format the code:
find ./src  -name '*.c' -o  -name '*.h'| xargs clang-format-18 -style=file -i
```

The style of the example does not correspond to the original one. Configurations are necessary for this to be the case. Tags can be placed to prevent certain macros from being formatted

```C
int formatted_code;
// clang-format off
    void    unformatted_code  ;
// clang-format on
void formatted_code_again;
```

I have tuned some settings for clang-format :

```bash
- IndentWidth: 4  // original code size indentation
- ColumnLimit: 0  // avoid breaking long line (like macros)
- PointerAlignment: Left  // like original files (mostly)
```

With preprocessor indentation, the result is a bit strange in some cases. It's possible with the option IndentPPDirectives ([doc](https://releases.llvm.org/18.1.6/tools/clang/docs/ClangFormatStyleOptions.html)).

Staying as close as possible to a default version (LLVM in this case) makes it easier to regenerate the style if necessary.

See also: https://github.com/rokath/trice/pull/487#issuecomment-2318003072

## `.clang-format-ignore`

*Contributor: [Sazerac4](https://github.com/Sazerac4)*

Sazerac4 commented Aug 30, 2024:
I have added .clang-format-ignore to ignore formatting for specific files


## `.editorconfig`

*Contributor: [Sazerac4](https://github.com/Sazerac4)*

The`.editorconfig` file allows to better identify the basic style for every files. (endline, charset, ...). It is a file accepted by a wide list of IDEs and editors : [link](https://editorconfig.org/#file-format-details)
This addition is motivated by forgetting the end of line in the .gitattributes file.

## `.gitattributes`

*Contributor: [Sazerac4](https://github.com/Sazerac4)*

With the`.gitattributes` file avoid problems with "diff" and end of lines. [Here](https://www.aleksandrhovhannisyan.com/blog/crlf-vs-lf-normalizing-line-endings-in-git/) is an article that presents the problem.

To fill the`.gitattributes`, I used the command below to view all the extensions currently used.

```bash
git ls-tree -r HEAD --name-only | perl -ne 'print $1 if m/\.([^.\/]+)$/' | sort -u
```
