import os
import shutil
import zipfile
import json

# create bin folder if it doesn't exist
if not os.path.exists("./bin"):
    os.mkdir("./bin")
    print("Bin folder created")

# build project
os.system("go build -o ./bin/gws.exe")
print("Project files built")

# read repo_config from gws-data.json
with open("gws-data.json", "r") as data_file:
    data = json.load(data_file)
    repo_config = data.get("repo_config")

# create config.json
config_data = {
    "port": ":8080",
    "domain": "localhost",
    "static_dir": "html",
    "tls_config": {
        "cert_file": "server.crt",
        "key_file": "server.key"
    },
    "repo_config": repo_config
}

with open("./bin/config.json", "w") as config_file:
    json.dump(config_data, config_file, indent=4)

print("Config created")

# copy html folder to bin folder
html_dest = "./bin/html"
if os.path.exists(html_dest):
    shutil.rmtree(html_dest)
shutil.copytree("html", html_dest)
print("Template code copied to bin")

# delete existing Release.zip file
if os.path.exists("./bin/Release.zip"):
    os.remove("./bin/Release.zip")

# zip contents of bin folder to Release.zip
zip_file = zipfile.ZipFile("./bin/Release.zip", "w")

for foldername, subfolders, filenames in os.walk("./bin"):
    for filename in filenames:
        if filename != "Release.zip":
            file_path = os.path.join(foldername, filename)
            arcname = os.path.relpath(file_path, "./bin")
            zip_file.write(file_path, arcname)

zip_file.close()

print("Content zipped to Release.zip")

print("Build completed")
