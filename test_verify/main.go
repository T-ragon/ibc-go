package main

import (
	"bytes"
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
	return &Node{Data: hash[:], Left: nil, Right: nil}
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

func doHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func VerifyLeafWithProof(leafData []byte, subProofs []SubProof, root []byte) bool {
	currentHash := doHash(leafData)

	//迭代计算上一层的哈希值
	for _, subProof := range subProofs {
		found := false
		for _, proofMeta := range *subProof.ProofMetaList {
			//check if the current hash matches one of the meta hashes
			if bytes.Equal(currentHash, proofMeta.M1) || bytes.Equal(currentHash, proofMeta.M2) {
				combinedDate := append(proofMeta.M1, proofMeta.M2...)
				currentHash = doHash(combinedDate)
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return bytes.Equal(currentHash, root)
}

// PrintTree prints the Merkle tree in a tree structure
func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}
	indent := strings.Repeat("  ", level)
	fmt.Printf("%slevel:%d, Node Hash: %s\n", indent, level, hex.EncodeToString(node.Data))
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
		*(*proofs)[level-1].ProofMetaList = append(*(*proofs)[level-1].ProofMetaList, ProofMeta{
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
			fmt.Printf("  (M1: %s, M2: %s)\n", hex.EncodeToString(meta.M1), hex.EncodeToString(meta.M2))
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
	rootHash := root.Data
	//// Print the Merkle tree nodes' hashes in tree structure
	//fmt.Println("Merkle Tree Nodes' Hashes:")
	PrintTree(root, 0)
	// Collect proofs
	var proofs []SubProof
	CollectProofs(root, 1, &proofs)
	var subProofs []SubProof
	for i := 3; i > 0; i-- {
		subProofs = append(subProofs, proofs[i-1])
	}
	PrintSubProofs(subProofs)

	if VerifyLeafWithProof(leafData[0], subProofs, rootHash) {
		fmt.Println("Leaf 1 Verified")
	} else {
		fmt.Println("Leaf 1 Verified failed!")
	}
}
