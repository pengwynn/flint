# Flint

Check your project for common sources of contributor friction.

### Usage


```sh
❯ ./flint -h
flint [OPTION...]

Checks a project for common sources of contributor friction

-r, --readme                                                  Check for existence of README
-c, --contributing                                            Check for contributing guidelines
-l, --license                                                 Check for a project license
-b, --bootstrap                                               Check for a bootstrap script
-t, --test                                                    Check for a test script
--no-readme                                                   Skip README check
--no-contributing                                             Skip contributing guide check
--no-license                                                  Skip license check
--no-bootstrap                                                Skip bootstrap script check
--no-test                                                     Skip test script check
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
