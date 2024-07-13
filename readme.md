# Go terminal colors

## Use
1. import
```sh
go get "github.com/kerogs/KerogsGo"
```
2. use
```Go
import "github.com/kerogs/KerogsGo/[pkg]"
```

## Func 
|Name|Import|Desc|
|-|-|-|
|colors.[Colors]|``/colors``|Add colors to CLI|
|cli.AsciiStart|``/cli``|Add kerogs watermark|
|cli.ClearScreen|``/cli``|Clear cmd|

## Colors
|Call|Name|Code|
|:---|:---|:---|
colors.Reset|Reset|\033[0m|
colors.Black|Black|\033[30m|
colors.Red|Red|\033[31m|
colors.Green|Green|\033[32m|
colors.Yellow|Yellow|\033[33m|
colors.Blue|Blue|\033[34m|
colors.Magenta|Magenta|\033[35m|
colors.Cyan|Cyan|\033[36m|
colors.White|White|\033[37m|
colors.Gray|Gray|\033[90m|
colors.LightRed|LightRed|\033[91m|
colors.LightGreen|LightGreen|\033[92m|
colors.LightYellow|LightYellow|\033[93m|
colors.LightBlue|LightBlue|\033[94m|
colors.LightMagenta|LightMagenta|\033[95m|
colors.LightCyan|LightCyan|\033[96m|
colors.LightWhite|LightWhite|\033[97m|
colors.Orange|Orange|\033[38;5;208m|