package tempconv

// CToF conversion Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC conversion Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK conversion Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) + (-AbsoluteZeroC)) }

// KToF conversion Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(Celsius(k) + AbsoluteZeroC)) }

// KToC conversion Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// CToK conversion Kelvin temperature to Celsius
func CToK(c Celsius) Kelvin { return Kelvin(c + (-AbsoluteZeroC)) }
