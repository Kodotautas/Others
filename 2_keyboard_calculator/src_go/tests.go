package main

// func main() {
// 	// switch stdin into 'raw' mode
// 	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer term.Restore(int(os.Stdin.Fd()), oldState)

// 	b := make([]byte, 1)
// 	_, err = os.Stdin.Read(b)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("the char %q was hit", string(b[0]))
// }
