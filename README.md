# Cron parser

Uses Go modules.

## Build

Run `make build` in the root directory to build the parser

## Test

Run `make test` in the root directory to test and return coverage

## Execution

With the parser build run with cron definition as a single argument.

Example execution

`./cron-parser "*/15, 0, 1,15, *, 1-5, /usr/bin/find"`