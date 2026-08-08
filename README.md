# EVE Trader

A CLI assistant for analyzing EVE Online Jita market opportunities and improving Alpha account trading decisions.

## Current Gameplay Progress

### Alpha Trading Status

Current goal:

> Build a sustainable Alpha-only Jita trading operation using market analysis instead of random flipping.

Current focus:

- Station trading in Jita IV - Moon 4 - Caldari Navy Assembly Plant
- Buying from player buy orders
- Selling into player sell orders
- Maximizing profit with limited Alpha skills and order slots

---

## Current Account Situation

Example:

```
Wallet:
36M ISK

Open Orders:
10 / 17

Remaining Slots:
7

Total Buy Order Escrow:
43.3M ISK

Sell Orders:
105.5M ISK
```

Current limitation:

- Capital is the main bottleneck
- Alpha order slots limit diversification
- Expensive items are usually not practical despite high ROI

---

## Trading Strategy

Current strategy:

### Margin Trading

Find items where:

```
Sell Price - Buy Price
```

creates a profitable spread after:

- Broker fees
- Sales tax
- Limited Alpha skills

Example:

```
Buy:
3.79M ISK

Sell:
4.17M ISK

Gross Profit:
380K ISK

After Fees:
~250K ISK
```

---

## Current Trading Tools

### Market Check

Check individual items:

```bash
./eve-trader margin-item ITEM_NAME
```

Example:

```bash
./eve-trader margin-item "Ballistic Control System II"
```

---

### Profit Calculator

Test a trade:

```bash
./eve-trader calculate SELL BUY QUANTITY
```

Example:

```bash
./eve-trader calculate 4.2m 400k 1
```

---

### Margin Scanner

Search Jita for opportunities:

```bash
./eve-trader margin-trade 20%
```

Example:

```
[1] Medium 'Integrative' Hull Repair Unit

Buy:
10.02M ISK

Sell:
13.99M ISK

ROI:
24.96%

Profit:
2.50M ISK
```

---

### ISK Challenge

Test a trade:

```bash
./eve-trader isk-challenge TARGET CURRENT
```

Description:

I saves the current challange to a `.json` file.

Example:

```bash
./eve-trader isk-challenge 500m 65m
```

To update the current challange:

```bash
./eve-trader isk-challenge add 420m
```

You can reset it by command below:

```bash
./eve-trader isk-challenge reset
```

---

# Gameplay Roadmap

## Phase 1 - Starting Capital ✅

Goal:

```
0 ISK → 50M ISK
```

Focus:

- Cheap modules
- Ammo
- Popular PvE/PvP items
- Fast turnover

Avoid:

- Billion ISK collector items
- SKINs
- Rare modules with 1 volume

---

## Phase 2 - Alpha Station Trader 🟡

Goal:

```
50M ISK → 500M ISK
```

Focus:

- 17 order slots
- High liquidity items
- Multiple small positions

Preferred trades:

```
Buy:
1M-20M ISK

ROI:
10-30%

Daily volume:
high
```

---

## Phase 3 - Scale Capital 🔜

Goal:

```
500M ISK → several billion ISK
```

Improve:

- More order slots
- Better trading skills
- Larger inventory
- More expensive items

---

# Current Weaknesses

The scanner can find:

```
20B ISK item
30% ROI
```

but Alpha cannot realistically trade it.

Future filters needed:

- Wallet limit
- Maximum buy cost
- Minimum market volume
- Liquidity score
- Expected turnover time

---

# Long Term Goal

Create an Alpha-friendly EVE trading assistant:

```
ESI API
   |
Market Scanner
   |
Profit Calculator
   |
Risk Filter
   |
Trade Recommendation
```

The goal is not finding the biggest ROI.

The goal is finding:

```
Affordable
+
Liquid
+
Repeatable
+
Profitable
```

trades.
