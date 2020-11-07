# GoEnv

### Description
Go environment variables helper functions

### Usage

##### Simple strings
```go
foo := goenv.String("SOME_ENV_VAR")
// or
foo := goenv.StringOrDefault("SOME_ENV_VAR", "defaultValue")
// or
foo := MustString("SOME_ENV_VAR") // will panic if not found
```

##### Integers
```go
var foo int = goenv.Int("SOME_ENV_VAR")
// or
defaultValue := 5
foo := goenv.IntOrDefault("SOME_ENV_VAR", defaultValue)
// or
foo := MustInt("SOME_ENV_VAR") // will panic if not found
```

##### Booleans
```go
var foo bool = goenv.IsTruthy("SOME_ENV_VAR")
// or
foo := goenv.IsFalthy("SOME_ENV_VAR")
```