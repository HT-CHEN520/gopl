package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rootIndex := indexOf(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: BuildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

func (root *TreeNode) InOrder(Visit func(node *TreeNode)) {
	if root == nil {
		return
	}
	root.Left.InOrder(Visit)
	Visit(root)
	root.Right.InOrder(Visit)
}

func (root *TreeNode) PreOrder(Visit func(node *TreeNode)) {
	if root == nil {
		return
	}
	Visit(root)
	root.Left.PreOrder(Visit)
	root.Right.PreOrder(Visit)
}

func indexOf(order []int, val int) int {
	for i, v := range order {
		if val == v {
			return i
		}
	}
	return -1
}

func Serialize(root *TreeNode) string {
	inArray, preArray := []int{}, []int{}

	root.InOrder(func(node *TreeNode) {
		inArray = append(inArray, node.Val)
	})
	root.PreOrder(func(node *TreeNode) {
		preArray = append(preArray, node.Val)
	})

	builder := strings.Builder{}
	for _, v := range inArray {
		builder.WriteString(fmt.Sprintf("%d,", v))
	}
	builder.WriteString("||")
	for _, v := range preArray {
		builder.WriteString(fmt.Sprintf("%d,", v))
	}

	return builder.String()
}

// Deserializes your encoded data to tree.
func Deserialize(data string) *TreeNode {
	inArray, preArray := []int{}, []int{}
	strList := strings.SplitN(data, "||", 2)
	inStr := strList[0]
	preStr := strList[1]
	inStrArray := strings.Split(inStr, ",")
	for _, v := range inStrArray {
		if len(v) < 0 {
			continue
		}

		iv, _ := strconv.Atoi(v)
		inArray = append(inArray, iv)
	}

	preStrArray := strings.Split(preStr, ",")
	for _, v := range preStrArray {
		if len(v) < 0 {
			continue
		}

		iv, _ := strconv.Atoi(v)
		preArray = append(preArray, iv)
	}

	return BuildTree(preArray, inArray)
}

func main() {
	tree := BuildTree([]int{3, 4, 1, 2, 5}, []int{1, 4, 2, 3, 5})
	s := Serialize(tree)
	fmt.Println(s)
	tree1 := Deserialize(s)

	fmt.Println(tree1.Val)
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	counts := regexp.MustCompile("red").FindAllStringIndex(input, -1)
	n, _ := fmt.Println(len(counts))
	if n >= 2 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
} */

/* import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return []int{}
	}
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	total := m * n
	res := make([]int, total)
	row, column := 0, 0
	dir := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	dirIndex := 0
	for i := 0; i < total; i++ {
		res[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+dir[dirIndex][0], column+dir[dirIndex][1]
		if nextRow >= m || nextRow < 0 || nextColumn >= n || nextColumn < 0 || visited[nextRow][nextColumn] {
			dirIndex = (dirIndex + 1) % 4
		}
		row += dir[dirIndex][0]
		column += dir[dirIndex][1]
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanln(&arr[i][j])
		}
	}
	ans := spiralOrder(arr)


	ans[0] = ans[n]
	for i := n; i > 0; i-- {
		ans[i] = ans[i-1]
	}
	res := spiralMatrix(n, n, ans)
	fmt.Println(res)
}

var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralMatrix(n int, m int, res []int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	for x, y, di, i := 0, 1, 0, 0; i < len(res); i++ {
		ans[x][y] = res[i]
		d := dirs[di&3]
		if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= m || ans[xx][yy] != -1 {
			di++
			d = dirs[di&3]
		}
		x += d.x
		y += d.y
	}
	return ans
}
*/

/* package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
*/

/* type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) Serialize(root *TreeNode) string {
	// write code here
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("#,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}
func (this *Codec) Deserialize(s string) *TreeNode {
	// write code here
	sp := strings.Split(s, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
*/
