package com.example.currencyconverter.data.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * Raw response shape returned by the free open.er-api.com endpoint
 * (e.g. https://open.er-api.com/v6/latest/USD), which covers ~160 ISO
 * currencies with no API key required.
 *
 * Unlike some providers, this one *does* include the base currency in
 * [rates] (with a value of 1.0).
 */
@Serializable
data class RatesResponse(
    val result: String = "",
    @SerialName("base_code") val baseCode: String = "USD",
    @SerialName("time_last_update_unix") val timeLastUpdateUnix: Long = 0L,
    @SerialName("time_last_update_utc") val timeLastUpdateUtc: String = "",
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
