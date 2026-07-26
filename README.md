# EVE Trader

A CLI tool for analyzing EVE Online Jita market opportunities.

## Features

- Fetch Jita market orders from EVE ESI API
- Calculate buy/sell spread
- Calculate broker fees and sales tax
- Calculate real ROI after fees

## Usage

```bash
go run ./cmd/trader
```

Use calculator directly:

```bash
go build -o eve-trader ./cmd/trader
./eve-trader calculate 422 100 1
```

## Example

```text
Buy: 826300 ISK
Sell: 885900 ISK
ROI: -4.04%
```
