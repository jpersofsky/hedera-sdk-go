package com.example.currencyconverter.data

import com.example.currencyconverter.data.model.RatesSnapshot

/**
 * Outcome of a rates load, capturing where the data came from so the UI can
 * tell the user whether they're looking at live or cached numbers.
 */
data class RatesResult(
    val snapshot: RatesSnapshot,
    val fromCache: Boolean,
    /** Set when a live fetch failed but cached data was available. */
    val staleReason: String? = null
)

/**
 * Hybrid rates source: prefer a live network fetch, but transparently fall
 * back to the last cached snapshot when offline or when the request fails.
 * Every successful live fetch refreshes the cache.
 */
class RatesRepository(
    private val api: RatesApi,
    private val cache: RatesCache,
) {
    /**
     * Load rates. Tries the network first; on failure returns cached rates if
     * present. Returns null only when there is neither network nor cache.
     */
    suspend fun getRates(): RatesResult? {
        return try {
            val fresh = api.fetchLatest()
            cache.save(fresh)
            RatesResult(snapshot = fresh, fromCache = false)
        } catch (e: Exception) {
            val cached = cache.load() ?: return null
            RatesResult(
                snapshot = cached,
                fromCache = true,
                staleReason = e.message ?: "Couldn't reach the rates service"
            )
        }
    }
}
