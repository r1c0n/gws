import shutil
import zipfile
import json
import os
import logging
import argparse
import psutil
import platform
from pathlib import Path

# Please, read /docs/build.md to understand how this Python script works.

# Constants
BIN_PATH = Path("./bin")
CONFIG_FILE_PATH = BIN_PATH / "config.json"
HTML_DIR_PATH = BIN_PATH / "html"
RELEASE_ZIP_PATH = BIN_PATH / "Release.zip"
GWS_EXE_TILDE_PATH = BIN_PATH / "gws.exe~"

logging.basicConfig(level=logging.INFO)


def check_and_close_process(process_name):
    """Check for a process with the given name and close it if found."""
    for proc in psutil.process_iter():
        if proc.name() == process_name:
            logging.info(f"Closing {process_name} process (PID: {proc.pid})")
            proc.kill()


def create_bin_folder():
    """Create the 'bin' folder if it doesn't exist."""
    if not BIN_PATH.exists():
        BIN_PATH.mkdir()
        logging.info("Bin folder created")


def build_project(linux=False):
    """Build the project files."""
    if linux: # use if on linux (tested & confirmed working on arch linux)
        os.system("go build -o ./bin/gws")
    else:
        os.system("go build -buildmode=exe -o ./bin/gws.exe")
    logging.info("Project files built")


def create_config_file(enable_ssl, enable_logging_middleware, enable_gzip_middleware):
    """Create the 'config.json' file with the given repository configuration."""
    if platform.system() == "Linux":
        port = ":8080"  # by default port 80 is protected on linux so we want to use 8080 instead
    else:
        port = ":80"  # default to port 80 for other operating systems (basically just windows lol)
    config_data = {
        "port": port,
        "domain": "localhost",
        "static_dir": "html",
        "tls_config": {
            "enabled": enable_ssl,
            "cert_file": "server.crt",
            "key_file": "server.key",
        },
        "middleware": {
            "logging_middleware_enabled": enable_logging_middleware,
            "gzip_middleware_enabled": enable_gzip_middleware,
        },
    }

    with open(CONFIG_FILE_PATH, "w") as config_file:
        json.dump(config_data, config_file, indent=4)
    logging.info("Config created")


def copy_html_files():
    """Copy the HTML template code to the 'bin/html' directory."""
    if HTML_DIR_PATH.exists():
        shutil.rmtree(HTML_DIR_PATH)
    shutil.copytree("html", HTML_DIR_PATH)
    logging.info("Template code copied to bin")


def zip_bin_contents():
    """Zip the contents of the 'bin' directory (excluding unnecessary files)."""
    if RELEASE_ZIP_PATH.exists():
        RELEASE_ZIP_PATH.unlink()

    with zipfile.ZipFile(RELEASE_ZIP_PATH, "w") as zip_file:
        for foldername, subfolders, filenames in os.walk(
            BIN_PATH
        ):  # DO NOT REMOVE SUBFOLDERS! IT WILL BREAK THE BUILD SCRIPT!!
            for filename in filenames:
                file_path = Path(foldername) / filename
                arcname = file_path.relative_to(BIN_PATH)
                if arcname.name != "Release.zip" and arcname.name not in [
                    "server.crt",
                    "server.key",
                    ".gws.exe.old",
                ]:
                    zip_file.write(file_path, arcname)

    logging.info("Content zipped to Release.zip")


def remove_gws_exe_tilde():
    """Remove the 'gws.exe~' file if it exists."""
    if GWS_EXE_TILDE_PATH.exists():
        GWS_EXE_TILDE_PATH.unlink()
        logging.info("gws.exe~ file removed")


def main(run, no_deploy, enable_ssl, linux):
    try:
        check_and_close_process("gws.exe")
        create_bin_folder()
        build_project(linux=linux)
        create_config_file(
            enable_ssl, enable_logging_middleware, enable_gzip_middleware
        )
        copy_html_files()
        if no_deploy:
            remove_gws_exe_tilde()
            logging.info("Build completed")
        else:
            zip_bin_contents()
            remove_gws_exe_tilde()
            logging.info("Build completed")

        if run:
            os.system("run.bat" if not linux else "./run.sh")
            logging.info("Run script executed")
    except FileNotFoundError as e:
        logging.error(f"File not found: {e}")
    except json.JSONDecodeError as e:
        logging.error(f"JSON decoding failed: {e}")
    except Exception as e:
        logging.error(f"Build failed: {e}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Build and deploy script")
    parser.add_argument(
        "--run", action="store_true", help="Run Gamma Web Server after build"
    )
    parser.add_argument("--no-deploy", action="store_true", help="Don't zip contents")
    parser.add_argument(
        "--enable-ssl", action="store_true", help="Enable SSL in config"
    )
    parser.add_argument(
        "--middleware",
        choices=["logging", "gzip", "all"],
        nargs="+",
        default=[],
        help="Enable middleware (logging, gzip, all)",
    )
    parser.add_argument(
        "--debug",
        action="store_true",
        help="Debug build configuration (--run, --enable-ssl, --no-deploy, --middleware all)",
    )
    parser.add_argument(
        "--linux",
        action="store_true",
        help="Compile the project to work with Linux",
    )
    args = parser.parse_args()

    if args.debug:
        args.run = True
        args.enable_ssl = True
        args.no_deploy = True
        args.middleware = ["all"]

    enable_logging_middleware = "logging" in args.middleware or "all" in args.middleware
    enable_gzip_middleware = "gzip" in args.middleware or "all" in args.middleware

    main(args.run, args.no_deploy, args.enable_ssl, args.linux)
