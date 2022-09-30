# Configurations Loading using Viper and PFlag

An example of implementing configurations in an application using the `viper` and `pflag` packages.

* CLI arguments (top precedence)
* environment variables
* configuration file
* default values

The `pflag` package can be configured to allow POSIX/GNU style `--` flags or the short form flags `-`.

## Setup

<<<<<<< Updated upstream
1. Build the program.

   ```bash
   $ make
   ```

2. Run the program

   ```bash
   $ CYBER_DIRECTORY=env bin/viper --postgres-url=arg
   Dump of config: {true 5000 arg env default}
   ```

## Reference and Credits

* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
