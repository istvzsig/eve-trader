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

Use calculator:

```bash
go build -o eve-trader ./cmd/trader
./eve-trader calculate 422 100 1
```

Example:

```text
Buy: 826300 ISK
Sell: 885900 ISK
ROI: -4.04%
```

Use margin-trader:

```bash
go build -o eve-trader ./cmd/trader
./eve-trader margin-trader 20
```

Example:

```text
[1] Makra's Modified Small Focused Pulse Laser (TypeID=85012) Buy=2.80B ISK Sell=3.91B ISK ROI=25.00% Verdict=✅ GOOD Net=700.24M ISK
[2] Medium 'Integrative' Hull Repair Unit (TypeID=21506) Buy=10.02M ISK Sell=13.99M ISK ROI=24.96% Verdict=✅ GOOD Net=2.50M ISK
[3] Corpum B-Type Thermal Energized Membrane (TypeID=18861) Buy=50.48M ISK Sell=70.47M ISK ROI=24.94% Verdict=✅ GOOD Net=12.59M ISK
```
