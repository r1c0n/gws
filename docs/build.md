# ✨ Build Documentation ✨

Building Gamma Web Server is all done through [/src/build.py](/src/build.py). The script performs several tasks including building the project files, creating configuration files, and zipping contents for release.

## Usage 🚀

```bash
py build.py [--run] [--run-headless] [--deploy] [--enable-ssl] [--linux] [--debug] [--debug-ssl] [--middleware <middleware_type> [<middleware_type> ...]]
```

### Arguments 📋

- `--run`: Run Gamma Web Server after build.
- `--run-headless`: Run Gamma Web Server headless version after build.
- `--deploy`: Build release packages for both Windows (Release-Windows.zip) and Linux (Release-Linux.tar.gz).
- `--enable-ssl`: Enable SSL in configuration (uses port 443).
- `--middleware` : Enable middleware. Options include `logging`, `gzip`, `cors`, `ratelimit`, or `all`.
- `--linux` : Compiles Gamma Web Server for Linux instead of Windows.
- `--debug` : Build configuration for debugging. Runs `--run` and `--middleware all` (port 80).
- `--debug-ssl` : Build configuration for debugging with SSL. Runs `--run`, `--enable-ssl`, and `--middleware all` (port 443).

## Functions 🛠️

### `check_and_close_process(process_name)`

Checks for a process with the given name and closes it if found.

### `create_bin_folder()`

Creates the 'bin' folder if it doesn't exist.

### `build_project(linux=linux)`

Builds the project files.

### `create_config_file(enable_ssl, enable_logging_middleware, enable_gzip_middleware)`

Creates the 'config.json' file with the given repository configuration.

### `copy_html_files()`

Copies the HTML template code to the 'bin/html' directory.

### `zip_bin_contents(linux=linux)`

Zips the contents of the 'bin' directory (excluding unnecessary files).

### `remove_gws_exe_tilde()`

Removes the 'gws.exe~' file if it exists.

### `main(run, no_deploy, enable_ssl, linux)`

Main function to orchestrate the build process. It handles various exceptions such as file not found errors and JSON decoding errors.

## Example Usage 🚀

```bash
python build.py --run --enable-ssl --middleware all
```

This command will run the build script, start Gamma Web Server after the build, enable SSL in the configuration, and enable all of the available middleware.

`Note: Please make sure you have the necessary dependencies installed and the project structure set up correctly before running the script!`
