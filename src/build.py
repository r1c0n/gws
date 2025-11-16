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
RELEASE_ZIP_WINDOWS = BIN_PATH / "Release-Windows.zip"
RELEASE_TAR_LINUX = BIN_PATH / "Release-Linux.tar.gz"
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
        os.system("go build -o ./bin/gws.exe")
        logging.info("Building gwsvc.exe")
        os.system("go build -ldflags -H=windowsgui -o ./bin/gwsvc.exe")
    logging.info("Project files built")


def create_config_file(
    enable_ssl,
    enable_logging_middleware,
    enable_gzip_middleware,
    enable_cors,
    enable_rate_limit,
):
    """Create the 'config.json' file with the given repository configuration."""
    if platform.system() == "Linux":
        # By default port 80/443 are protected on Linux, use 8080/8443 instead
        port = ":8443" if enable_ssl else ":8080"
    else:
        # Windows: use standard HTTP (80) or HTTPS (443) ports
        port = ":443" if enable_ssl else ":80"
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
            "enabled": enable_cors,
            "allowed_origins": ["*"],
            "allowed_methods": ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
            "allowed_headers": ["Content-Type", "Authorization", "X-Custom-Header"],
            "allow_credentials": False,
            "max_age": 3600,
        },
        "rate_limit": {
            "enabled": enable_rate_limit,
            "requests_per_minute": 100,
            "burst": 20,
            "whitelist": ["127.0.0.1", "::1"],
            "exempt_paths": ["/html/", "/favicon/"],
        },
    }

    with open(CONFIG_FILE_PATH, "w") as config_file:
        json.dump(config_data, config_file, indent=4, ensure_ascii=False)
    logging.info("Config created")


def copy_html_files():
    """Copy the HTML template code to the 'bin/html' directory."""
    if HTML_DIR_PATH.exists():
        shutil.rmtree(HTML_DIR_PATH)
    shutil.copytree("html", HTML_DIR_PATH)
    logging.info("Template code copied to bin")


def zip_bin_contents(linux=False):
    """Zip the contents of the 'bin' directory with platform-specific naming."""
    if linux:
        release_path = RELEASE_TAR_LINUX
        if release_path.exists():
            release_path.unlink()

        with tarfile.open(release_path, "w:gz") as tar_file:
            # DO NOT REMOVE SUBFOLDERS! IT WILL BREAK THE BUILD SCRIPT!!
            for foldername, subfolders, filenames in os.walk(BIN_PATH):
                # Skip logs directory entirely
                if "logs" in Path(foldername).parts:
                    continue
                for filename in filenames:
                    file_path = Path(foldername) / filename
                    arcname = file_path.relative_to(BIN_PATH)
                    if arcname.name not in [
                        "Release-Linux.tar.gz",
                        "Release-Windows.zip",
                        "server.crt",
                        "server.key",
                        ".gws.exe.old",
                        "gws.exe",
                        "gwsvc.exe",
                    ] and not str(arcname).endswith(".log"):
                        tar_file.add(file_path, arcname)
        logging.info(f"Content archived to {release_path.name}")
    else:
        release_path = RELEASE_ZIP_WINDOWS
        if release_path.exists():
            release_path.unlink()

        with zipfile.ZipFile(release_path, "w") as zip_file:
            # DO NOT REMOVE SUBFOLDERS! IT WILL BREAK THE BUILD SCRIPT!!
            for foldername, subfolders, filenames in os.walk(BIN_PATH):
                # Skip logs directory entirely
                if "logs" in Path(foldername).parts:
                    continue
                for filename in filenames:
                    file_path = Path(foldername) / filename
                    arcname = file_path.relative_to(BIN_PATH)
                    if arcname.name not in [
                        "Release-Windows.zip",
                        "Release-Linux.tar.gz",
                        "server.crt",
                        "server.key",
                        ".gws.exe.old",
                        "gws",
                    ] and not str(arcname).endswith(".log"):
                        zip_file.write(file_path, arcname)

        logging.info(f"Content zipped to {release_path.name}")


def remove_gws_exe_tilde():
    """Remove the 'gws.exe~' file if it exists."""
    if GWS_EXE_TILDE_PATH.exists():
        GWS_EXE_TILDE_PATH.unlink()
        logging.info("gws.exe~ file removed")


def main(run, deploy, enable_ssl, linux, run_headless):
    try:
        if deploy:
            # Build for both Windows and Linux
            logging.info("Building for Windows...")
            check_and_close_process("gws.exe")
            create_bin_folder()
            build_project(linux=False)
            create_config_file(
                enable_ssl,
                enable_logging_middleware,
                enable_gzip_middleware,
                enable_cors,
                enable_rate_limit,
            )
            copy_html_files()
            zip_bin_contents(linux=False)
            remove_gws_exe_tilde()

            logging.info("Building for Linux...")
            build_project(linux=True)
            create_config_file(
                enable_ssl,
                enable_logging_middleware,
                enable_gzip_middleware,
                enable_cors,
                enable_rate_limit,
            )
            zip_bin_contents(linux=True)
            logging.info("Deployment builds completed for both platforms")
        else:
            # Regular build for current platform
            process_name = "gws" if linux else "gws.exe"
            check_and_close_process(process_name)
            create_bin_folder()
            build_project(linux=linux)
            create_config_file(
                enable_ssl,
                enable_logging_middleware,
                enable_gzip_middleware,
                enable_cors,
                enable_rate_limit,
            )
            copy_html_files()
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
    parser = argparse.ArgumentParser(
        description="Build and deploy script for Gamma Web Server"
    )
    parser.add_argument(
        "--run", action="store_true", help="Run Gamma Web Server after build"
    )
    parser.add_argument(
        "--deploy",
        action="store_true",
        help="Build release packages for both Windows and Linux",
    )
    parser.add_argument(
        "--enable-ssl", action="store_true", help="Enable SSL in config"
    )
    parser.add_argument(
        "--middleware",
        choices=["logging", "gzip", "cors", "ratelimit", "all"],
        nargs="+",
        default=[],
        help="Enable middleware (logging, gzip, cors, ratelimit, all)",
    )

    parser.add_argument(
        "--run-headless", action="store_true", help="Run in headless mode"
    )

    parser.add_argument(
        "--debug",
        action="store_true",
        help="Debug build configuration (--run, --middleware all)",
    )
    parser.add_argument(
        "--debug-ssl",
        action="store_true",
        help="Debug build with SSL (--run, --enable-ssl, --middleware all)",
    )
    parser.add_argument(
        "--linux",
        action="store_true",
        help="Compile the project to work with Linux",
    )
    args = parser.parse_args()

    if args.debug:
        args.run = True
        args.middleware = ["all"]

    if args.debug_ssl:
        args.run = True
        args.enable_ssl = True
        args.middleware = ["all"]

    enable_logging_middleware = "logging" in args.middleware or "all" in args.middleware
    enable_gzip_middleware = "gzip" in args.middleware or "all" in args.middleware
    enable_cors = "cors" in args.middleware or "all" in args.middleware
    enable_rate_limit = "ratelimit" in args.middleware or "all" in args.middleware

    main(
        args.run,
        args.deploy,
        args.enable_ssl,
        args.linux,
        args.run_headless,
    )
