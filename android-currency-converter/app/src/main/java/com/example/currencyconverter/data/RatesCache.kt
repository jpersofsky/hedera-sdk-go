package com.example.currencyconverter.data

import android.content.Context
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import androidx.datastore.preferences.preferencesDataStore
import com.example.currencyconverter.data.model.RatesSnapshot
import kotlinx.coroutines.flow.first
import kotlinx.serialization.json.Json

private val Context.dataStore by preferencesDataStore(name = "rates_cache")

/**
 * Persists the most recently fetched [RatesSnapshot] to disk via DataStore
 * so the converter keeps working when the device is offline.
 */
class RatesCache(private val context: Context) {
    private val json = Json { ignoreUnknownKeys = true }
    private val key = stringPreferencesKey("latest_snapshot")

    suspend fun save(snapshot: RatesSnapshot) {
        context.dataStore.edit { prefs ->
            prefs[key] = json.encodeToString(RatesSnapshot.serializer(), snapshot)
        }
    }

    suspend fun load(): RatesSnapshot? {
        val raw = context.dataStore.data.first()[key] ?: return null
        return runCatching {
            json.decodeFromString(RatesSnapshot.serializer(), raw)
        }.getOrNull()
    }
}
