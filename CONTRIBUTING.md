# Contributing to Pulumi

First, thanks for contributing to Pulumi and helping make it better. We appreciate the help! If you're looking for an issue to start with, we've tagged some issues with the [help-wanted](https://github.com/pulumi/pulumi/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22) tag but feel free to pick up any issue that looks interesting to you or fix a bug you stumble across in the course of using Pulumi. No matter the size, we welcome all improvements.

For larger features, we'd appreciate it if you open a [new issue](https://github.com/pulumi/pulumi/issues/new) before doing a ton of work to discuss the feature before you start writing a lot of code.

## Hacking on Pulumi

To hack on Pulumi, you'll need to get a development environment set up. You'll want to install the following on your machine:

- Go 1.17
- NodeJS 12.X.X or later
- Python 3.6 or later
- [.NET Core](https://dotnet.microsoft.com/download)
- [pipenv](https://github.com/pypa/pipenv)
- [Golangci-lint](https://github.com/golangci/golangci-lint)
- [Yarn](https://yarnpkg.com/)
- [Pulumictl](https://github.com/pulumi/pulumictl)

## Configuring Windows machine

Follow these instructions to setup Pulumi development environment on a Windows-based machine. You need to install these tools:

- [Windows Subsystem (WSL) for Linux](https://docs.microsoft.com/windows/wsl/install-manual)
- [Docker Desktop WSL backend](https://docs.docker.com/desktop/windows/install/)
- [Visual Studio Code](https://code.visualstudio.com/) and [Remote - WSL](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) extension for VS Code

Once you've installed the required components, open `Microsoft Store` and search for `Linux`. You will see a list of distros to pick from including Ubuntu. Although you can work with any suitable Linux distribution, we assume that Ubuntu is picked. Launch Ubuntu as an app, which will open a command line window to its bash shell.

Run the following command to install brew as Homebrew:

```bash
sudo apt update
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Pay close attention to the output - you'll need to add run additional commands to add brew to your PATH as shown below:

```bash
echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /home/my_demo_username/.profile
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
```

FInally, you need to clone pulumi repository from within the WSL terminal and launch the VS Code as following:

```bash
git clone https://github.com/pulumi/pulumi.git
cd pulumi/
code .
```

Congratulations, the configuration of the Windows development environment is complete! You can now proceed with [Installing dependencies](#installing-dependencies), just use the VS Code bash terminal to run brew and the follow-up commands. 

## Installing dependencies

You can easily get all required dependencies with brew and npm

```bash
brew install node pipenv python@3 typescript yarn go@1.17 golangci/tap/golangci-lint pulumi/tap/pulumictl
curl https://raw.githubusercontent.com/Homebrew/homebrew-cask/0272f0d33f/Casks/dotnet-sdk.rb > dotnet-sdk.rb  # v3.1.0
brew install --HEAD -s dotnet-sdk.rb
rm dotnet-sdk.rb
```

## Hacking on Pulumi in Gitpod

If you have a web browser, you can get a fully pre-configured Pulumi development environment in one click:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/pulumi/pulumi)

## Make build system

We use `make` as our build system, so you'll want to install that as well, if you don't have it already. We build Pulumi in `$PULUMI_ROOT`, which defaults to `$HOME/.pulumi`. If you would like to build Pulumi in another location, you do so by setting `$PULUMI_ROOT`. 

```bash
export PATH=$HOME/.pulumi/bin:$PATH
```

You'll also need to make sure your maximum open file descriptor limit is set to 5000 at a minimum.

```bash
ulimit -n # to test
ulimit -n 5000
```

Across our projects, we try to use a regular set of make targets. The ones you'll care most about are:

0. `make ensure`, which restores/installs any build dependencies
1. `make`, which builds Pulumi and runs a quick set of tests
2. `make all` which builds Pulumi and runs the quick tests and a larger set of tests.

We make heavy use of integration level testing where we invoke `pulumi` to create and then delete cloud resources. This requires you to have a Pulumi account (so [sign up for free](https://pulumi.com) today if you haven't already) and login with `pulumi login`.

This repository does not actually create any real cloud resources as part of testing, but still uses Pulumi.com to store information abot some synthetic resources it creates during testing. Other repositories may require additional setup before running tests (most often this is just setting a few environment variables that tell the tests some information about how to use the cloud provider we are testing). Please see the `CONTRIBUTING.md` file in the repository, which will explain what additional configuration needs to be done before running tests.

## Debugging

The Pulumi tools have extensive logging built in.  In fact, we encourage liberal logging in new code, and adding new logging when debugging problems.  This helps to ensure future debugging endeavors benefit from your sleuthing.

All logging is done using a fork of Google's [Glog library](https://github.com/pulumi/glog).  It is relatively bare-bones, and adds basic leveled logging, stack dumping, and other capabilities beyond what Go's built-in logging routines offer.

The `pulumi` command line has two flags that control this logging and that can come in handy when debugging problems. The `--logtostderr` flag spews directly to stderr, rather than the default of logging to files in your temp directory. And the `--verbose=n` flag (`-v=n` for short) sets the logging level to `n`.  Anything greater than 3 is reserved for debug-level logging, greater than 5 is going to be quite verbose, and anything beyond 7 is extremely noisy.

For example, the command

```sh
$ pulumi preview --logtostderr -v=5
```

is a pretty standard starting point during debugging that will show a fairly comprehensive trace log of a compilation.

## Submitting a Pull Request

For contributors we use the standard fork based workflow. Fork this repository, create a topic branch, and start hacking away.  When you're ready, make sure you've run the tests (`make travis_pull_request` will run the exact flow we run in CI) and open your PR.
When adding a changelog entry, please be sure to use `CHANGELOG_PENDING.md` for the entry - we will then be able to ensure your PR gets into the next release.

## Getting Help

We're sure there are rough edges and we appreciate you helping out. If you want
to talk with other folks hacking on Pulumi (or members of the Pulumi team!)
come hang out `#contribute` channel in the
[Pulumi Community Slack](https://slack.pulumi.com/).
