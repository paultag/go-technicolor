package technicolor

func (w Writer) Red() Writer {
	return w.Add(RedFg)
}

func (w Writer) Black() Writer {
	return w.Add(BlackFg)
}

func (w Writer) Green() Writer {
	return w.Add(GreenFg)
}

func (w Writer) Yellow() Writer {
	return w.Add(YellowFg)
}

func (w Writer) Blue() Writer {
	return w.Add(BlueFg)
}

func (w Writer) Magenta() Writer {
	return w.Add(MagentaFg)
}

func (w Writer) White() Writer {
	return w.Add(WhiteFg)
}

func (w Writer) ResetColor() Writer {
	return w.Reset().White()
}

func (w Writer) Cyan() Writer {
	return w.Add(CyanFg)
}

func Color(color string) ANSISequence {
	return NewANSISequence([]string{color}, 'm')
}

func ResetColor() ANSISequence {
	return Color("0")
}

func XTermColor(color string) ANSISequence {
	return NewANSISequence([]string{"38", "5", color}, 'm')
}

func XTermBackgroundColor(color string) ANSISequence {
	return NewANSISequence([]string{"48", "5", color}, 'm')
}

func XTermTheme(fg, bg string) ANSISequences {
	return ANSISequences{XTermColor(fg), XTermBackgroundColor(bg)}
}

var Primary = "025"
var PrimaryDarker = "024"
var PrimaryDarkest = "017"

var Base = "016"

var White = "231"

var PrimaryAlt = "038"
var PrimaryAltDarkest = "024"
var PrimaryAltDark = "038"
var PrimaryAltLight = "153"
var PrimaryAltLightest = "195"

var Secondary = "161"
var SecondaryDarkest = "088"
var SecondaryDark = "160"
var SecondaryLight = "174"
var SecondaryLightest = "224"

var GrayDark = "059"
var Gray = "059"
var GrayLight = "145"
var GrayLighter = "188"
var GrayLightest = "231"

var GrayWarmDark = "059"
var GrayWarmLight = "188"

var GrayCoolLight = "189"

var Gold = "214"
var GoldLight = "221"
var GoldLighter = "222"
var GoldLightest = "230"

var Green = "029"
var GreenLight = "071"
var GreenLighter = "109"
var GreenLightest = "194"

var CoolBlue = "024"
var CoolBlueLight = "067"
var CoolBlueLighter = "110"
var CoolBlueLightest = "189"

var Focus = "068"
var Visited = "054"

var BlackFg = XTermColor(Base)
var RedFg = XTermColor(SecondaryDark)
var GreenFg = XTermColor(GreenLight)
var YellowFg = XTermColor(GoldLight)
var BlueFg = XTermColor(PrimaryAltLightest)
var MagentaFg = XTermColor(Visited)
var CyanFg = XTermColor(Focus)
var WhiteFg = Color("0")

var PrimaryDarkestOnWhite = XTermTheme(PrimaryDarkest, White)
var PrimaryDarkerOnWhite = XTermTheme(PrimaryDarker, White)
var PrimaryOnWhite = XTermTheme(Primary, White)
var CoolBlueLightOnWhite = XTermTheme(CoolBlueLight, White)
var PrimaryAltDarkestOnWhite = XTermTheme(PrimaryAltDarkest, White)
var GreenOnWhite = XTermTheme(Green, White)
var VisitedOnWhite = XTermTheme(Visited, White)
var BaseOnWhite = XTermTheme(Base, White)
var GrayDarkOnWhite = XTermTheme(GrayDark, White)
var GrayOnWhite = XTermTheme(Gray, White)
var GrayWarmDarkOnWhite = XTermTheme(GrayWarmDark, White)
var SecondaryDarkestOnWhite = XTermTheme(SecondaryDarkest, White)
var SecondaryDarkOnWhite = XTermTheme(SecondaryDark, White)
var SecondaryOnWhite = XTermTheme(Secondary, White)

var WhiteOnBase = XTermTheme(White, Base)
var WhiteOnGrayWarmDark = XTermTheme(White, GrayWarmDark)
var WhiteOnGrayDark = XTermTheme(White, GrayDark)
var WhiteOnGray = XTermTheme(White, Gray)
var WhiteOnPrimaryDarkest = XTermTheme(White, PrimaryDarkest)
var WhiteOnPrimaryDarker = XTermTheme(White, PrimaryDarker)
var WhiteOnPrimary = XTermTheme(White, Primary)
var WhiteOnCoolBlueLight = XTermTheme(White, CoolBlueLight)
var WhiteOnPrimaryAltDarkest = XTermTheme(White, PrimaryAltDarkest)
var BaseOnPrimaryAltDark = XTermTheme(Base, PrimaryAltDark)
var BaseOnPrimaryAlt = XTermTheme(Base, PrimaryAlt)
var WhiteOnGreen = XTermTheme(White, Green)
var BaseOnGreenLight = XTermTheme(Base, GreenLight)
var BaseOnGold = XTermTheme(Base, Gold)
var BaseOnGoldLight = XTermTheme(Base, GoldLight)
var WhiteOnSecondaryDarkest = XTermTheme(White, SecondaryDarkest)
var WhiteOnSecondaryDark = XTermTheme(White, SecondaryDark)
var WhiteOnSecondary = XTermTheme(White, Secondary)

var BaseOnGrayLight = XTermTheme(Base, GrayLight)
var BaseOnGrayLighter = XTermTheme(Base, GrayLighter)
var BaseOnGrayWarmLight = XTermTheme(Base, GrayWarmLight)
var BaseOnCoolBlueLighter = XTermTheme(Base, CoolBlueLighter)
var BaseOnCoolBlueLightest = XTermTheme(Base, CoolBlueLightest)
var BaseOnPrimaryAltLightest = XTermTheme(Base, PrimaryAltLightest)
var BaseOnGreenLighter = XTermTheme(Base, GreenLighter)
var BaseOnGreenLightest = XTermTheme(Base, GreenLightest)
var BaseOnGoldLighter = XTermTheme(Base, GoldLighter)
var BaseOnGoldLightest = XTermTheme(Base, GoldLightest)
var BaseOnSecondaryLightest = XTermTheme(Base, SecondaryLightest)
