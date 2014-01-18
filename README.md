# Flint

Check your project for common sources of contributor friction.

### Usage

Run `flint` from your project root to check for some common ways to improve the
experience for potential contributors. Here's the output for a blank folder to
show the full gamut of suggestions:

```sh
~/projects/dream
❯ flint
[ERROR] README not found
[FIXME] Every project begins with a README. http://bit.ly/1dqUYQF
[ERROR] License file not found
[FIXME] Add a license to protect yourself and your users. http://choosealicense.com/
[WARNING] Contributing guide not found
[FIXME] Add a CONTRIBUTING guide for potential contributors. http://git.io/z-TiGg
[WARNING] Bootstrap script not found
[FIXME] A bootstrap script makes setup a snap. http://bit.ly/JZjVL6
[CRITICAL] Some critical problems found. Please fix right away!
```

You can also run this in older projects which were created by lazy you, or by
younger, less wise you.

Passing the `-h` flag will show full usage options:

```
❯ flint -h
NAME:
   flint - Check a project for common sources of contributor friction

USAGE:
   flint [global options] command [command options] [arguments...]

VERSION:
   0.0.2

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --skip-readme	skip check for README
   --skip-contributing	skip check for contributing guide
   --skip-license	skip check for license
   --skip-bootstrap	skip check for bootstrap script
   --skip-test		skip check for test script
   --skip-scripts	skip check for all scripts
   --version, -v	print the version
   --help, -h		show help
```

### Philosophy

If you want people to use and contribute to your project, you need to start by
answering their most basic questions. Flint is a command line script that will
check your project for common answers to these questions.

#### What is this?

Since it is so important, GitHub founder [Tom Preston-Werner][mojombo]
suggests you [should write the README before you write a single line of
code][RDD]. A well crafted README includes:

- A description of problems your project solves.
- The philosophy behind your project.
- Basic usage and getting started instructions.
- A list of comparable projects that inspired yours or would be suitable
  alternatives.

#### How am I allowed to use it?

Providing the source to your project isn't enough. While you don't _have to_
provide a license, doing so will make it clear to users and potential
contributors how they can legally use your software and what happens to
contributions they make. [Choose A License][choose] can help you pick the right
license for your project.

#### How do I contribute?

You'll want to tell folks about your development workflow so they'll know how
to submit patches for bugfixes and new features. When you add [CONTRIBUTING
guidelines][contributing] to your project, GitHub will make those available at
the top of every new Pull Request screen.

#### How do I get up and running in development?

A bootstrap script is a thoughtful way to let new users (and future versions of
yourself on new hardware) get up and running quickly. A good bootstrap script
detects and installs all project dependencies. Don't make your less technical
users learn devops. Make it as easy as running `script/bootstrap`.

#### How do I make sure my new features didn't break old functionality?

Good software projects have test suites that ensure the code works as
advertised. Even within language communities, there can be a myriad of test
frameworks. You can make it easy to run the test suite with a platform agnostic
`script/test` executable.

### Maintainers

[@pengwynn][pengwynn]

Copyright 2014 [Wynn Netherland][pengwynn].

[pengwynn]: https://github.com/pengwynn
[mojombo]: https://github.com/mojombo
[contributing]: https://github.com/blog/1184-contributing-guidelines
[octokit contrib]: https://github.com/octokit/octokit.rb/blob/master/CONTRIBUTING.md
[choose]: http://choosealicense.com/
[RDD]: http://tom.preston-werner.com/2010/08/23/readme-driven-development.html
