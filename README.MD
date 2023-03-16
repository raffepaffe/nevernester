# Nevernester

Nevernester is a Go linter that finds nesting in functions. Many levels of nesting
reduces readability and adds cognitive load.

## Why nevernester?

To many levels of nesting makes it hard to understand what the code does.
Here is a simple example with five levels of nesting being reduced to three:
```go
func calculate(order *Order) int {
    sum := 0
    if order != nil {
        if order.Header.IsValid {
            for _, row := range order.Rows {
                if row.IsValid {
                    sum = sum + row.Price
                }
            }
        }
    }
    return sum
}
```

Here the calculate function is easier to understand:
```go
func calculate(order *Order) int {
    sum := 0
    if order == nil {
        return sum
    }

    if order.Header.IsValid {
        sum = getPrice(order.Rows)
    }

    return sum
}

func getPrice(rows []OrderRow) int {
    sum := 0
    for _, row := range rows {
        if row.IsValid {
            sum = sum + row.Price
        }
    }
    return sum
}
```

## Installation & usage

```
$go install github.com/raffepaffe/nevernester@latest
$nevernester ./...
```

The default level of nesting is 4. If you want another level, say 5,  use the maxNesting option
```
$nevernester -maxNesting 5 ./...
```

If you want to skip checking tests or benchmarks use the skipTests and skipBenchmarks options
```
$nevernester -skipTests=true -skipBenchmarks=true ./...
```
### Still not convinced?
Linus Torvalds is pretty good developer and the 
[coding style rules](https://www.kernel.org/doc/Documentation/process/coding-style.rst) for the linux kernel states:
>The answer to that is that if you need more than 3 levels of indentation, you're screwed anyway, and should fix
your program.

Well, most of us are not as cool as Linus so nevernester has a default value of 4 levels. 
You can change that to whatever you want.

### Examples
<details>
  <summary>running nevernester on go 1.20.1 source with skipTests, skipBenchmarks and maxNesting set to 6</summary>

```
$nevernester -skipTests=true -skipBenchmarks=true -maxNesting 6 ./...
/usr/local/go/src/net/dnsclient_unix.go:595:1: calculated nesting for function goLookupIPCNAMEOrder is 7, max is 6
/usr/local/go/src/crypto/x509/parser.go:643:1: calculated nesting for function processExtensions is 7, max is 6
/usr/local/go/src/crypto/tls/handshake_messages.go:867:1: calculated nesting for function marshal is 7, max is 6
/usr/local/go/src/crypto/tls/handshake_messages.go:1085:1: calculated nesting for function marshal is 8, max is 6
/usr/local/go/src/crypto/tls/handshake_messages.go:1336:1: calculated nesting for function marshalCertificate is 9, max is 6
/usr/local/go/src/go/parser/parser.go:1055:1: calculated nesting for function parseMethodSpec is 7, max is 6
/usr/local/go/src/go/types/decl.go:836:1: calculated nesting for function declStmt is 7, max is 6
/usr/local/go/src/image/jpeg/scan.go:51:1: calculated nesting for function processSOS is 8, max is 6
/usr/local/go/src/image/png/writer.go:304:1: calculated nesting for function writeImage is 7, max is 6
/usr/local/go/src/database/sql/fakedb_test.go:881:1: calculated nesting for function QueryContext is 8, max is 6
/usr/local/go/src/go/parser/parser_test.go:360:1: calculated nesting for function getField is 9, max is 6
/usr/local/go/src/runtime/pprof/pprof_test.go:568:1: calculated nesting for function matchAndAvoidStacks is 7, max is 6
```
<br>
<br>
</details>

<details>
  <summary>running nevernester on golangci-lint 1.51.2 source with skipTests, skipBenchmarks and maxNesting set to 5</summary>

```
$nevernester -skipTests=true -skipBenchmarks=true -maxNesting 5 ./...
/Users/ralf/code/temp/golangci-lint/pkg/golinters/goanalysis/runners.go:126:1: calculated nesting for function saveIssuesToCache is 6, max is 5
/Users/ralf/code/temp/golangci-lint/pkg/golinters/deadcode.go:17:1: calculated nesting for function NewDeadcode is 6, max is 5
/Users/ralf/code/temp/golangci-lint/pkg/golinters/gocritic.go:169:1: calculated nesting for function runGocriticOnFile is 6, max is 5
/Users/ralf/code/temp/golangci-lint/pkg/golinters/gofmt_common.go:237:1: calculated nesting for function extractIssuesFromPatch is 6, max is 5
/Users/ralf/code/temp/golangci-lint/pkg/commands/version.go:35:1: calculated nesting for function initVersion is 6, max is 5
```
</details>

## Credits

The idea for this linter came after I watched [this video](https://www.youtube.com/watch?v=CFRhGnuXG-4) discussing 
nesting in code.  
Before I started to code I looked at other linters like [cyclop](https://github.com/bkielbasa/cyclop), 
[nilnil](https://github.com/Antonboom/nilnil) and [wrapcheck](https://github.com/tomarrell/wrapcheck).  
Credit also goes to (in alphabetic order) Anders, Björn, Hampus and Oscar for 
helping me becoming a better Go developer.

## Contributing

As with most tools, this will likely miss some cases. If you come across a case which you 
think should be covered and isn't, please file an issue including a minimum reproducible example of the case.

## License

This project is licensed under the MIT license. See the [LICENSE](./LICENSE) file for more
details.