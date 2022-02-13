package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//var resourceSimheiTtf = &fyne.StaticResource{
//	StaticName: "simhei.ttf",
//}

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

// return bundled font resource
func (*MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceAlibabaPuHuiTiMediumTtf
}

func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
