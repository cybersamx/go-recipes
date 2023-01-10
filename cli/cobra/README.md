# Configurations Loading using Viper and PFlag

An example of using the viper and cobra packages to build a cli program.

The program takes in environment variables, command, command arguments, and pflags with the following pattern.

`ENV_NAME=ENV_VALUE APPNAME COMMAND ARG --PFLAG`

## Setup

1. Build the program.

   ```bash
   $ make
   ```

1. Run the program

   ```bash
   $ bin/cobra
   $ # Run the program with help subcommand, we get a help screen with the short descriptor of all commands and pflags.
   $ bin/cobra help
   ```
1. Try out different arguments.

   ```bash
   $ bin/cobra -a localhost:7000 list -lr bucket/folder
   list args: bucket/folder
   Config: &{localhost:7000 admin password false true true}
   $ CY_ADDR=example.com:1234 bin/cobra list -lr bucket/folder
   list args: bucket/folder
   Config: &{example.com:1234 admin password false true true}
   $ CY_ADDR=example.com:1234 bin/cobra -a localhost:7000 list -lr bucket/folder
   list args: bucket/folder
   Config: &{localhost:7000 admin password false true true}
   ```

## Reference and Credits

* [Github: Cobra](https://github.com/spf13/cobra)
* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
