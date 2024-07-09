package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Node represents a node in the Merkle tree
type Node struct {
	Left  *Node
	Right *Node
	Data  []byte
}

type ProofMeta struct {
	M1 []byte
	M2 []byte
}

type SubProof struct {
	Number        uint64
	ProofMetaList *[]ProofMeta
}

// NewNode creates a new node from left and right children
func NewNode(left, right *Node) *Node {
	data := append(left.Data, right.Data...)
	hash := sha256.Sum256(data)
	return &Node{Left: left, Right: right, Data: hash[:]}
}

// NewLeafNode creates a new leaf node
func NewLeafNode(data []byte) *Node {
	hash := sha256.Sum256(data)
	return &Node{Data: hash[:]}
}

// BuildMerkleTree builds a Merkle tree from a list of leaf nodes
func BuildMerkleTree(leafNodes []*Node) *Node {
	if len(leafNodes) == 1 {
		return leafNodes[0]
	}

	var parentNodes []*Node
	for i := 0; i < len(leafNodes); i += 2 {
		if i+1 < len(leafNodes) {
			parentNodes = append(parentNodes, NewNode(leafNodes[i], leafNodes[i+1]))
		} else {
			// If odd number of nodes, duplicate the last one
			parentNodes = append(parentNodes, NewNode(leafNodes[i], leafNodes[i]))
		}
	}

	return BuildMerkleTree(parentNodes)
}

// PrintTree prints the Merkle tree in a tree structure
func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}
	indent := strings.Repeat("  ", level)
	fmt.Printf("%sNode Hash: %s\n", indent, hex.EncodeToString(node.Data))
	PrintTree(node.Left, level+1)
	PrintTree(node.Right, level+1)
}

// CollectProofs collects ProofMeta for each level of the Merkle tree
func CollectProofs(node *Node, level int, proofs *[]SubProof) {
	if node == nil {
		return
	}
	if len(*proofs) <= level {
		*proofs = append(*proofs, SubProof{
			Number:        uint64(level),
			ProofMetaList: &[]ProofMeta{}})
	}
	if node.Left != nil && node.Right != nil {
		*(*proofs)[level].ProofMetaList = append(*(*proofs)[level].ProofMetaList, ProofMeta{
			M1: node.Left.Data,
			M2: node.Right.Data,
		})
	}
	CollectProofs(node.Left, level+1, proofs)
	CollectProofs(node.Right, level+1, proofs)
}

// PrintSubProofs prints the SubProofs in a readable format
func PrintSubProofs(proofs []SubProof) {
	for _, proof := range proofs {
		fmt.Printf("Level %d:\n", proof.Number)
		for _, meta := range *proof.ProofMetaList {
			fmt.Printf("  (M1: %s, M2: %s\n)", hex.EncodeToString(meta.M1), hex.EncodeToString(meta.M2))
		}
	}
}

func main() {
	// Example data for the six leaf nodes
	leafData := [][]byte{
		[]byte("Leaf 1"),
		[]byte("Leaf 2"),
		[]byte("Leaf 3"),
		[]byte("Leaf 4"),
		[]byte("Leaf 5"),
		[]byte("Leaf 6"),
		[]byte("Leaf 7"),
		[]byte("Leaf 8"),
	}

	// Create leaf nodes
	var leafNodes []*Node
	for _, data := range leafData {
		leafNodes = append(leafNodes, NewLeafNode(data))
	}

	// Build the Merkle tree
	root := BuildMerkleTree(leafNodes)

	// Print the Merkle tree nodes' hashes in tree structure
	fmt.Println("Merkle Tree Nodes' Hashes:")
	PrintTree(root, 0)

	// Print the Merkle root
	fmt.Printf("Merkle Root: %s\n", hex.EncodeToString(root.Data))

	// Collect proofs
	var proofs []SubProof
	CollectProofs(root, 0, &proofs)

	//Print the Subproofs
	PrintSubProofs(proofs)
}
