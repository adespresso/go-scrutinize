# go-scrutinize
[![Build Status](https://scrutinizer-ci.com/g/phayes/go-scrutinize/badges/build.png?b=master)](https://scrutinizer-ci.com/g/phayes/go-scrutinize/build-status/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/phayes/go-scrutinize)](https://goreportcard.com/report/github.com/phayes/go-scrutinize)

Scrutinizer-CI support for golang.

This package generates code-coverage, does static analysis, and runs tests for [Scrutinizer-CI](https://scrutinizer-ci.com).  

## Getting Started

Copy and paste the yml below into your `.scrutinizer.yml` file to get started.

```yml
build:
    dependencies:
        before:
            - 'source <(curl -fsSL https://raw.githubusercontent.com/phayes/go-scrutinize/master/install-golang)'

    tests:
        override:
            -
                command: 'cd $PROJECTPATH && go-scrutinize'
                coverage:
                    file: 'coverage.xml'
                    format: 'clover'
                analysis:
                    file: 'checkstyle_report.xml'
                    format: 'general-checkstyle'
```

## Example Scrutinizer Report

This is an example of what the Scrutinizer Report page looks like for a golang project.

![Example Scrutinizer Report](http://i.imgur.com/1iBxgLb.png)
