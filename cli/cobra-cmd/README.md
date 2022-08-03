# Configurations Loading using Viper and PFlag

An example of using the viper and cobra packages to build a cli program.

The program takes in environment variables, command, command parameters, and pflags with the following pattern.

`ENV_NAME=ENV_VALUE APPNAME COMMAND ARG --PFLAG`

## Setup

1. Build the program.

   ```bash
   $ make
   ```

2. Run the program

   ```bash
   $ # Run the ./cobra-cmd binary - see section Outputs
   $ #
   $ bin/cobra-cmd
   $ # Run the program with help subcommand, we get a help screen with the short descriptor of all commands and flags.
   $ bin/cobra-cmd help
   ```

## Outputs

<<<<<<< Updated upstream
| Command Execution                                    | Output                                                                                                  |
|------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| `./cobra-cmd`                                        | Output from the root command.                                                                           |
| `./cobra-cmd help`                                   | A help screen with the short descriptor of all commands and flags.                                      |
| `./cobra-cmd run help`                               | A help screen for the run subcommand. The long description will be used.                                |
| `./cobra-cmd run arg`                                | Output from the run command, which shows default value of format flag ie. yaml                          |
| `./cobra-cmd run arg --format json`                  | Output from the run command, which shows passed param value of format flag ie. json                     |
| `./cobra-cmd run arg -f json`                        | Output from the run command, which shows passed param value of f (short) flag ie. json                  |
| `CYBER_FORMAT=csv ./cobra-cmd run arg`               | Output from the run command, which shows the environment variable for format ie. csv                    |
| `CYBER_FORMAT=csv ./cobra-cmd run arg --format json` | Output from the run command, if both env and flag variables are present, flag variable takes precedence |
=======
| Command Executed                                     |
|------------------------------------------------------|
| `./cobra-cmd`                                        |
| `./cobra-cmd help`                                   |
| `./cobra-cmd run help`                               |
| `./cobra-cmd run arg`                                |
| `./cobra-cmd run arg --format json`                  |
| `./cobra-cmd run arg -f json`                        |
| `CYBER_FORMAT=csv ./cobra-cmd run arg`               |
| `CYBER_FORMAT=csv ./cobra-cmd run arg --format json` |
>>>>>>> Stashed changes

## Reference and Credits

* [Github: Cobra](https://github.com/spf13/cobra)
* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
