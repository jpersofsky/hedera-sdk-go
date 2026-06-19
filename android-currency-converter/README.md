# Currency Converter (Android)

A small, self-contained Android currency-converter app built with **Kotlin** and
**Jetpack Compose**. It converts between ~30 world currencies using live
exchange rates, and keeps working offline from the last rates it downloaded.

> Note: this app lives in a subdirectory of the `hedera-sdk-go` repository but is
> entirely independent of the Go SDK — it has its own Gradle build and does not
> touch any SDK code.

## Features

- **Hybrid rates (online + offline):** fetches the latest rates from the free,
  key-less [Frankfurter API](https://www.frankfurter.app/) when there's a
  connection, caches them on-device with DataStore, and transparently falls back
  to the cached rates when offline. An "Offline — showing saved rates" banner
  appears when cached data is in use.
- **Live conversion** as you type, with a swap button to flip the currencies.
- **Material 3 UI** with dynamic color on Android 12+, light/dark support.
- **Locale-aware formatting** (correct fraction digits per currency, e.g. JPY
  shows no decimals).

## Project layout

```
app/src/main/java/com/example/currencyconverter/
├── MainActivity.kt              # Compose entry point
├── data/
│   ├── Currencies.kt            # code -> display-name lookup
│   ├── RatesApi.kt              # OkHttp client for the Frankfurter API
│   ├── RatesCache.kt            # DataStore-backed offline cache
│   ├── RatesRepository.kt       # hybrid: network-first, cache fallback
│   └── model/RatesModels.kt     # DTOs + RatesSnapshot (conversion math)
└── ui/
    ├── ConverterViewModel.kt    # UI state + actions
    ├── ConverterScreen.kt       # Compose UI
    └── theme/Theme.kt           # Material 3 theme
```

## How rates work

Rates are stored relative to a single base currency (USD). To convert an amount
from currency *A* to currency *B*:

```
amountInBase = amount / rate[A]      // rate[A] = units of A per 1 USD
result       = amountInBase * rate[B]
```

This logic lives in `RatesSnapshot.convert(...)` and is covered by unit tests in
`app/src/test/.../RatesSnapshotTest.kt`.

## Building & running

You need the Android SDK (Android Studio is the easiest path) and JDK 17+.
Network access is required on first build to download Gradle dependencies.

**Android Studio:** open the `android-currency-converter/` folder and press Run.

**Command line:**

```bash
cd android-currency-converter

# Run the unit tests
./gradlew test

# Build a debug APK (output in app/build/outputs/apk/debug/)
./gradlew assembleDebug

# Install onto a connected device / emulator
./gradlew installDebug
```

If Gradle complains it can't find the Android SDK, create a `local.properties`
file with `sdk.dir=/path/to/Android/sdk` (Android Studio does this for you).

## Configuration

- **Min SDK:** 24 (Android 7.0) · **Target SDK:** 34
- **Base currency / API:** change `RatesApi(baseCurrency = "USD")` in
  `ConverterViewModel.kt` to use a different base. The Frankfurter API requires
  no API key.
