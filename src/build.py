import shutil
import zipfile
import json
import os
import logging
import argparse
import psutil
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


def build_project():
    """Build the project files."""
    os.system("go build -buildmode=exe -o ./bin/gws.exe")
    logging.info("Project files built")


def create_config_file(enable_ssl, enable_logging_middleware, enable_gzip_middleware):
    """Create the 'config.json' file with the given repository configuration."""
    config_data = {
        "port": ":80",
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


def main(run, no_deploy, enable_ssl):
    try:
        check_and_close_process("gws.exe")
        create_bin_folder()
        build_project()
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
            os.system("run.bat")
            logging.info("run.bat executed")
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
    args = parser.parse_args()

    enable_logging_middleware = "logging" in args.middleware or "all" in args.middleware
    enable_gzip_middleware = "gzip" in args.middleware or "all" in args.middleware

    main(args.run, args.no_deploy, args.enable_ssl)
