# 开发Golang的CommandLine应用

Cobra is built on a structure of commands, arguments & flags.

***Commands*** represent actions, ***Args*** are things and ***Flags*** are modifiers for those actions.

```shell
go get -u github.com/spf13/cobra/cobra
```

## Generate Commands

```shell
cobra init --pkg-name qapi
```