package com.example.currencyconverter.ui

import android.app.Application
import androidx.lifecycle.AndroidViewModel
import androidx.lifecycle.viewModelScope
import com.example.currencyconverter.data.RatesApi
import com.example.currencyconverter.data.RatesCache
import com.example.currencyconverter.data.RatesRepository
import com.example.currencyconverter.data.model.RatesSnapshot
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

/** Immutable snapshot of everything the converter screen renders. */
data class ConverterUiState(
    val amountInput: String = "1",
    val fromCurrency: String = "USD",
    val toCurrency: String = "EUR",
    val snapshot: RatesSnapshot? = null,
    val isLoading: Boolean = true,
    val isOffline: Boolean = false,
    val errorMessage: String? = null,
) {
    val availableCurrencies: List<String>
        get() = snapshot?.currencyCodes ?: emptyList()

    /** Converted amount, or null if input is blank/invalid or rates missing. */
    val convertedAmount: Double?
        get() {
            val amount = amountInput.toDoubleOrNull() ?: return null
            return snapshot?.convert(amount, fromCurrency, toCurrency)
        }
}

class ConverterViewModel(app: Application) : AndroidViewModel(app) {

    private val repository = RatesRepository(
        api = RatesApi(baseCurrency = "USD"),
        cache = RatesCache(app.applicationContext)
    )

    private val _uiState = MutableStateFlow(ConverterUiState())
    val uiState: StateFlow<ConverterUiState> = _uiState.asStateFlow()

    init {
        refresh()
    }

    fun refresh() {
        _uiState.update { it.copy(isLoading = true, errorMessage = null) }
        viewModelScope.launch {
            val result = repository.getRates()
            if (result == null) {
                _uiState.update {
                    it.copy(
                        isLoading = false,
                        isOffline = true,
                        errorMessage = "No internet connection and no saved rates yet. " +
                            "Connect once to download the latest rates."
                    )
                }
                return@launch
            }
            _uiState.update { state ->
                state.copy(
                    snapshot = result.snapshot,
                    isLoading = false,
                    isOffline = result.fromCache,
                    errorMessage = result.staleReason,
                )
            }
        }
    }

    fun onAmountChanged(value: String) {
        // Keep only digits and a single decimal separator.
        val sanitized = value
            .filter { it.isDigit() || it == '.' }
            .let { input ->
                val firstDot = input.indexOf('.')
                if (firstDot == -1) input
                else input.substring(0, firstDot + 1) +
                    input.substring(firstDot + 1).replace(".", "")
            }
        _uiState.update { it.copy(amountInput = sanitized) }
    }

    fun onFromCurrencyChanged(code: String) {
        _uiState.update { it.copy(fromCurrency = code) }
    }

    fun onToCurrencyChanged(code: String) {
        _uiState.update { it.copy(toCurrency = code) }
    }

    fun swapCurrencies() {
        _uiState.update {
            it.copy(fromCurrency = it.toCurrency, toCurrency = it.fromCurrency)
        }
    }
}
