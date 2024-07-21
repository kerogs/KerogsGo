package base

import (
	"fmt"
	"github.com/kerogs/KerogsGo/colors"
)

// AsciiStart prints a Kerogs watermark in ASCII art.
//
// No parameters.
// No return value.
func AsciiStart() {
	fmt.Println(colors.Red)
	fmt.Println("                       $o")
	fmt.Println("                       $                     .........")
	fmt.Println("                      $$$      .oo..     'oooo'oooo'ooooooooo....")
	fmt.Println("                       $       $$$$$$$")
	fmt.Println("                   .ooooooo.   $$!!!!!")
	fmt.Println("                 .'.........'. $$!!!!!      o$$oo.   ...oo,oooo,oooo'ooo''")
	fmt.Println("    $          .o'  oooooo   '.$$!!!!!      $$!!!!!       'oo''oooo''")
	fmt.Println(" ..o$ooo...    $                '!!''!.     $$!!!!!")
	fmt.Println(" $    ..  '''oo$$$$$$$$$$$$$.    '    'oo.  $$!!!!!")
	fmt.Println(" !.......      '''..$$ $$ $$$   ..        '.$$!!''!")
	fmt.Println(" !!$$$!!!!!!!!oooo......   '''  $$ $$ :o           'oo.")
	fmt.Println(" !!$$$!!!$$!$$!!!!!!!!!!oo.....     ' ''  o$$o .      ''oo..")
	fmt.Println(" !!!$$!!!!!!!!!!!!!!!!!!!!!!!!!!!!ooooo..      'o  oo..    $")
	fmt.Println("  '!!$$!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!oooooo..  ''   ,$")
	fmt.Println("   '!!$!!!!!!!!!!!!!!!Kerogs!!!!!!!!!!!!!!!!!!!!!!!!oooo..$$")
	fmt.Println("    !!$!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!$'")
	fmt.Println("    '$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$!!!!!!!!!!!!!!!!!!,")
	fmt.Println(".....$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$.....")
	fmt.Println(colors.Reset)
}

