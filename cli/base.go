package cli

import (
	"fmt"
	"github.com/kerogs/KerogsGo/colors"
	"os"
	"os/exec"
)

func AsciiStart(version string) {
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
	fmt.Println("  '!!$$!!!!!!oneFileTransfer!!!!!!!!!!!!!!!oooooo..  ''   ,$")
	fmt.Println("   '!!$!!!!!!!!v" + version + " kerogs!!!!!!!!!!!!!!!!!!!!!!!!oooo..$$")
	fmt.Println("    !!$!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!$'")
	fmt.Println("    '$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$!!!!!!!!!!!!!!!!!!,")
	fmt.Println(".....$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$.....")
	fmt.Println(colors.Reset)
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
