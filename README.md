<p align="center">
    <a href="https://www.gammaws.gq" target="_blank">
        <img src="./branding/gws-wordmark-01.png" alt="logo" width="125"/>
    </a>
</p>

<h1 align="center">Introducing <code>Gamma Web Server</code> 🚀</h1>
<h4 align="center">A lightweight and fast web server written in Go.</h4>

<p align="center">

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

<a href="https://github.com/r1c0n/gws/blob/main/LICENSE" target="blank">
<img src="https://img.shields.io/github/license/r1c0n/gws?style=flat-square" alt="gws license" />
</a>
<a href="https://github.com/r1c0n/gws/fork" target="blank">
<img src="https://img.shields.io/github/forks/r1c0n/gws?style=flat-square" alt="gws forks"/>
</a>
<a href="https://github.com/r1c0n/gws/stargazers" target="blank">
<img src="https://img.shields.io/github/stars/r1c0n/gws?style=flat-square" alt="gws stars"/>
</a>
<a href="https://github.com/r1c0n/gws/issues" target="blank">
<img src="https://img.shields.io/github/issues/r1c0n/gws?style=flat-square" alt="gws issues"/>
</a>
<a href="https://github.com/r1c0n/gws/pulls" target="blank">
<img src="https://img.shields.io/github/issues-pr/r1c0n/gws?style=flat-square" alt="gws pull-requests"/>
</a>
<a href="https://app.codacy.com/gh/r1c0n/gws/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade">
    <img src="https://app.codacy.com/project/badge/Grade/b4242484e7b840e6b1f5dd877723a8df"/>
</a>
<a href="https://twitter.com/intent/tweet?text=👋%20Check%20out%20this%20amazing%20webserver!%20https://github.com/r1c0n/gws"><img src="https://img.shields.io/twitter/url?label=Share%20on%20Twitter&style=social&url=https%3A%2F%2Fgithub.com%2Fr1c0n%2Fgws"></a>
</p>

![Alt](https://repobeats.axiom.co/api/embed/6eabbe5b07af02d9e866551848f1d4b0ac35c53a.svg 'Repobeats analytics image')

## 🔍 File Content

These are what the different files / directories contain in this repository.

| File/folder           | Description                                                           | Directory |
| --------------------- | --------------------------------------------------------------------- | --------- |
| `.github`             | Contains GitHub related files, such as workflows.                     | Yes       |
| `branding`            | Contains Gamma artwork.                                               | Yes       |
| `docs`                | Contains documentation for GWS.                                       | Yes       |
| `src`                 | Source code of Gamma Web Server.                                      | Yes       |
| `.all-contributorsrc` | All-Contributors bot configuration file                               | No        |
| `.gitattributes`      | Defines attribute rules for Git repository.                           | No        |
| `.gitignore`          | Defines what to not commit to Git.                                    | No        |
| `.markdownlint.json`  | Markdownlint configuration file.                                      | No        |
| `.prettierignore`     | Defines folders or files to not format in Prettier                    | No        |
| `.prettierrc`         | Prettier configuration file.                                          | No        |
| `CHANGELOG.md`        | Contains automatically generated GWS Changelogs.                      | No        |
| `CODE_OF_CONDUCT.md`  | This contains the contributor covenant code of conduct.               | No        |
| `LICENSE`             | The Gamma Web Server license.                                         | No        |
| `README.md`           | Before asking questions, read this file.                              | No        |
| `SECURITY.md`         | Contains the security policy of Gamma, along with supported versions. | No        |
| `gen_changelog.cmd`   | Command file for generating changelogs in CHANGELOG.md                | No        |

## ✨ Features

- Supports `HTTP` and `HTTPS`
- Easy & quick setup
- Easy to use configuration file

## 🚧 Build

Read our build documentation @ [/docs/build.md](/docs/build.md) to learn how to build Gamma Web Server.

## 🚀 Installation

Download the latest release from `GitHub`, or find the latest release on our website: [https://www.gammaws.gq](https://www.gammaws.gq)

## 📖 Usage

To start `Gamma Web Server`, double click on the `.exe` file, or run it through the `command-line`.

```cmd
.\gws.exe
```

To see your website, visit the link showed on the command-line below the product information. By default, the link is [localhost](localhost).

## ⚙️ Configuration

`Gamma Web Server` uses a `JSON configuration file` to customize its behavior. You can specify the port, the document root & more in the configuration file. You can find an example config.json in the src/json directory. One will be generated when the project is built.

Here is an example of what the **config.json** should look like.

```json
{
  "port": ":80",
  "tls_config": {
    "cert_file": "server.crt",
    "key_file": "server.key"
  },
  "static_dir": "public",
  "repo_config": {
    "version": "1.1.0",
    "author": "Official B",
    "product": "Gamma Web Server",
    "repository": "https://github.com/gamma-gws/gws"
  }
}
```

## 📄 License

Gamma Web Server is licensed under the GNU General Public License v3.0 - see the [`LICENSE`](LICENSE) file for details.

## 💪 Contributors

Thank you to the wonderful people below! ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://www.recon.best"><img src="https://avatars.githubusercontent.com/u/86677439?v=4?s=100" width="100px;" alt="recon"/><br /><sub><b>recon</b></sub></a><br /><a href="https://github.com/r1c0n/gws/commits?author=r1c0n" title="Code">💻</a> <a href="https://github.com/r1c0n/gws/commits?author=r1c0n" title="Documentation">📖</a> <a href="#design-r1c0n" title="Design">🎨</a> <a href="#example-r1c0n" title="Examples">💡</a> <a href="#infra-r1c0n" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/r1c0n/gws/pulls?q=is%3Apr+reviewed-by%3Ar1c0n" title="Reviewed Pull Requests">👀</a> <a href="#platform-r1c0n" title="Packaging/porting to new platform">📦</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/porokimun"><img src="https://avatars.githubusercontent.com/u/80103152?v=4?s=100" width="100px;" alt="porokimun"/><br /><sub><b>porokimun</b></sub></a><br /><a href="#design-porokimun" title="Design">🎨</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/kimotpe19"><img src="https://avatars.githubusercontent.com/u/80103152?v=4?s=100" width="100px;" alt="kimotpe"/><br /><sub><b>kimotpe</b></sub></a><br /><a href="#platform-kimotpe19" title="Packaging/porting to new platform">📦</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/zauceee"><img src="https://avatars.githubusercontent.com/u/37784801?v=4?s=100" width="100px;" alt="zauce"/><br /><sub><b>zauce</b></sub></a><br /><a href="#platform-zauceee" title="Packaging/porting to new platform">📦</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

---

<h3 align="center">
Don't forget to give a ⭐️ to <b>Gamma Web Server</b>! It's a great motivation booster.
</h3>
