# ✨ Build Documentation ✨

Building Gamma Web Server is all done through [/src/build.py](/src/build.py). The script performs several tasks including building the project files, creating configuration files, and zipping contents for release.

## Usage 🚀

```bash
py build.py [--run] [--no-deploy] [--enable-ssl] [--debug] [--middleware <middleware_type> [<middleware_type> ...]]
```

### Arguments 📋

- `--run`: Run Gamma Web Server after build.
- `--no-deploy`: Do not zip contents for release.
- `--enable-ssl`: Enable SSL in configuration.
- `--middleware` : Enable middleware. Options include `logging`, `gzip`, or `all`.
- `--debug` : Build configuration for debugging. Runs `--run`, `--no-deploy`, and `--middleware all`.

## Functions 🛠️

### `check_and_close_process(process_name)`

Checks for a process with the given name and closes it if found.

### `create_bin_folder()`

Creates the 'bin' folder if it doesn't exist.

### `build_project()`

Builds the project files.

### `create_config_file(enable_ssl, enable_logging_middleware, enable_gzip_middleware)`

Creates the 'config.json' file with the given repository configuration.

### `copy_html_files()`

Copies the HTML template code to the 'bin/html' directory.

### `zip_bin_contents()`

Zips the contents of the 'bin' directory (excluding unnecessary files).

### `remove_gws_exe_tilde()`

Removes the 'gws.exe~' file if it exists.

### `main(run, no_deploy, enable_ssl)`

Main function to orchestrate the build process. It handles various exceptions such as file not found errors and JSON decoding errors.

## Example Usage 🚀

```bash
python build.py --run --enable-ssl --middleware all
```

This command will run the build script, start Gamma Web Server after the build, enable SSL in the configuration, and enable all of the available middleware.

`Note: Please make sure you have the necessary dependencies installed and the project structure set up correctly before running the script!`
