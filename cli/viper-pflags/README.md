# Configurations Loading using Viper and PFlag

An example of implementing configurations in an aaplication using the `viper` and `pflag` packages.

<<<<<<< Updated upstream
* CLI arguments (top precedence)
=======
Configurations can be loaded into the application in the following ways with configurations set by cli arguments taking highest precedence.

* CLI arguments
>>>>>>> Stashed changes
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
   $ bin/viper-pflags
=======
1. Run the program.

   ```bash
   $ make run
>>>>>>> Stashed changes
   ```

## Reference and Credits

* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
