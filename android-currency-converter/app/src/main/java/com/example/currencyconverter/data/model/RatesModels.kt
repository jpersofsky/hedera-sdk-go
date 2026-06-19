package com.example.currencyconverter.data.model

import kotlinx.serialization.Serializable

/**
 * Raw response shape returned by the Frankfurter API
 * (e.g. https://api.frankfurter.app/latest?base=USD).
 *
 * Note: the API does not include the base currency itself in [rates]
 * (its rate is implicitly 1.0); the repository adds it explicitly.
 */
@Serializable
data class RatesResponse(
    val base: String = "USD",
    val date: String = "",
    val rates: Map<String, Double> = emptyMap()
)

/**
 * A normalized, self-contained snapshot of exchange rates that the app
 * works with. Every rate is expressed as "units of currency per 1 [base]".
 *
 * This is the type that gets cached to disk so the app can convert offline.
 */
@Serializable
data class RatesSnapshot(
    val base: String,
    /** ISO date the rates were published for, e.g. "2024-06-18". */
    val date: String,
    /** Epoch millis when we fetched these rates. */
    val fetchedAtMillis: Long,
    val rates: Map<String, Double>
) {
    /**
     * Convert [amount] from [from] currency to [to] currency using these
     * rates. Returns null if either currency is missing from the snapshot.
     */
    fun convert(amount: Double, from: String, to: String): Double? {
        val fromRate = rates[from] ?: return null
        val toRate = rates[to] ?: return null
        if (fromRate == 0.0) return null
        // amount(from) -> amount(base) -> amount(to)
        val amountInBase = amount / fromRate
        return amountInBase * toRate
    }

    val currencyCodes: List<String>
        get() = rates.keys.sorted()
}
