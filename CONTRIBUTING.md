# Contributing to prawf

Thank you for your interest in contributing to prawf!

📃 🐛 💡 ⚙️ Contributions of any kind are welcome here! It could be changes to documentation, reporting bugs, requesting for features and also contributing code.

Before you start, please go through the [Code of Conduct](CODE_OF_CONDUCT.md).

All changes are made through [pull requests](https://github.com/prawf/prawf-cli/pulls). Please check the [GithHub Flow](https://guides.github.com/introduction/flow/index.html)

We use [discussions](https://github.com/prawf/prawf-cli/discussions) for communicating. Please use it to share your ideas.

To report bugs and track features, we use issues. Please use the appropriate issue template while creating a new issue.

The docs are built in a [separate repository](https://github.com/prawf/prawf.github.io). Any changes to the docs should be made there.

## Setting up the development environment

prawf-cli is written in [Golang](https://golang.org/). 

To run the development environment, you should install the latest version of Go. See their [downloads page](https://golang.org/dl/).

Fork this repo and clone your fork to your local machine.

```
git clone https://github.com/<your-user-name>/prawf-cli.git
```

Or

```
gh repo clone <your-user-name>/prawf-cli
```

Inside the repo run-

```
make
```

to build prawf-cli locally.

You can then run-

```
./prawf <name of the command>
```

to use the local binary.
