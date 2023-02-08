package golang

const GolangCiYml = `run:
  timeout: 3m

linters-settings:
  cyclop:
    max-complexity: 20
    package-average: 10.0

  errcheck:
    check-type-assertions: true

  exhaustive:
    check:
      - switch
      - map

  funlen:
    lines: 100
    statements: 50

  gocognit:
    min-complexity: 20

  gocritic:
    settings:
      captLocal:
        paramsOnly: false
      underef:
        skipRecvDeref: false

  gomodguard:
    blocked:
      modules:
        - github.com/golang/protobuf:
            recommendations:
              - google.golang.org/protobuf
            reason: "see https://developers.google.com/protocol-buffers/docs/reference/go/faq#modules"
        - github.com/satori/go.uuid:
            recommendations:
              - github.com/google/uuid
            reason: "satori's package is not maintained"
        - github.com/gofrs/uuid:
            recommendations:
              - github.com/google/uuid
            reason: "see recommendation from dev-infra team: https://confluence.gtforge.com/x/gQI6Aw"

  govet:
    # Enable all analyzers.
    # Default: false
    enable-all: true
    # Disable analyzers by name.
    # Run 'go tool vet help' to see all analyzers.
    # Default: []
    disable:
      - fieldalignment # too strict
    # Settings per analyzer.
    # settings:
    #   shadow:
    #     # Whether to be strict about shadowing; can be noisy.
    #     # Default: false
    #     strict: true

  nakedret:
    # Make an issue if func has more lines of code than this setting, and it has naked returns.
    # Default: 30
    max-func-lines: 0

  nolintlint:
    # Exclude following linters from requiring an explanation.
    # Default: []
    allow-no-explanation: [funlen, gocognit, lll]
    # Enable to require an explanation of nonzero length after each nolint directive.
    # Default: false
    require-explanation: true
    # Enable to require nolint directives to mention the specific linter being suppressed.
    # Default: false
    require-specific: true

  rowserrcheck:
    # database/sql is always checked
    # Default: []
    packages:
      - github.com/jmoiron/sqlx

  tenv:
    # The option 'all' will run against whole test files ('_test.go') regardless of method/function signatures.
    # Otherwise, only methods that take '*testing.T', '*testing.B', and 'testing.TB' as arguments are checked.
    # Default: false
    all: true

linters:
  disable-all: true
  enable:
    ## enabled by default
    - errcheck # checking for unchecked errors, these unchecked errors can be critical bugs in some cases
    - gosimple # specializes in simplifying a code
    # - govet # reports suspicious constructs, such as Printf calls whose arguments do not align with the format string
    - ineffassign # detects when assignments to existing variables are not used
    - staticcheck # is a go vet on steroids, applying a ton of static analysis checks
    - typecheck # like the front-end of a Go compiler, parses and type-checks Go code
    - unused # checks for unused constants, variables, functions and types
    ## disabled by default
    - asasalint # checks for pass []any as any in variadic func(...any)
    - asciicheck # checks that your code does not contain non-ASCII identifiers
    - bidichk # checks for dangerous unicode character sequences
    - bodyclose # checks whether HTTP response body is closed successfully
    - cyclop # checks function and package cyclomatic complexity
    - dupl # tool for code clone detection
    - durationcheck # checks for two durations multiplied together
    - errname # checks that sentinel errors are prefixed with the Err and error types are suffixed with the Error
    - errorlint # finds code that will cause problems with the error wrapping scheme introduced in Go 1.13
    - execinquery # checks query string in Query function which reads your Go src files and warning it finds
    - exhaustive # checks exhaustiveness of enum switch statements
    - exportloopref # checks for pointers to enclosing loop variables
    - forbidigo # forbids identifiers
    - funlen # tool for detection of long functions
    # - gochecknoglobals # checks that no global variables exist
    # - gochecknoinits # checks that no init functions are present in Go code
    # - gocognit # computes and checks the cognitive complexity of functions
    - goconst # finds repeated strings that could be replaced by a constant
    - gocritic # provides diagnostics that check for bugs, performance and style issues
    - gocyclo # computes and checks the cyclomatic complexity of functions
    - godot # checks if comments end in a period
    - goimports # in addition to fixing imports, goimports also formats your code in the same style as gofmt
    # - gomnd # detects magic numbers
    - gomoddirectives # manages the use of 'replace', 'retract', and 'excludes' directives in go.mod
    - gomodguard # allow and block lists linter for direct Go module dependencies. This is different from depguard where there are different block types for example version constraints and module recommendations
    - goprintffuncname # checks that printf-like functions are named with f at the end
    - gosec # inspects source code for security problems
    - lll # reports long lines
    - loggercheck # checks key value pairs for common logger libraries (kitlog,klog,logr,zap)
    - makezero # finds slice declarations with non-zero initial length
    - nakedret # finds naked returns in functions greater than a specified function length
    - nestif # reports deeply nested if statements
    - nilerr # finds the code that returns nil even if it checks that the error is not nil
    - nilnil # checks that there is no simultaneous return of nil error and an invalid value
    - noctx # finds sending http request without context.Context
    - nolintlint # reports ill-formed or insufficient nolint directives
    - nonamedreturns # reports all named returns
    - nosprintfhostport # checks for misuse of Sprintf to construct a host with port in a URL
    - predeclared # finds code that shadows one of Go's predeclared identifiers
    - promlinter # checks Prometheus metrics naming via promlint
    - reassign # checks that package variables are not reassigned
    - rowserrcheck # checks whether Err of rows is checked successfully
    - sqlclosecheck # checks that sql.Rows and sql.Stmt are closed
    - tenv # detects using os.Setenv instead of t.Setenv since Go1.17
    - testableexamples # checks if examples are testable (have an expected output)
    - tparallel # detects inappropriate usage of t.Parallel() method in your Go test codes
    - unconvert # removes unnecessary type conversions
    - unparam # reports unused function parameters
    - usestdlibvars # detects the possibility to use variables/constants from the Go standard library
    - wastedassign # finds wasted assignment statements
    - whitespace # detects leading and trailing whitespace
    # - revive # fast, configurable, extensible, flexible, and beautiful linter for Go, drop-in replacement of golint
    # - stylecheck # is a replacement for golint
    # - testpackage # makes you use a separate _test package

    ## you may want to enable
    #- decorder # checks declaration order and count of types, constants, variables and functions
    #- exhaustruct # checks if all structure fields are initialized
    #- gci # controls golang package import order and makes it always deterministic
    #- godox # detects FIXME, TODO and other comment keywords
    #- goheader # checks is file header matches to pattern
    #- interfacebloat # checks the number of methods inside an interface
    #- ireturn # accept interfaces, return concrete types
    #- prealloc # [premature optimization, but can be used in some cases] finds slice declarations that could potentially be preallocated
    #- varnamelen # [great idea, but too many false positives] checks that the length of a variable's name matches its scope
    #- wrapcheck # checks that errors returned from external packages are wrapped

    ## disabled
    #- containedctx # detects struct contained context.Context field
    #- contextcheck # [too many false positives] checks the function whether use a non-inherited context
    #- depguard # [replaced by gomodguard] checks if package imports are in a list of acceptable packages
    #- dogsled # checks assignments with too many blank identifiers (e.g. x, _, _, _, := f())
    #- dupword # [useless without config] checks for duplicate words in the source code
    #- errchkjson # [don't see profit + I'm against of omitting errors like in the first example https://github.com/breml/errchkjson] checks types passed to the json encoding functions. Reports unsupported types and optionally reports occasions, where the check for the returned error can be omitted
    #- forcetypeassert # [replaced by errcheck] finds forced type assertions
    #- goerr113 # [too strict] checks the errors handling expressions
    #- gofmt # [replaced by goimports] checks whether code was gofmt-ed
    #- gofumpt # [replaced by goimports, gofumports is not available yet] checks whether code was gofumpt-ed
    #- grouper # analyzes expression groups
    #- importas # enforces consistent import aliases
    #- maintidx # measures the maintainability index of each function
    #- misspell # [useless] finds commonly misspelled English words in comments
    #- nlreturn # [too strict and mostly code is not more readable] checks for a new line before return and branch statements to increase code clarity
    #- paralleltest # [too many false positives] detects missing usage of t.Parallel() method in your Go test
    #- tagliatelle # checks the struct tags
    #- thelper # detects golang test helpers without t.Helper() call and checks the consistency of test helpers
    #- wsl # [too strict and mostly code is not more readable] whitespace linter forces you to use empty lines

    ## deprecated
    #- deadcode # [deprecated, replaced by unused] finds unused code
    #- exhaustivestruct # [deprecated, replaced by exhaustruct] checks if all struct's fields are initialized
    #- golint # [deprecated, replaced by revive] golint differs from gofmt. Gofmt reformats Go source code, whereas golint prints out style mistakes
    #- ifshort # [deprecated] checks that your code uses short syntax for if-statements whenever possible
    #- interfacer # [deprecated] suggests narrower interface types
    #- maligned # [deprecated, replaced by govet fieldalignment] detects Go structs that would take less memory if their fields were sorted
    #- nosnakecase # [deprecated, replaced by revive var-naming] detects snake case of variable naming and function name
    #- scopelint # [deprecated, replaced by exportloopref] checks for unpinned variables in go programs
    #- structcheck # [deprecated, replaced by unused] finds unused struct fields
    #- varcheck # [deprecated, replaced by unused] finds unused global variables and constants

issues:
  # Maximum count of issues with the same text.
  # Set to 0 to disable.
  # Default: 3
  max-same-issues: 50

  exclude-rules:
    - path: postgres/subnodes_queries.go
      linters: [lll, gomnd, govet]
    - path: postgres/subnodes_queries_test.go
      linters: [lll, gomnd, govet, gocognit, gocyclo, cyclop]
    - source: "^//\\s*go:generate\\s"
      linters: [lll]
    - source: "(noinspection|TODO)"
      linters: [godot]
    - source: "//noinspection"
      linters: [gocritic]
    - source: "^\\s+if _, ok := err\\.\\([^.]+\\.InternalError\\); ok {"
      linters: [errorlint]
    - path: "_test\\.go"
      linters:
        - bodyclose
        - dupl
        - funlen
        - goconst
        - gosec
        - noctx
        - wrapcheck
        - errcheck
`
