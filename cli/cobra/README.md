# Configurations Loading using Viper and PFlag

An example of using the viper and cobra packages to build a cli program.

The program takes in environment variables, command, command arguments, and pflags with the following pattern.

`ENV_NAME=ENV_VALUE APPNAME COMMAND ARG --PFLAG`

## Setup

1. Build the program.

   ```bash
   $ make
   ```

2. Run the program

   ```bash
   $ bin/cobra
   $ # Run the program with help subcommand, we get a help screen with the short descriptor of all commands and pflags.
   $ bin/cobra help
   ```

## Outputs

| Command Execution                                | Output                                                                                           |
|--------------------------------------------------|--------------------------------------------------------------------------------------------------|
| `./cobra`                                        | Output from the root command, which calls the help command.                                      |
| `./cobra help`                                   | Help with short descriptors of all commands and flags.                                           |
| `./cobra run --help`                             | Help for the run subcommand. The long description will be used.                                  |
| `./cobra run arg`                                | Output from the run command.                                                                     |
| `./cobra run arg --format json`                  | Output from the run command, which shows flag-name `format` with a value of `json`.              |
| `./cobra run arg -f json`                        | Output from the run command, which shows (short) flag-name `format` with a value of `json`.      |
| `CYBER_FORMAT=csv ./cobra run arg`               | Output from the run command, which shows (env var) flag-name `format` with a value of `csv`.     |
| `CYBER_FORMAT=csv ./cobra run arg --format json` | Output from the run command, which shows flag-name `format` with a value of `json`. pflags > env |

## Reference and Credits

* [Github: Cobra](https://github.com/spf13/cobra)
* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
