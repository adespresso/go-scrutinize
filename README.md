# go-scrutinize

Scrutinizer-CI support for golang.

This package generates code-coverage, does static analysis, and runs tests for Scrutinizer-CI (scrutinizer-ci.com).  

Example .scrutinizer.yml file:

```yml
build:
    dependencies:
        before:
            - '"$(curl -fsSL https://raw.githubusercontent.com/phayes/go-scrutinize/master/install-golang)" | source /dev/stdin'

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
