# bitcal

```go
import "github.com/cxtuttle/go-bitcal"
```

[API documentation](doc.md)

bitcal provides 2 functions,  one to add work days to a date and another to compute the delta work days between 2 dates. Inspiration was taken from Date::Calendar.

## Examples
```go
	bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

	start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)

	new_date := bc.AddWorkDays(start_date, -55)
```

```go
	bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

	start_date := time.Date(2019, 8, 4, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC)

	delta := bc.DeltaWorkDays(start_date, end_date)
```