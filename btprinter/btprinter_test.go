package btprinter

import (
	"fmt"
	"testing"
)

func TestPrintTree(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		fmt.Println("=== printTreeStringTest begin ====")
		fmt.Println("printTree(\"\")")
		PrintTree("")
		fmt.Println("\nprintTree(\"1\")")
		PrintTree("1")
		fmt.Println("\nprintTree(\"11\")")
		PrintTree("11")
		fmt.Println("\nprintTree(\"111\")")
		PrintTree("111")
		fmt.Println("\nprintTree(\"abcdefghijklmn\")")
		PrintTree("abcdefghijklmn")
		fmt.Println("\nprintTree(\"aaaa,bbbb,cccccc\")")
		PrintTree("aaaa,bbbb,cccccc")
		fmt.Println("\nprintTree(\"1,#,2,#,#,3,4\")")
		PrintTree("1,#,2,#,#,3,4")
		fmt.Println("\nprintTree(\"1,2,3,4,5,#,#,6,7,8,1,#,#,#,#,#,#,2,3,4,5,6,7,8,9,10,11,12,13,14,15\")")
		PrintTree("1,2,3,4,5,#,#,6,7,8,1,#,#,#,#,#,#,2,3,4,5,6,7,8,9,10,11,12,13,14,15")
		fmt.Println("===== printTreeStringTest end ======")
	})
	fmt.Println()
	t.Run("slice", func(t *testing.T) {
		fmt.Println("=== printTreeSliceTest begin ====")
		fmt.Println("printTree(\"\")")
		PrintTree(make([]string, 0))
		fmt.Println("\nprintTree(\"1\")")
		PrintTree([]string{"1"})
		fmt.Println("\nprintTree(\"11\")")
		PrintTree([]string{"11"})
		fmt.Println("\nprintTree(\"111\")")
		PrintTree([]string{"111"})
		fmt.Println("\nprintTree(\"abcdefghijklmn\")")
		PrintTree([]string{"abcdefghijklmn"})
		fmt.Println("\nprintTree(\"aaaa,bbbb,cccccc\")")
		PrintTree([]string{"aaaa", "bbbb", "cccccc"})
		fmt.Println("\nprintTree(\"1,#,2,#,#,3,4\")")
		PrintTree([]string{"1", "#", "2", "#", "#", "3", "4"})
		fmt.Println("\nprintTree(\"1,2,3,4,5,#,#,6,7,8,1,#,#,#,#,#,#,2,3,4,5,6,7,8,9,10,11,12,13,14,15\")")
		PrintTree([]string{"1", "2", "3", "4", "5", "#", "#", "6", "7", "8", "1", "#", "#", "#", "#", "#", "#", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"})
		fmt.Println("===== printTreeSliceTest end ======")
	})
}

func TestPrintTreeLevelOrder(t *testing.T) {
	t.Run("stringLevelOrder", func(t *testing.T) {
		fmt.Println("\n=== printTreeLevelOrderStringTest begin ===")
		fmt.Println("\nprintTreeLevelOrder(\"\")")
		PrintTreeLevelOrder("")
		fmt.Println("\nprintTreeLevelOrder(\"1,2,3\")")
		PrintTreeLevelOrder("1,2,3")
		fmt.Println("\nprintTreeLevelOrder(\"root,left-child-with-long-name,right-child-with-long-name\")")
		PrintTreeLevelOrder("root,left-child-with-long-name,right-child-with-long-name")
		fmt.Println("\nprintTreeLevelOrder(\"1,2,3,#,#,4\")")
		PrintTreeLevelOrder("1,2,3,#,#,4")
		fmt.Println("\nprintTreeLevelOrder(\"1,#,#,2,3\")")
		PrintTreeLevelOrder("1,#,#,2,3")
		fmt.Println("\nprintTreeLevelOrder(\"1,#,2,#,#,#,3,#,#,#,#,#,#,#,4\")")
		PrintTreeLevelOrder("1,#,2,#,#,#,3,#,#,#,#,#,#,#,4")
		fmt.Println("=== printTreeLevelOrderStringTest end ===")
	})
	t.Run("sliceLevelOrder", func(t *testing.T) {
		fmt.Println("\n=== printTreeLevelOrderSliceTest begin ===")
		fmt.Println("\nprintTreeLevelOrder(\"\")")
		PrintTreeLevelOrder(make([]string, 0))
		fmt.Println("\nprintTreeLevelOrder(\"1,2,3\")")
		PrintTreeLevelOrder([]string{"1", "2", "3"})
		fmt.Println("\nprintTreeLevelOrder(\"root,left-child-with-long-name,right-child-with-long-name\")")
		PrintTreeLevelOrder([]string{"root", "left-child-with-long-name", "right-child-with-long-name"})
		fmt.Println("\nprintTreeLevelOrder(\"1,2,3,#,#,4\")")
		PrintTreeLevelOrder([]string{"1", "2", "3", "#", "#", "4"})
		fmt.Println("\nprintTreeLevelOrder(\"1,#,#,2,3\")")
		PrintTreeLevelOrder([]string{"1", "#", "#", "2", "3"})
		fmt.Println("\nprintTreeLevelOrder(\"1,#,2,#,#,#,3,#,#,#,#,#,#,#,4\")")
		PrintTreeLevelOrder([]string{"1", "#", "2", "#", "#", "#", "3", "#", "#", "#", "#", "#", "#", "#", "4"})
		fmt.Println("=== printTreeLevelOrderSliceTest end ===")
	})
}
