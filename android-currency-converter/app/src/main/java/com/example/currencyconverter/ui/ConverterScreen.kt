package com.example.currencyconverter.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material.icons.filled.SwapVert
import androidx.compose.material.icons.outlined.CloudOff
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ExposedDropdownMenuBox
import androidx.compose.material3.ExposedDropdownMenuDefaults
import androidx.compose.material3.FilledTonalIconButton
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.currencyconverter.data.Currencies
import java.text.NumberFormat
import java.text.SimpleDateFormat
import java.util.Currency
import java.util.Date
import java.util.Locale

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ConverterScreen(viewModel: ConverterViewModel = viewModel()) {
    val state by viewModel.uiState.collectAsStateWithLifecycle()

    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text("Currency Converter") },
                actions = {
                    IconButton(onClick = viewModel::refresh) {
                        Icon(Icons.Filled.Refresh, contentDescription = "Refresh rates")
                    }
                },
                colors = TopAppBarDefaults.topAppBarColors(
                    containerColor = MaterialTheme.colorScheme.primary,
                    titleContentColor = MaterialTheme.colorScheme.onPrimary,
                    actionIconContentColor = MaterialTheme.colorScheme.onPrimary,
                )
            )
        }
    ) { innerPadding ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(innerPadding)
                .padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(16.dp)
        ) {
            when {
                state.isLoading && state.snapshot == null -> LoadingState()
                state.snapshot == null -> ErrorState(
                    message = state.errorMessage ?: "Unable to load rates.",
                    onRetry = viewModel::refresh
                )
                else -> ConverterContent(state = state, viewModel = viewModel)
            }
        }
    }
}

@Composable
private fun LoadingState() {
    Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
        Column(horizontalAlignment = Alignment.CenterHorizontally) {
            CircularProgressIndicator()
            Spacer(Modifier.height(12.dp))
            Text("Fetching the latest rates…")
        }
    }
}

@Composable
private fun ErrorState(message: String, onRetry: () -> Unit) {
    Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
        Column(
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            Icon(
                Icons.Outlined.CloudOff,
                contentDescription = null,
                modifier = Modifier.size(48.dp),
                tint = MaterialTheme.colorScheme.error
            )
            Text(message, style = MaterialTheme.typography.bodyLarge)
            FilledTonalIconButton(onClick = onRetry) {
                Icon(Icons.Filled.Refresh, contentDescription = "Retry")
            }
        }
    }
}

@Composable
private fun ConverterContent(
    state: ConverterUiState,
    viewModel: ConverterViewModel,
) {
    if (state.isOffline) {
        OfflineBanner(date = state.snapshot?.date)
    }

    // Amount input
    OutlinedTextField(
        value = state.amountInput,
        onValueChange = viewModel::onAmountChanged,
        label = { Text("Amount") },
        singleLine = true,
        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Decimal),
        modifier = Modifier.fillMaxWidth()
    )

    // From / Swap / To — full width and searchable.
    SearchableCurrencyField(
        label = "From",
        selected = state.fromCurrency,
        options = state.availableCurrencies,
        onSelected = viewModel::onFromCurrencyChanged,
        modifier = Modifier.fillMaxWidth()
    )
    Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
        FilledTonalIconButton(onClick = viewModel::swapCurrencies) {
            Icon(Icons.Filled.SwapVert, contentDescription = "Swap currencies")
        }
    }
    SearchableCurrencyField(
        label = "To",
        selected = state.toCurrency,
        options = state.availableCurrencies,
        onSelected = viewModel::onToCurrencyChanged,
        modifier = Modifier.fillMaxWidth()
    )

    ResultCard(state = state)
}

@Composable
private fun OfflineBanner(date: String?) {
    Card(
        modifier = Modifier.fillMaxWidth(),
        colors = CardDefaults.cardColors(
            containerColor = MaterialTheme.colorScheme.secondaryContainer
        )
    ) {
        Row(
            modifier = Modifier.padding(12.dp),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(8.dp)
        ) {
            Icon(Icons.Outlined.CloudOff, contentDescription = null)
            Text(
                text = "Offline — showing saved rates" +
                    (date?.let { " from $it" } ?: ""),
                style = MaterialTheme.typography.bodyMedium
            )
        }
    }
}

