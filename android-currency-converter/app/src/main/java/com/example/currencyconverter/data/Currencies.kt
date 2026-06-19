package com.example.currencyconverter.data

/**
 * Human-friendly display names for the currencies supported by the
 * Frankfurter API. Used to show "US Dollar" next to "USD" in the picker.
 * Codes not present here simply fall back to showing the raw code.
 */
object Currencies {
    val displayNames: Map<String, String> = mapOf(
        "AUD" to "Australian Dollar",
        "BGN" to "Bulgarian Lev",
        "BRL" to "Brazilian Real",
        "CAD" to "Canadian Dollar",
        "CHF" to "Swiss Franc",
        "CNY" to "Chinese Yuan",
        "CZK" to "Czech Koruna",
        "DKK" to "Danish Krone",
        "EUR" to "Euro",
        "GBP" to "British Pound",
        "HKD" to "Hong Kong Dollar",
        "HUF" to "Hungarian Forint",
        "IDR" to "Indonesian Rupiah",
        "ILS" to "Israeli Shekel",
        "INR" to "Indian Rupee",
        "ISK" to "Icelandic Krona",
        "JPY" to "Japanese Yen",
        "KRW" to "South Korean Won",
        "MXN" to "Mexican Peso",
        "MYR" to "Malaysian Ringgit",
        "NOK" to "Norwegian Krone",
        "NZD" to "New Zealand Dollar",
        "PHP" to "Philippine Peso",
        "PLN" to "Polish Zloty",
        "RON" to "Romanian Leu",
        "SEK" to "Swedish Krona",
        "SGD" to "Singapore Dollar",
        "THB" to "Thai Baht",
        "TRY" to "Turkish Lira",
        "USD" to "US Dollar",
        "ZAR" to "South African Rand"
    )

    fun nameFor(code: String): String = displayNames[code] ?: code
}
