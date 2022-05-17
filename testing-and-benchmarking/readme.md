# Remember to BET

- Benchmark
- Example
- Test

``` text

BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

## Commands

```bash
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```
