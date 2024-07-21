package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/kerogs/KerogsGo/colors"
)

// ClearScreen clears the terminal screen.
//
// This function executes the "cls" command in the command prompt to clear the screen.
// It uses the "cmd" package to execute the command and sets the output to the standard output.
//
// There are no parameters.
// There is no return value.
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CliReturn moves the cursor up the console screen by the specified number of times and clears the line.
//
// Parameters:
// - repeat: the number of times to move the cursor up and clear the line (int)
//
// Return type: None.
func CliReturn(repeat int) {
	for i := 0; i < repeat; i++ {
		fmt.Printf("\033[1A\033[K")
	}
}

// StopProgram prints a message asking the user to press a key to stop the program and then exits the program.
//
// No parameters.
// No return value.
func StopProgram(){
	fmt.Println("\n\n\n"+colors.Red+"Press a key to stop the program"+colors.Reset)
	fmt.Scanln()

	os.Exit(0)
}

// StopProgramMsg prints the given message and prompts the user to press a key to stop the program.
//
// Parameters:
// - message: the message to be printed (string)
//
// Return type: None.
func StopProgramMsg(message string) {
	fmt.Println(colors.Red+message+colors.Reset)

	fmt.Println("\n\n\n"+colors.Red+"Press a key to stop the program"+colors.Reset)
	fmt.Scanln()

	os.Exit(0)
}

// StopProgramErr prints the given error and message, prompts the user to press a key to stop the program, and exits the program.
//
// Parameters:
// - errorReturn: the error to be printed (error)
// - message: the message to be printed (string)
//
// Return type: None.
func StopProgramErr(errorReturn error, message string) {
	fmt.Println(errorReturn)
	
	if(message != "") {
		fmt.Println(colors.Red+message+colors.Reset)
	}

	fmt.Println("\n\n\n"+colors.Red+"Press a key to stop the program"+colors.Reset)
	fmt.Scanln()

	os.Exit(0)
}