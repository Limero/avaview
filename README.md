# AvaView

A tool to display stocks and funds from [Avanza](https://avanza.se) in the terminal.

## Get started

Install AvaView (it will install to `~/go/bin/avaview` unless changed)

```
go get -u github.com/limero/avaview
```

Login to Avanza and then export the two required .csv-files:

* [konto.csv](https://www.avanza.se/_cqbe/ff/gdpr/export/accounts)
* [positioner.csv](https://www.avanza.se/_cqbe/ff/gdpr/export/positions)

Put the .csv-files in the same directory as the executable and run it!

## License

Licensed under the MIT license. (See LICENSE)
