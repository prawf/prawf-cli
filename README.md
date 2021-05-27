<div align="center">
    <h1>prawf-cli</h1> 
    <p>🧪 Easy to use HTTP API testing framework built into an elegant CLI</p>
    <a href="https://github.com/prawf/prawf-cli/relea" target="_blank">
        <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/prawf/prawf-cli">
    </a>
    <a href="https://github.com/prawf/prawf-cli/commits/master" target="_blank">
        <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/prawf/prawf-cli">
    </a>
    <a href="https://github.com/prawf/prawf-cli/issues" target="_blank">
        <img alt="GitHub issues" src="https://img.shields.io/github/issues/prawf/prawf-cli">
    </a>
    <a href="https://github.com/prawf/prawf-cli/blob/master/LICENSE" target="_blank">
        <img alt="GitHub" src="https://img.shields.io/github/license/prawf/prawf-cli">
    </a>
    <a href="https://ctt.ac/MfgmK" target="_blank">
        <img alt="Twitter URL" src="https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fctt.ac%2FMfgmK">
    </a>
</div>

<p align="center">
    <sub>
        Made with ❤︎ by
        <a href="https://github.com/navendu-pottekkat">Navendu Pottekkat</a>
    </sub>
</p>

<div align="center">
    <a href="https://github.com/prawf/prawf-cli" target="_blank">
        <img alt="screenshot" src="https://raw.githubusercontent.com/prawf/prawf-cli/master/screenshot.png">
    </a>
</div>

# Why use prawf?

🏋️‍♂️ Lightweight- Does not add any overhead to your software

🧰 Cross platform- Compiled to a binary and works on Windows, Mac and Linux

📝 Declarative tests- Forget all those flags you use to send a request and write your tests in a file

🧱 Structured logs- Get structured logs so you do not have to spend hours debugging

🚰 Built-in CI/CD support- Ship your applications bug free by adding to your CI/CD pipelines

🔓 Free and open-source- It is and always will be

# Quick Start

* Install prawf- See the [Installation](#installation) guide.

* Open up your project folder. If you do have a project yet and is just testing prawf, you can create an empty folder.

* Run `prawf init` to create a `prawf.json` configuration file and initialise it.

> By default, prawf will initialise the prawf.json configuration file with the API endpoints from jsonplaceholder.typicode.com. You can use this for testing out the capabilities of prawf.

* Edit the `prawf.json` configuration file if you are testing a custom application. You can leave it as it is if you are just testing out prawf.

* Run `prawf run` to send requests to the endponits specified in your configuration file.

* Run `prawf test` to test the endpoints with the expected responses.

# Installation

Create a new directory to download prawf-

```shell
mkdir prawf
```

Download prawf- Go to the [releases page](https://github.com/prawf/prawf-cli/releases) to view available downloads.

Replace the URL below with the link to the `tar.gz` file for your particular operating system and run the command to download prawf-

```shell
curl -OL https://github.com/prawf/prawf-cli/releases/latest/download/prawf_0.0.1-alpha_Linux_x86_64.tar.gz
```

Extract the package. Replace the file name with the filename of your download-

```
tar -xvzf prawf_0.0.1-alpha_Linux_x86_64.tar.gz
```

Add prawf to your path(Linux, macOS)-

```
export PATH=$PWD:$PATH
```

Check your installation-

```
prawf version
```

