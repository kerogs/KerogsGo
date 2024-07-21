# Go terminal colors
KerogsGo is just another function I use often. It's simply intended to simplify my life.

## Use
1. import
```sh
go get "github.com/kerogs/KerogsGo"
```
2. use
```Go
import "github.com/kerogs/KerogsGo/[Import]"
```

## Func 
|Name|Import|In|Out|Desc|
|-|-|-|
|colors.[Colors]|``/colors``|N/A|N/A|Add colors to CLI|
|base.AsciiStart|``/base``|N/A|N/A|Add kerogs watermark|
|terminal.ClearScreen|``/terminal``|N/A|N/A|Clear cmd|
|terminal.CliReturn|``/terminal``|``repeat`` int|N/A|return x line above|
|terminal.StopProgram|``/terminal``|N/A|N/A|Stop the program, but first wait for an interaction.|
|terminal.StopProgramMsg|``/terminal``|``message`` string|N/A|Same as above, just add a custom message directly with|
|terminal.StopProgramErr|``/terminal``|``errorReturn`` error, ``message`` string|N/A|Same as above, just you can add a custom message directly to it and also give the error.|
|filedir.FileMake|``/filedir``|``name`` string, ``content`` string|N/A|make a file + write in|

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