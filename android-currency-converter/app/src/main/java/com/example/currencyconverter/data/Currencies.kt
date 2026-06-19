package com.example.currencyconverter.data

import java.util.Currency
import java.util.Locale

/**
 * Human-friendly display names for currency codes. The vast majority resolve
 * via the JVM's built-in ISO 4217 data (`java.util.Currency`), so this works
 * for every standard currency the rates API returns without a hand-maintained
 * table. [fallbackNames] only covers the few non-ISO codes the API uses (and
 * older Android versions that may lack a display name for some codes).
 */
object Currencies {

    fun nameFor(code: String): String {
        runCatching {
            val name = Currency.getInstance(code).getDisplayName(Locale.getDefault())
            if (name.isNotBlank() && !name.equals(code, ignoreCase = true)) {
                return name
            }
        }
        return fallbackNames[code] ?: code
    }

    // Codes that aren't standard ISO 4217 (or are commonly missing a name on
    // older Android API levels) but are returned by open.er-api.com.
    private val fallbackNames: Map<String, String> = mapOf(
        "CLF" to "Chilean Unit of Account (UF)",
        "CNH" to "Chinese Yuan (Offshore)",
        "FOK" to "Faroese Króna",
        "GGP" to "Guernsey Pound",
        "IMP" to "Isle of Man Pound",
        "JEP" to "Jersey Pound",
        "KID" to "Kiribati Dollar",
        "TVD" to "Tuvaluan Dollar",
        "XDR" to "IMF Special Drawing Rights",
        "ZWL" to "Zimbabwean Dollar"
    )
}
