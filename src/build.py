import shutil
import zipfile
import tarfile
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
RELEASE_TAR_PATH = BIN_PATH / "Release.tar.gz"
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
    if linux:  # use if on linux (tested & confirmed working on arch linux)
        os.system("go build -o ./bin/gws")
    else:
        logging.info("Building gws.exe")
        os.system("go build -buildmode=exe -o ./bin/gws.exe")
        logging.info("Bulding gwsvc.exe")
        os.system("go build -buildmode=exe -ldflags -H=windowsgui -o ./bin/gwsvc.exe")
    logging.info("Project files built")


def create_config_file(enable_ssl, enable_logging_middleware, enable_gzip_middleware):
    """Create the 'config.json' file with the given repository configuration."""
    if platform.system() == "Linux":
        port = ":8080"  # by default port 80 is protected on linux so we want to use 8080 instead
    else:
        if enable_ssl:
            port = ":443"  # use port 443 for HTTPS when SSL is enabled
        else:
            port = ":80"  # default to port 80 for HTTP
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
        "error_pages": {
            "enabled": True,
            "error_pages_dir": "html/errors",
            "pages": {
                "404": "404.html",
                "500": "500.html",
                "403": "403.html",
                "429": "429.html",
            },
        },
        "cors": {
            "enabled": False,
            "allowed_origins": ["*"],
            "allowed_methods": ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
            "allowed_headers": ["Content-Type", "Authorization", "X-Custom-Header"],
            "allow_credentials": False,
            "max_age": 3600,
        },
        "rate_limit": {
            "enabled": False,
            "requests_per_minute": 100,
            "burst": 20,
            "whitelist": ["127.0.0.1", "::1"],
            "exempt_paths": ["/html/", "/favicon/"],
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


def zip_bin_contents(linux=False):
    """Zip the contents of the 'bin' directory (or create a tar.gz if on Linux)."""
    if linux:
        if RELEASE_TAR_PATH.exists():
            RELEASE_TAR_PATH.unlink()

        with tarfile.open(RELEASE_TAR_PATH, "w:gz") as tar_file:
            # DO NOT REMOVE SUBFOLDERS! IT WILL BREAK THE BUILD SCRIPT!!
            for foldername, subfolders, filenames in os.walk(BIN_PATH):
                for filename in filenames:
                    file_path = Path(foldername) / filename
                    arcname = file_path.relative_to(BIN_PATH)
                    if arcname.name not in [
                        "Release.tar.gz",
                        "server.crt",
                        "server.key",
                        ".gws.exe.old",
                    ]:
                        tar_file.add(file_path, arcname)
        logging.info("Content archived to Release.tar.gz")
    else:
        if RELEASE_ZIP_PATH.exists():
            RELEASE_ZIP_PATH.unlink()

        with zipfile.ZipFile(RELEASE_ZIP_PATH, "w") as zip_file:
            # DO NOT REMOVE SUBFOLDERS! IT WILL BREAK THE BUILD SCRIPT!!
            for foldername, subfolders, filenames in os.walk(BIN_PATH):
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


def main(run, no_deploy, enable_ssl, linux, run_headless):
    try:
        check_and_close_process("gws.exe")
        create_bin_folder()
        build_project(linux=linux)
        create_config_file(
            enable_ssl,
            enable_logging_middleware,
            enable_gzip_middleware,
        )
        copy_html_files()
        if no_deploy:
            remove_gws_exe_tilde()
            logging.info("Build completed")
        else:
            zip_bin_contents(linux=linux)
            remove_gws_exe_tilde()
            logging.info("Build completed")

        if run:
            os.system("run.bat" if not linux else "./run.sh")
            logging.info("Run script executed")
        elif run_headless:
            os.system("run.bat -h")
            logging.info("Run headless script executed")
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
        "--run-headless", action="store_true", help="Run in headless mode"
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

    main(
        args.run,
        args.no_deploy,
        args.enable_ssl,
        args.linux,
        args.run_headless,
    )
