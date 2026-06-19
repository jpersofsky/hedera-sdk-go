package com.example.currencyconverter

import com.example.currencyconverter.data.model.RatesSnapshot
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

class RatesSnapshotTest {

    private val snapshot = RatesSnapshot(
        base = "USD",
        date = "2024-06-18",
        fetchedAtMillis = 0L,
        rates = mapOf(
            "USD" to 1.0,
            "EUR" to 0.5,   // 1 USD = 0.5 EUR
            "JPY" to 100.0  // 1 USD = 100 JPY
        )
    )

    @Test
    fun convertsFromBaseCurrency() {
        // 10 USD -> EUR at 0.5 = 5 EUR
        assertEquals(5.0, snapshot.convert(10.0, "USD", "EUR")!!, 1e-9)
    }

    @Test
    fun convertsToBaseCurrency() {
        // 5 EUR -> USD = 10 USD
        assertEquals(10.0, snapshot.convert(5.0, "EUR", "USD")!!, 1e-9)
    }

    @Test
    fun convertsBetweenTwoNonBaseCurrencies() {
        // 10 EUR -> JPY: 10 / 0.5 = 20 USD -> * 100 = 2000 JPY
        assertEquals(2000.0, snapshot.convert(10.0, "EUR", "JPY")!!, 1e-9)
    }

    @Test
    fun sameCurrencyReturnsSameAmount() {
        assertEquals(42.0, snapshot.convert(42.0, "EUR", "EUR")!!, 1e-9)
    }

    @Test
    fun unknownCurrencyReturnsNull() {
        assertNull(snapshot.convert(1.0, "USD", "GBP"))
        assertNull(snapshot.convert(1.0, "GBP", "USD"))
    }

    @Test
    fun currencyCodesAreSorted() {
        assertEquals(listOf("EUR", "JPY", "USD"), snapshot.currencyCodes)
    }

    @Test
    fun zeroSourceRateReturnsNull() {
        val broken = snapshot.copy(rates = snapshot.rates + ("BAD" to 0.0))
        assertNull(broken.convert(1.0, "BAD", "USD"))
    }

    @Test
    fun roundTripConversionIsStable() {
        val usd = 123.45
        val eur = snapshot.convert(usd, "USD", "EUR")!!
        val back = snapshot.convert(eur, "EUR", "USD")!!
        assertTrue(kotlin.math.abs(usd - back) < 1e-9)
    }
}
