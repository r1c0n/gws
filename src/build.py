import shutil
import zipfile
import json
import os
from pathlib import Path
import logging
import argparse
import psutil

logging.basicConfig(level=logging.INFO)

def check_and_close_process(process_name):
    for proc in psutil.process_iter():
        if proc.name() == process_name:
            logging.info(f"Closing {process_name} process (PID: {proc.pid})")
            proc.kill()

def create_bin_folder():
    bin_path = Path("./bin")
    if not bin_path.exists():
        bin_path.mkdir()
        logging.info("Bin folder created")

def build_project():
    os.system("go build -o ./bin/gws.exe")
    logging.info("Project files built")

def read_repo_config():
    with open("gws-data.json", "r") as data_file:
        data = json.load(data_file)
        return data.get("repo_config")

def create_config_file(repo_config):
    config_data = {
        "port": ":8080",
        "domain": "localhost",
        "static_dir": "html",
        "tls_config": {
            "enabled": False,
            "cert_file": "server.crt",
            "key_file": "server.key"
        },
        "repo_config": repo_config
    }

    with open("./bin/config.json", "w") as config_file:
        json.dump(config_data, config_file, indent=4)
    logging.info("Config created")

def copy_html_files():
    html_dest = Path("./bin/html")
    if html_dest.exists():
        shutil.rmtree(html_dest)
    shutil.copytree("html", html_dest)
    logging.info("Template code copied to bin")

def zip_bin_contents():
    release_zip_path = Path("./bin/Release.zip")
    if release_zip_path.exists():
        release_zip_path.unlink()

    with zipfile.ZipFile(release_zip_path, "w") as zip_file:
        for foldername, subfolders, filenames in os.walk("./bin"):
            for filename in filenames:
                file_path = Path(foldername) / filename
                arcname = file_path.relative_to("./bin")
                if arcname.name != "Release.zip" and arcname.name not in ["server.crt", "server.key"]:
                    zip_file.write(file_path, arcname)

    logging.info("Content zipped to Release.zip")

def remove_gws_exe_tilde():
    gws_exe_tilde_path = Path("./bin/gws.exe~")
    if gws_exe_tilde_path.exists():
        gws_exe_tilde_path.unlink()
        logging.info("gws.exe~ file removed")

def main(run_dev):
    try:
        check_and_close_process("gws.exe")
        create_bin_folder()
        build_project()
        repo_config = read_repo_config()
        create_config_file(repo_config)
        copy_html_files()
        zip_bin_contents()
        remove_gws_exe_tilde()
        logging.info("Build completed")

        if run_dev:
            os.system("run-dev.bat")
            logging.info("run-dev.bat executed")
    except Exception as e:
        logging.error(f"Build failed: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Build and deploy script")
    parser.add_argument("--run-dev", action="store_true", help="Run run-dev.bat after build")
    args = parser.parse_args()

    main(args.run_dev)
