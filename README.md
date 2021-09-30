<div align="center">
    <h1>prawf-cli</h1> 
    <p>🧪 Easy to use API testing framework built into an elegant CLI</p>
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

# What is prawf?

🧪 prawf is an API testing platform that is easy to use and does not add any overhead to your project.

You can use prawf to define tests for your API endpoints and use that definitions to run tests.

Try the [Quick Start](#quick-start) guide to run prawf without any configurations.

See the [docs](https://prawf.github.io) for additional usage information.

# Why use prawf?

🏋️‍♂️ Lightweight- Does not add any overhead to your software

🧰 Cross platform- Compiled to a binary and works on Windows, Mac and Linux

📝 Declarative tests- Forget all those flags you use to send a request and write your tests in a file

🧱 Structured logs- Get structured logs so you do not have to spend hours debugging

🚰 Built-in CI/CD support- Ship your applications bug free by adding to your CI/CD pipelines

🔓 Free and open-source- It is and always will be

# Quick Start

Install prawf- See the [Installation](#installation) guide.

Open up your project folder. If you do have a project yet and is just testing prawf, you can create an empty folder.

Create a new `prawf.json` config file-

```
prawf init
```

```json
prawf.json file not found. Would you like to create one? [y/n]? y
INFO[0005] File created.                                 file=prawf.json
INFO[0005] File loaded.                                  file=prawf.json
```

> By default, prawf will initialise the prawf.json configuration file with the API endpoints from [jsonplaceholder.typicode.com](https://jsonplaceholder.typicode.com). You can use this for testing out the capabilities of prawf.

Edit the `prawf.json` configuration file if you are testing a custom application. You can leave it as it is if you are just testing out prawf. See [Configuring prawf.json](#configuring-prawf.json) for more details.

Send requests to endpoints specified in the `prawf.json` file-

```
prawf run
```

```json
INFO[0000] File loaded.                                  file=prawf.json
INFO[0000] Running test.                                 test=sample-test url="https://jsonplaceholder.typicode.com"

INFO[0000] Creating request.                             method=GET name=get-pass path=/posts
INFO[0001] Response received.                            status code="200 OK"
INFO[0001] [
  {
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  }
] 
```

Test the endpoints as specified in the `prawf.json` file-

```
prawf test
```

```
INFO[0000] File loaded.                                  file=prawf.json
INFO[0000] Running test.                                 test=sample-test url="https://jsonplaceholder.typicode.com"

INFO[0000] Creating request.                             method=GET name=get-pass path=/posts
INFO[0000] Response received.                            status code="200 OK"
INFO[0000] [
  {
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  }
] 

INFO[0000] Expected response.                            contain=yes equal=no keys=no test=pass
INFO[0000] {
  "id": 1
  }                               type=contain
```

Check out the [docs](https://prawf.github.io/) for more examples and info on using prawf.

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

See the [Quick Start](#quick-start) guide or the [docs](https://prawf.github.io/) to getting started.

# Configuring prawf.json

The `prawf.json` config file will contain all your tests. You can define tests by specifying the endpoints request details and provide the expected responses to test with.

To start, open the `prawf.json` file in your project folder. If you do not have a `prawf.json` file, run-

```
prawf init
```

See the [Quick Start](#quick-start) guide for more details on getting started.

Once you have the `prawf.json` file initialised, open it on a text editor.

## Available Configurations

The following could be configured in the `prawf.json` file. See the [example below](#example-prawf.json-file)-

```
`current`- Represents the focused test. Users can define multiple tests and the test mentioned in `current` will be used for running tests.

`tests` - Array of tests. Users can define multiple tests here with a test name as mentioned below.

    `test-name`- Name of the test. Its value contains the test.

        `url`- The URL the current test is interested in.

        `methods`- Array of methods to test the URL on.
            
            `name`- Name of the method.
            
            `path`- The path which will be added to the URL.
            
            `method`- Type of request.
            
            `query`- The query to add to the request.
            
            `header`- The header to add to the request.
            
            `body`- The body to add to the request.
            
            `expect`- Represents what to look for in the response for testing.
            
                `keys`- To check if the keys mentioned are present in the response.
                
                `contain`- To check if the key-value pairs are present in the response.
                
                `equal`- To check if the response if exactly equal to the present value.
```

### Example prawf.json file

```json
{
  "current": "sample-test",
  "tests": {
    "sample-test": {
      "url": "https://jsonplaceholder.typicode.com",
      "methods": [
        {
          "name": "get-pass",
          "path": "/posts",
          "method": "get",
          "query": {
            "id": 1
          },
          "expect": {
            "contain": {
              "id": 1
            }
          }
        },
        {
          "name": "get-fail",
          "path": "/posts",
          "method": "get",
          "query": {
            "id": 3
          },
          "expect": {
            "keys": [
              "uuid"
            ]
          }
        },
        {
          "name": "post-pass",
          "path": "/posts",
          "method": "post",
          "header": {
            "Content-type": "application/json; charset=UTF-8"
          },
          "body": {
            "body": "If you haven't already, check out prawf to test your REST API endpoints",
            "title": "prawf is amazing!",
            "userId": 1
          },
          "expect": {
            "equal": {
              "body": "If you haven't already, check out prawf to test your REST API endpoints",
              "title": "prawf is amazing!",
              "userId": 1
            }
          }
        },
        {
          "name": "put-pass",
          "path": "/posts/1",
          "method": "put",
          "header": {
            "Content-type": "application/json; charset=UTF-8"
          },
          "body": {
            "body": "Give us a star on GitHub/prawf!",
            "id": 1,
            "title": "prawf is awesome!",
            "userId": 1
          },
          "expect": {
            "contain": {
              "title": "prawf is awesome!"
            }
          }
        }
      ]
    }
  }
}
```

# Sending requests

Using prawf, you can send requests to your API endpoints and view the responses.

You can configure your endpoint URLs, paths and methods on your `prawf.json file`. ([See Available Configurations](#available-configurations))

Once you have finished adding your configuration and saved the file, you can run-

`prawf run`

to start sending requests.

# Running tests

You can also automate the testing process by declaring it in the `prawf.json` file. ([See Available Configurations](#available-configurations))

These tests will be run automatically and structured logs will printed on the screen.

Once the configuration is done, run-

`prawf test`

to start running the test.

# Contributing

All contributions are welcome! Please check the [Contributing Guidelines](https://github.com/prawf/prawf-cli/blob/master/CONTRIBUTING.md) to get started.

Join the [discussions](https://github.com/prawf/prawf-cli/discussions) and let us know what you think!

# License

[GNU General Public License v3.0](https://github.com/prawf/prawf-cli/blob/master/LICENSE)
