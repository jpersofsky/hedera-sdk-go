package com.example.currencyconverter.data

import com.example.currencyconverter.data.model.RatesResponse
import com.example.currencyconverter.data.model.RatesSnapshot
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.json.Json
import okhttp3.OkHttpClient
import okhttp3.Request
import java.util.concurrent.TimeUnit

/**
 * Thin network client for the free, key-less Frankfurter exchange-rate API.
 * https://www.frankfurter.app/docs/
 */
class RatesApi(
    private val baseCurrency: String = "USD",
    private val client: OkHttpClient = defaultClient,
) {
    private val json = Json { ignoreUnknownKeys = true }

    /**
     * Fetch the latest rates. Runs on the IO dispatcher. Throws on network
     * or parsing failure so the repository can fall back to cache.
     */
    suspend fun fetchLatest(): RatesSnapshot = withContext(Dispatchers.IO) {
        val url = "https://api.frankfurter.app/latest?base=$baseCurrency"
        val request = Request.Builder().url(url).build()

        client.newCall(request).execute().use { response ->
            if (!response.isSuccessful) {
                throw RatesApiException("HTTP ${response.code} fetching rates")
            }
            val body = response.body?.string()
                ?: throw RatesApiException("Empty response body")

            val parsed = json.decodeFromString<RatesResponse>(body)

            // Frankfurter omits the base currency from the rates map; add it
            // back so conversions involving the base work uniformly.
            val rates = parsed.rates.toMutableMap()
            rates[parsed.base] = 1.0

            RatesSnapshot(
                base = parsed.base,
                date = parsed.date,
                fetchedAtMillis = System.currentTimeMillis(),
                rates = rates
            )
        }
    }

    companion object {
        private val defaultClient: OkHttpClient by lazy {
            OkHttpClient.Builder()
                .connectTimeout(10, TimeUnit.SECONDS)
                .readTimeout(10, TimeUnit.SECONDS)
                .build()
        }
    }
}

class RatesApiException(message: String) : Exception(message)
