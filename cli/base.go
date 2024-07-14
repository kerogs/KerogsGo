package cli

import (
	"fmt"
	"github.com/kerogs/KerogsGo/colors"
	"os"
	"os/exec"
)

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

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CliReturn(repeat int) {
	for i := 0; i < repeat; i++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func FileMake(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("file creation error :", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}

}