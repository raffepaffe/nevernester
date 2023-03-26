# Nevernester

Nevernester is a Go linter that finds nesting in golang functions. Many levels of nesting
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
$nevernester -max-nesting 5 ./...
```

If you want to skip checking tests or benchmarks use the skipTests and skipBenchmarks options
```
$nevernester -skip-tests=true -skip-benchmarks=true ./...
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
  <summary>running nevernester on go 1.20.2 source with skip-tests, skip-benchmarks and max-nesting set to 6</summary>

```
$nevernester -skip-tests=true -skip-benchmarks=true -max-nesting 6 ./...
/usr/local/go/src/runtime/map.go:1152:1: calculated nesting for function evacuate is 7, max is 6
/usr/local/go/src/crypto/x509/parser.go:643:1: calculated nesting for function processExtensions is 7, max is 6
/usr/local/go/src/encoding/xml/xml.go:544:1: calculated nesting for function rawToken is 7, max is 6
/usr/local/go/src/regexp/syntax/parse.go:894:1: calculated nesting for function parse is 7, max is 6
/usr/local/go/src/go/parser/parser.go:1055:1: calculated nesting for function parseMethodSpec is 7, max is 6
/usr/local/go/src/image/draw/draw.go:956:1: calculated nesting for function drawPaletted is 7, max is 6
/usr/local/go/src/image/jpeg/scan.go:51:1: calculated nesting for function processSOS is 8, max is 6
/usr/local/go/src/database/sql/fakedb_test.go:881:1: calculated nesting for function QueryContext is 7, max is 6
/usr/local/go/src/go/parser/parser_test.go:360:1: calculated nesting for function getField is 9, max is 6
/usr/local/go/src/runtime/pprof/pprof_test.go:568:1: calculated nesting for function matchAndAvoidStacks is 7, max is 6
```
<br>
<br>
</details>

<details>
  <summary>running nevernester on golangci-lint 1.52.1 source with skip-tests, skip-benchmarks with default nesting</summary>

```
$nevernester -skip-tests=true -skip-benchmarks=true ./...
golangci-lint/pkg/golinters/goanalysis/errors.go:22:1: calculated nesting for function buildIssuesFromIllTypedError is 5, max is 4
golangci-lint/pkg/golinters/goanalysis/runners.go:126:1: calculated nesting for function saveIssuesToCache is 5, max is 4
golangci-lint/pkg/golinters/gofmt_common.go:237:1: calculated nesting for function extractIssuesFromPatch is 5, max is 4
golangci-lint/pkg/golinters/nakedret.go:109:1: calculated nesting for function Visit is 5, max is 4
```
</details>

<details>
  <summary>running nevernester on opentelemetry-go 1.15.0-rc.2 source with default settings</summary>

```
$nevernester ./...
opentelemetry-go/attribute/set_test.go:51:1: calculated nesting for function TestSetDedup is 5, max is 4
opentelemetry-go/exporters/otlp/internal/envconfig/envconfig_test.go:59:1: calculated nesting for function TestEnvConfig is 6, max is 4
```
</details>

## Credits

The idea for this linter came after I watched [this video](https://www.youtube.com/watch?v=CFRhGnuXG-4) discussing 
nesting in code.  

Before I started to code I looked at other linters like [cyclop](https://github.com/bkielbasa/cyclop), 
[nilnil](https://github.com/Antonboom/nilnil) and [wrapcheck](https://github.com/tomarrell/wrapcheck).  

Special thanks to Björn for helping out with testing.

Credit also goes to (in alphabetic order) Anders, Björn, Hampus and Oscar for 
helping me becoming a better Go developer.

## Contributing

As with most tools, this will likely miss some cases. If you come across a case which you 
think should be covered and isn't, please file an issue including a minimum reproducible example of the case.

## License

This project is licensed under the MIT license. See the [LICENSE](./LICENSE) file for more
details.