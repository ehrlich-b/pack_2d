package pack_2d

import "fmt"
import "strconv"
import "sort"

type Block struct {
	X, Y, Width, Height, Id int
}

type Packer2d struct {
	blocks []Block
}

func (p *Packer2d) AddBlock(b Block) {
	p.blocks = append(p.blocks, b)
}

func (p *Packer2d) AddNewBlock(width, height, id int) {
	p.blocks = append(p.blocks, Block{0, 0, width, height, id})
}

func (p Packer2d) GetBlocks() []Block {
	return p.blocks
}

func (p Packer2d) Pack() ([]Block, int, int) {
	blocks := p.blocks
	sort.Stable(byWidth(blocks))
	sort.Stable(byHeight(blocks))
	rootNode := &Node{false, 0, 0, 0, 0, nil, nil}

	for i := 0; i < len(blocks); i++ {
		block := &blocks[i]
		var found *Node
		found = nil
		for found == nil {
			found = getNode(rootNode, block.Height, block.Width)
			if found == nil {
				rootNode = growRootNode(rootNode, block.Height, block.Width)
			}
		}
		found.used = true
		block.X = found.x
		block.Y = found.y
		splitNode(found, block.Height, block.Width)
	}
	return blocks, rootNode.width, rootNode.height
}

type Node struct {
	used                bool
	x, y, width, height int
	down, right         *Node
}

//Sorting helpers
type byHeight []Block

func (a byHeight) Len() int           { return len(a) }
func (a byHeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byHeight) Less(i, j int) bool { return a[i].Height > a[j].Height }

type byWidth []Block

func (a byWidth) Len() int           { return len(a) }
func (a byWidth) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWidth) Less(i, j int) bool { return a[i].Width > a[j].Width }

func (node Node) printNode() {
	fmt.Print("used: ")
	fmt.Println(node.used)
	fmt.Println("x: " + strconv.Itoa(node.x))
	fmt.Println("y: " + strconv.Itoa(node.y))
	fmt.Println("width: " + strconv.Itoa(node.width))
	fmt.Println("height: " + strconv.Itoa(node.height))
	fmt.Println("--------")
	if node.right != nil {
		node.right.printNode()
	}
	if node.down != nil {
		node.down.printNode()
	}
}

func PrintBlocks(blocks []Block) {
	boundsX := 0
	boundsY := 0
	for _, val := range blocks {
		if val.Width+val.X > boundsX {
			boundsX = val.Width + val.X
		}
		if val.Height+val.Y > boundsY {
			boundsY = val.Height + val.Y
		}
	}
	etch := make([][]int, boundsY)
	for i := 0; i < len(etch); i++ {
		etch[i] = make([]int, boundsX)
	}
	for _, val := range blocks {
		etch[val.Y][val.X] = 3
		etch[val.Y+val.Height-1][val.X] = 3
		etch[val.Y][val.X+val.Width-1] = 3
		etch[val.Y+val.Height-1][val.X+val.Width-1] = 3
		for i := val.Y + 1; i < val.Y+val.Height-1; i++ {
			etch[i][val.X] = 2
		}
		for i := val.Y + 1; i < val.Y+val.Height-1; i++ {
			etch[i][val.X+val.Width-1] = 2
		}
		for i := val.X + 1; i < val.X+val.Width-1; i++ {
			etch[val.Y][i] = 1
		}
		for i := val.X + 1; i < val.X+val.Width-1; i++ {
			etch[val.Y+val.Height-1][i] = 1
		}
	}
	for i := 0; i < len(etch); i++ {
		for j := 0; j < len(etch[i]); j++ {
			switch {
			case etch[i][j] == 3:
				fmt.Print("+")
			case etch[i][j] == 2:
				fmt.Print("|")
			case etch[i][j] == 1:
				fmt.Print("-")
			case true:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func getNode(root *Node, height int, width int) *Node {
	if root == nil {
		return nil
	}
	if root.used {
		right := getNode(root.right, height, width)
		if right != nil {
			return right
		}
		down := getNode(root.down, height, width)
		return down
	} else if root.width >= width && root.height >= height {
		return root
	}
	return nil
}

func growRight(root *Node, height, width int) *Node {
	return &Node{
		true,
		0,
		0,
		root.width + width,
		root.height,
		root,
		&Node{false, root.width, 0, width, root.height, nil, nil},
	}
}

func growDown(root *Node, height, width int) *Node {
	return &Node{
		true,
		0,
		0,
		root.width,
		root.height + height,
		&Node{false, 0, root.height, root.width, height, nil, nil},
		root,
	}
}

func growRootNode(root *Node, height int, width int) *Node {
	if root.width > root.height {
		return growDown(root, height, width)
	} else {
		return growRight(root, height, width)
	}
}

func splitNode(node *Node, height, width int) {
	node.down = &Node{false, node.x, node.y + height, node.width, node.height - height, nil, nil}
	node.right = &Node{false, node.x + width, node.y, node.width - width, height, nil, nil}
}