@Composable
private fun ResultCard(state: ConverterUiState) {
    Card(
        modifier = Modifier.fillMaxWidth(),
        colors = CardDefaults.cardColors(
            containerColor = MaterialTheme.colorScheme.primaryContainer
        )
    ) {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .padding(20.dp),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.spacedBy(4.dp)
        ) {
            val converted = state.convertedAmount
            if (converted == null) {
                Text(
                    "Enter an amount to convert",
                    style = MaterialTheme.typography.titleMedium
                )
            } else {
                Text(
                    text = "${formatAmount(state.amountInput.toDoubleOrNull() ?: 0.0, state.fromCurrency)} ${state.fromCurrency} =",
                    style = MaterialTheme.typography.bodyLarge
                )
                Text(
                    text = "${formatAmount(converted, state.toCurrency)} ${state.toCurrency}",
                    style = MaterialTheme.typography.headlineMedium,
                    fontWeight = FontWeight.Bold,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
                state.snapshot?.let { snap ->
                    val perUnit = snap.convert(1.0, state.fromCurrency, state.toCurrency)
                    if (perUnit != null) {
                        Text(
                            text = "1 ${state.fromCurrency} = ${formatRate(perUnit)} ${state.toCurrency}",
                            style = MaterialTheme.typography.bodySmall
                        )
                    }
                    Text(
                        text = "Rates as of ${snap.date} · updated ${formatTimestamp(snap.fetchedAtMillis)}",
                        style = MaterialTheme.typography.labelSmall
                    )
                }
            }
        }
    }
}

/**
 * A full-width currency picker you can search by typing the first few letters
 * of the code (e.g. "eu") or the name (e.g. "yen"). Prefix matches on the code
 * are surfaced first.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun SearchableCurrencyField(
    label: String,
    selected: String,
    options: List<String>,
    onSelected: (String) -> Unit,
    modifier: Modifier = Modifier,
) {
    var expanded by remember { mutableStateOf(false) }
    var query by remember { mutableStateOf("") }

    val matches = remember(query, options) {
        val q = query.trim().lowercase()
        if (q.isEmpty()) {
            options
        } else {
            options
                .filter { code ->
                    code.lowercase().contains(q) ||
                        Currencies.nameFor(code).lowercase().contains(q)
                }
                .sortedBy { code -> if (code.lowercase().startsWith(q)) 0 else 1 }
        }
    }

    ExposedDropdownMenuBox(
        expanded = expanded,
        onExpandedChange = { open ->
            expanded = open
            if (!open) query = ""
        },
        modifier = modifier
    ) {
        OutlinedTextField(
            // When searching, show what the user typed; otherwise the selection.
            value = if (expanded) query else "$selected — ${Currencies.nameFor(selected)}",
            onValueChange = {
                query = it
                expanded = true
            },
            singleLine = true,
            label = { Text(label) },
            placeholder = { Text("Search currency…") },
            trailingIcon = {
                ExposedDropdownMenuDefaults.TrailingIcon(expanded = expanded)
            },
            modifier = Modifier
                .menuAnchor()
                .fillMaxWidth()
        )
        ExposedDropdownMenu(
            expanded = expanded,
            onDismissRequest = {
                expanded = false
                query = ""
            }
        ) {
            if (matches.isEmpty()) {
                DropdownMenuItem(
                    text = { Text("No matching currency") },
                    onClick = {},
                    enabled = false
                )
            } else {
                matches.forEach { code ->
                    DropdownMenuItem(
                        text = { Text("$code — ${Currencies.nameFor(code)}") },
                        onClick = {
                            onSelected(code)
                            query = ""
                            expanded = false
                        }
                    )
                }
            }
        }
    }
}

// --- formatting helpers ---

private fun formatAmount(value: Double, currencyCode: String): String {
    val format = NumberFormat.getNumberInstance(Locale.getDefault())
    val fractionDigits = runCatching {
        Currency.getInstance(currencyCode).defaultFractionDigits
    }.getOrDefault(2).coerceAtLeast(0)
    format.minimumFractionDigits = fractionDigits
    format.maximumFractionDigits = fractionDigits
    return format.format(value)
}

private fun formatRate(value: Double): String {
    val format = NumberFormat.getNumberInstance(Locale.getDefault())
    format.minimumFractionDigits = 2
    format.maximumFractionDigits = 6
    return format.format(value)
}

private fun formatTimestamp(millis: Long): String {
    val formatter = SimpleDateFormat("MMM d, HH:mm", Locale.getDefault())
    return formatter.format(Date(millis))
}
