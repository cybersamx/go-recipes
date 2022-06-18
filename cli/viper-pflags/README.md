# Configurations Loading using Viper and PFlag

An example of using the viper and pflag packages to load configurations into an app with the following precedence:

* CLI arguments (highest precedence)
* environment variables
* configuration file
* default values

The pflag package allows POSIX/GNU style `--` flags.

## Setup

1. Run the tests.

   ```bash
   $ make test
   ```

## Reference and Credits

* [Github: Viper](https://github.com/spf13/viper)
* [Github: PFlag](https://github.com/spf13/pflag)
