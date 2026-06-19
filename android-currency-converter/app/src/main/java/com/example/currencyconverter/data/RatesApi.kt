package com.example.currencyconverter.data

import com.example.currencyconverter.data.model.RatesResponse
import com.example.currencyconverter.data.model.RatesSnapshot
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.json.Json
import okhttp3.OkHttpClient
import okhttp3.Request
import java.text.SimpleDateFormat
import java.util.Date
import java.util.Locale
import java.util.TimeZone
import java.util.concurrent.TimeUnit

/**
 * Thin network client for the free, key-less open.er-api.com exchange-rate
 * API, which covers ~160 ISO currencies. https://www.exchangerate-api.com/docs/free
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
        val url = "https://open.er-api.com/v6/latest/$baseCurrency"
        val request = Request.Builder().url(url).build()

        client.newCall(request).execute().use { response ->
            if (!response.isSuccessful) {
                throw RatesApiException("HTTP ${response.code} fetching rates")
            }
            val body = response.body?.string()
                ?: throw RatesApiException("Empty response body")

            val parsed = json.decodeFromString<RatesResponse>(body)

            if (parsed.result != "success" || parsed.rates.isEmpty()) {
                throw RatesApiException("Rates service returned no data")
            }

            RatesSnapshot(
                base = parsed.baseCode,
                date = formatDate(parsed.timeLastUpdateUnix),
                fetchedAtMillis = System.currentTimeMillis(),
                rates = parsed.rates
            )
        }
    }

    private fun formatDate(unixSeconds: Long): String {
        if (unixSeconds <= 0L) return ""
        val formatter = SimpleDateFormat("yyyy-MM-dd", Locale.US).apply {
            timeZone = TimeZone.getTimeZone("UTC")
        }
        return formatter.format(Date(unixSeconds * 1000))
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
