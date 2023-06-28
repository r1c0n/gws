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
shutil.copytree("html", "./bin/html")
print("Template code copied to bin")

# delete existing Release.zip file
if os.path.exists("./bin/Release.zip"):
    os.remove("./bin/Release.zip")

# zip contents of bin folder to Release.zip
zipfile.ZipFile("./bin/Release.zip", "w").write("./bin", arcname=os.path.basename("./bin"))
print("Content zipped to Release.zip")

print("Build completed")
