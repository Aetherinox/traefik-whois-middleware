<div align="center">
<h6>Traefik v3 middleware which displays user connection information about each client that accesses a container with this middleware applied.</h6>
<h2>♾️ Traefik Whois Middleware ♾️</h1>

<br />

<p>

(Under Development): This middleware allows you to view detailed information about a connecting client when they access a container where this middleware is enabled.

</p>

<br />

<br />
<br />

</div>

<div align="center">

<!-- prettier-ignore-start -->
[![Version][github-version-img]][github-version-uri]
[![Downloads][github-downloads-img]][github-downloads-uri]
[![Build Status][github-build-img]][github-build-uri]
[![Size][github-size-img]][github-size-img]
[![Last Commit][github-commit-img]][github-commit-img]
[![Contributors][contribs-all-img]](#contributors-)

[![Built with Material for MkDocs](https://img.shields.io/badge/Powered_by_Material_for_MkDocs-526CFE?style=for-the-badge&logo=MaterialForMkDocs&logoColor=white)](https://aetherinox.github.io/traefik-whois-middleware/)
<!-- prettier-ignore-end -->

</div>

<br />

---

<br />

- [Configuration](#configuration)
  - [Static File](#static-file)
    - [File (YAML)](#file-yaml)
    - [File (TOML)](#file-toml)
    - [CLI](#cli)
  - [Dynamic File](#dynamic-file)
    - [File (YAML)](#file-yaml-1)
    - [File (TOML)](#file-toml-1)
    - [Kubernetes Custom Resource Definition](#kubernetes-custom-resource-definition)
- [Parameters](#parameters)
  - [debugLogs](#debuglogs)
- [Full Examples](#full-examples)
- [Local Install](#local-install)
  - [Static File](#static-file-1)
    - [File (YAML)](#file-yaml-2)
    - [File (TOML)](#file-toml-2)
  - [Dynamic File](#dynamic-file-1)
    - [File (YAML)](#file-yaml-3)
    - [File (TOML)](#file-toml-3)
- [Contributors ✨](#contributors-)

<br />

---

<br />

## Configuration
The following provides examples for usage scenarios.

<br />

### Static File
If you are utilizing a Traefik **Static File**, review the following examples:

<br />

#### File (YAML)

```yaml
## Static configuration
experimental:
  plugins:
    traefik-whois-middleware:
      moduleName: "github.com/Aetherinox/traefik-whois-middleware"
      version: "v0.1.0"
```

<br />

#### File (TOML)

```toml
## Static configuration
[experimental.plugins.traefik-whois-middleware]
  moduleName = "github.com/Aetherinox/traefik-whois-middleware"
  version = "v0.1.0"
```

<br />

#### CLI

```bash
## Static configuration
--experimental.plugins.traefik-whois-middleware.modulename=github.com/Aetherinox/traefik-whois-middleware
--experimental.plugins.traefik-whois-middleware.version=v0.1.0
```

<br />

### Dynamic File
If you are utilizing a Traefik **Dynamic File**, review the following examples:

<br />

#### File (YAML)

```yaml
# Dynamic configuration
http:
  middlewares:
    whois:
      plugin:
        traefik-whois-middleware:
          debugLogs: false
```

<br />

#### File (TOML)

```toml
# Dynamic configuration
[http]
  [http.middlewares]
    [http.middlewares.whois]
      [http.middlewares.whois.plugin]
        [http.middlewares.whois.plugin.traefik-whois-middleware]
          debugLogs = false
```

<br />

#### Kubernetes Custom Resource Definition

```yaml
# Dynamic configuration
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: whois
spec:
  plugin:
    traefik-whois-middleware:
      debugLogs: false
```

<br />

---

<br />

## Parameters
This plugin accepts the following parameters:

<br />

| Parameter | Description | Default | Type | Required |
| --- | --- | --- | --- | --- |
| <sub>`debugLogs`</sub> | <sub>Shows debug logs in console</sub> | <sub>false</sub> | <sub>bool</sub> | <sub>⭕ Optional</sub> |

<br />
<br />

### debugLogs
If set `true`, gives you a more detailed outline of what is going on behind the scenes in your console. 

<br />

---

<br />

## Full Examples
A few extra examples have been provided.

<br />

```yml
http:
  middlewares:
    whois:
      plugin:
        traefik-whois-middleware:
          debugLogs: true

    routers:
        traefik-http:
            service: "traefik"
            rule: "Host(`yourdomain.com`)"
            entryPoints:
                - http
            middlewares:
                - https-redirect@file

        traefik-https:
            service: "traefik"
            rule: "Host(`yourdomain.com`)"
            entryPoints:
                - https
            middlewares:
                - whois@file
            tls:
                certResolver: cloudflare
                domains:
                    - main: "yourdomain.com"
                      sans:
                          - "*.yourdomain.com"
```

<br />

---

<br />

## Local Install
Traefik comes with the ability to install this plugin locally without fetching it from Github. 

<br />

Download a local copy of this plugin to your server within your Traefik installation folder.
```shell
git clone https://github.com/Aetherinox/traefik-whois-middleware.git
```

<br />

If you are running **Docker**, you need to mount a new volume:

<br />

> [!WARNING]
> The path to the plugin is **case sensitive**, do not change the casing of the folders, or the plugin will fail to load.

<br />

```yml
services:
    traefik:
        container_name: traefik
        image: traefik:latest
        restart: unless-stopped
        volumes:
            - ./traefik-whois-middleware:/plugins-local/src/github.com/Aetherinox/traefik-whois-middleware/
```

<br />

### Static File
Open your **Traefik Static File** and change `plugins` to `localPlugins`.

<br />

#### File (YAML)

```yaml
# Static configuration
experimental:
  localPlugins:
    traefik-whois-middleware:
      moduleName: "github.com/Aetherinox/traefik-whois-middleware"
      version: "v0.1.0"
```

<br />

#### File (TOML)

```toml
# Static configuration
[experimental.localPlugins.traefik-whois-middleware]
  moduleName = "github.com/Aetherinox/traefik-whois-middleware"
  version = "v0.1.0"
```

<br />

### Dynamic File
For local installation, your dynamic file will contain the same contents as it would if you installed the plugin normally.

<br />

#### File (YAML)

```yaml
# Dynamic configuration
http:
  middlewares:
    whois:
      plugin:
        traefik-whois-middleware:
          debugLogs: true
```

<br />

#### File (TOML)

```toml
# Dynamic configuration
[http]
  [http.middlewares]
    [http.middlewares.whois]
      [http.middlewares.whois.plugin]
        [http.middlewares.whois.plugin.traefik-whois-middleware]
          debugLogs = true
```


<br />

---

<br />

## Contributors ✨
We are always looking for contributors. If you feel that you can provide something useful to Gistr, then we'd love to review your suggestion. Before submitting your contribution, please review the following resources:

- [Pull Request Procedure](.github/PULL_REQUEST_TEMPLATE.md)
- [Contributor Policy](CONTRIBUTING.md)

<br />

Want to help but can't write code?
- Review [active questions by our community](https://github.com/Aetherinox/traefik-whois-middleware/labels/help%20wanted) and answer the ones you know.

<br />

![Alt](https://repobeats.axiom.co/api/embed/3a528b94b5433fa1f9763340435c6071716e7dca.svg "analytics image")

<br />

The following people have helped get this project going:

<br />

<div align="center">

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![Contributors][contribs-all-img]](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top"><a href="https://gitlab.com/Aetherinox"><img src="https://avatars.githubusercontent.com/u/118329232?v=4?s=40" width="80px;" alt="Aetherinox"/><br /><sub><b>Aetherinox</b></sub></a><br /><a href="https://github.com/Aetherinox/traefik-whois-middleware/commits?author=Aetherinox" title="Code">💻</a> <a href="#projectManagement-Aetherinox" title="Project Management">📆</a> <a href="#fundingFinding-Aetherinox" title="Funding Finding">🔍</a></td>
    </tr>
  </tbody>
</table>
</div>
<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

<br />
<br />

<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- BADGE > GENERAL -->
  [general-npmjs-uri]: https://npmjs.com
  [general-nodejs-uri]: https://nodejs.org
  [general-npmtrends-uri]: http://npmtrends.com/traefik-whois-middleware

<!-- BADGE > VERSION > GITHUB -->
  [github-version-img]: https://img.shields.io/github/v/tag/Aetherinox/traefik-whois-middleware?logo=GitHub&label=Version&color=ba5225
  [github-version-uri]: https://github.com/Aetherinox/traefik-whois-middleware/releases

<!-- BADGE > VERSION > NPMJS -->
  [npm-version-img]: https://img.shields.io/npm/v/traefik-whois-middleware?logo=npm&label=Version&color=ba5225
  [npm-version-uri]: https://npmjs.com/package/traefik-whois-middleware

<!-- BADGE > VERSION > PYPI -->
  [pypi-version-img]: https://img.shields.io/pypi/v/traefik-whois-middleware-plugin
  [pypi-version-uri]: https://pypi.org/project/traefik-whois-middleware-plugin/

<!-- BADGE > LICENSE > MIT -->
  [license-mit-img]: https://img.shields.io/badge/MIT-FFF?logo=creativecommons&logoColor=FFFFFF&label=License&color=9d29a0
  [license-mit-uri]: https://github.com/Aetherinox/traefik-whois-middleware/blob/main/LICENSE

<!-- BADGE > GITHUB > DOWNLOAD COUNT -->
  [github-downloads-img]: https://img.shields.io/github/downloads/Aetherinox/traefik-whois-middleware/total?logo=github&logoColor=FFFFFF&label=Downloads&color=376892
  [github-downloads-uri]: https://github.com/Aetherinox/traefik-whois-middleware/releases

<!-- BADGE > NPMJS > DOWNLOAD COUNT -->
  [npmjs-downloads-img]: https://img.shields.io/npm/dw/%40aetherinox%2Ftraefik-whois-middleware?logo=npm&&label=Downloads&color=376892
  [npmjs-downloads-uri]: https://npmjs.com/package/traefik-whois-middleware

<!-- BADGE > GITHUB > DOWNLOAD SIZE -->
  [github-size-img]: https://img.shields.io/github/repo-size/Aetherinox/traefik-whois-middleware?logo=github&label=Size&color=59702a
  [github-size-uri]: https://github.com/Aetherinox/traefik-whois-middleware/releases

<!-- BADGE > NPMJS > DOWNLOAD SIZE -->
  [npmjs-size-img]: https://img.shields.io/npm/unpacked-size/traefik-whois-middleware/latest?logo=npm&label=Size&color=59702a
  [npmjs-size-uri]: https://npmjs.com/package/traefik-whois-middleware

<!-- BADGE > CODECOV > COVERAGE -->
  [codecov-coverage-img]: https://img.shields.io/codecov/c/github/Aetherinox/traefik-whois-middleware?token=MPAVASGIOG&logo=codecov&logoColor=FFFFFF&label=Coverage&color=354b9e
  [codecov-coverage-uri]: https://codecov.io/github/Aetherinox/traefik-whois-middleware

<!-- BADGE > ALL CONTRIBUTORS -->
  [contribs-all-img]: https://img.shields.io/github/all-contributors/Aetherinox/traefik-whois-middleware?logo=contributorcovenant&color=de1f6f&label=contributors
  [contribs-all-uri]: https://github.com/all-contributors/all-contributors

<!-- BADGE > GITHUB > BUILD > NPM -->
  [github-build-img]: https://img.shields.io/github/actions/workflow/status/Aetherinox/traefik-whois-middleware/release.yml?logo=github&logoColor=FFFFFF&label=Build&color=%23278b30
  [github-build-uri]: https://github.com/Aetherinox/traefik-whois-middleware/actions/workflows/release.yml

<!-- BADGE > GITHUB > BUILD > Pypi -->
  [github-build-pypi-img]: https://img.shields.io/github/actions/workflow/status/Aetherinox/traefik-whois-middleware/release-pypi.yml?logo=github&logoColor=FFFFFF&label=Build&color=%23278b30
  [github-build-pypi-uri]: https://github.com/Aetherinox/traefik-whois-middleware/actions/workflows/pypi-release.yml

<!-- BADGE > GITHUB > TESTS -->
  [github-tests-img]: https://img.shields.io/github/actions/workflow/status/Aetherinox/traefik-whois-middleware/tests.yml?logo=github&label=Tests&color=2c6488
  [github-tests-uri]: https://github.com/Aetherinox/traefik-whois-middleware/actions/workflows/tests.yml

<!-- BADGE > GITHUB > COMMIT -->
  [github-commit-img]: https://img.shields.io/github/last-commit/Aetherinox/traefik-whois-middleware?logo=conventionalcommits&logoColor=FFFFFF&label=Last%20Commit&color=313131
  [github-commit-uri]: https://github.com/Aetherinox/traefik-whois-middleware/commits/main/

<!-- prettier-ignore-end -->
<!-- markdownlint-restore -->