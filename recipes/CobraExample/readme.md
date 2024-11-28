
```
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
~/go/bin/cobra-cli init
```

Adding a new (sub-)command:
```
~/go/bin/cobra-cli add random
```

And another one:

```
~/go/bin/cobra-cli add greeter
```

When a subcommand is added, you can define arguments specific to the subcommand there in the init. Remember to reference the new subcommand. So in the case with having added 'greeter':

```go
func init() {
	rootCmd.AddCommand(greeterCmd)

	greeterCmd.Flags().StringVarP(&Name, "name", "n", "", "the name")
	greeterCmd.MarkFlagRequired("name")
}
```