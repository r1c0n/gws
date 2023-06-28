import os
import shutil
import zipfile

# create bin folder if it doesn't exist
if not os.path.exists("./bin"):
    os.mkdir("./bin")
    print("Bin folder created")

# build project
os.system("go build -o ./bin/gws.exe")
print("Project files built")

# create config.json
config_data = '''{
    "port": ":8080",
    "domain": "localhost",
    "static_dir": "html",
    "tls_config": {
        "cert_file": "server.crt",
        "key_file": "server.key"
    },
    "repo_config": {
        "version": "1.2.0",
        "author": "recon (contact@mail.recon.best)",
        "product": "Gamma Web Server",
        "repository": "https://github.com/gamma-gws/gws"
    }
}'''

with open("./bin/config.json", "w") as config_file:
    config_file.write(config_data)

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
