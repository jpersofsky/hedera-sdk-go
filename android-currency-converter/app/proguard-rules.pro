# Add project specific ProGuard rules here.

# Keep kotlinx.serialization generated serializers.
-keepattributes *Annotation*, InnerClasses
-dontnote kotlinx.serialization.**
-keepclassmembers class **$$serializer { *; }
-keepclasseswithmembers class * {
    kotlinx.serialization.KSerializer serializer(...);
}
-keep,includedescriptorclasses class com.example.currencyconverter.**$$serializer { *; }
-keepclassmembers class com.example.currencyconverter.** {
    *** Companion;
}
