// Copyright © 2019 TJ Hoplock <t.hoplock@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var pasteCmd = &cobra.Command{
	Use:     "paste",
	Aliases: []string{"out", "print"},
	Short:   "Print clipboard contents to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		err := writeClipboardToStdout()
		if err != nil {
			fmt.Printf("Call to print clipboard contents failed: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)
}

func writeClipboardToStdout() error {
	str, err := clipboard.ReadAll()
	if err != nil {
		return fmt.Errorf("Failed to dump clipboard contents to variable: %v\n", err)
	}

	fmt.Println(str)
	return nil
}
