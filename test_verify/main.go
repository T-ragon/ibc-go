package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	clienttypes "github.com/T-ragon/ibc-go/v9/modules/core/02-client/types"
	"github.com/T-ragon/ibc-go/v9/modules/core/04-channel/types"
	commitmenttypes "github.com/T-ragon/ibc-go/v9/modules/core/23-commitment/types"
	host "github.com/T-ragon/ibc-go/v9/modules/core/24-host"
	_ "github.com/T-ragon/ibc-go/v9/modules/light-clients/05-aggrelite"
	aggrelite "github.com/T-ragon/ibc-go/v9/modules/light-clients/05-aggrelite"
	ics23 "github.com/cosmos/ics23/go"
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
	packets := []*types.Packet{
		{Sequence: 1,
			SourcePort:         "blog",
			SourceChannel:      "channel-4",
			DestinationPort:    "blog",
			DestinationChannel: "channel-5",
			Data: []byte{18, 59, 10, 1, 49, 18, 7, 104, 101, 108, 108, 111, 32, 49, 26, 45, 99, 111, 115, 109, 111, 115,
				49, 100, 104, 106, 99, 109, 109, 122, 113, 100, 119, 118, 121, 113, 99, 108, 48, 104, 48, 57, 102, 114,
				117, 57, 104, 97, 120, 112, 99, 115, 109, 117, 121, 121, 119, 122, 104, 48, 52},
			TimeoutHeight: clienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: 1721323200856042000},
		{Sequence: 2,
			SourcePort:         "blog",
			SourceChannel:      "channel-4",
			DestinationPort:    "blog",
			DestinationChannel: "channel-5",
			Data: []byte{18, 59, 10, 1, 50, 18, 7, 104, 101, 108, 108, 111, 32, 50, 26, 45, 99, 111, 115, 109, 111, 115,
				49, 107, 55, 99, 108, 104, 110, 97, 57, 119, 103, 112, 110, 103, 97, 119, 109, 115, 56, 110, 103, 52, 118,
				112, 99, 101, 97, 57, 120, 109, 51, 99, 53, 103, 113, 117, 120, 101, 108},
			TimeoutHeight: clienttypes.Height{
				RevisionNumber: 0,
				RevisionHeight: 0,
			},
			TimeoutTimestamp: 1721323200856042000},
		{Sequence: 3,
			SourcePort:         "blog",
			SourceChannel:      "channel-4",
			DestinationPort:    "blog",
			DestinationChannel: "channel-5",
			Data: []byte{18, 59, 10, 1, 51, 18, 7, 104, 101, 108, 108, 111, 32, 51, 26, 45, 99, 111, 115, 109, 111, 115,
				49, 56, 48, 113, 56, 57, 110, 121, 122, 48, 51, 112, 106, 50, 116, 119, 120, 110, 114, 102, 100, 107, 110,
				52, 115, 103, 52, 55, 100, 109, 115, 103, 50, 121, 118, 53, 109, 101, 56},
			TimeoutHeight: clienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: 1721323200856042000},
		{Sequence: 4,
			SourcePort:         "blog",
			SourceChannel:      "channel-4",
			DestinationPort:    "blog",
			DestinationChannel: "channel-5",
			Data: []byte{18, 59, 10, 1, 52, 18, 7, 104, 101, 108, 108, 111, 32, 52, 26, 45, 99, 111, 115, 109, 111, 115,
				49, 100, 114, 116, 110, 100, 57, 54, 108, 100, 52, 99, 113, 115, 113, 117, 114, 50, 117, 48, 109, 112,
				106, 55, 109, 104, 57, 109, 100, 97, 99, 119, 112, 103, 106, 97, 108, 113, 115},
			TimeoutHeight: clienttypes.Height{
				RevisionHeight: 0,
				RevisionNumber: 0,
			},
			TimeoutTimestamp: 1721323200856042000},
	}

	packets_leaf_number := []uint64{12, 11, 12, 12}

	proofMeta1_1 := types.ProofMeta{
		HashValue: []byte{41, 140, 2, 24, 89, 47, 191, 239, 182, 64, 201, 183, 13, 96, 145, 233, 62, 41, 220, 8, 242, 56,
			228, 207, 220, 200, 61, 86, 87, 190, 158, 58},
		PathInnerOp: &types.InnerOp{
			Hash: types.HashOp_SHA256,
			Prefix: []byte{2, 4, 174, 2, 32, 89, 198, 220, 26, 61, 155, 75, 157, 49, 143, 158, 6, 204, 137, 99, 229, 45,
				58, 165, 63, 221, 20, 196, 129, 198, 52, 8, 43, 33, 94, 25, 168, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	proofMeta1_2 := types.ProofMeta{
		HashValue: []byte{253, 131, 153, 104, 115, 227, 234, 48, 234, 64, 128, 158, 96, 227, 55, 107, 35, 97, 87, 221, 212,
			234, 208, 21, 254, 132, 3, 89, 64, 224, 184, 11},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{2, 4, 174, 2, 32},
			Suffix: []byte{32, 23, 202, 40, 112, 252, 50, 152, 76, 171, 175, 18, 116, 184, 108, 214, 160, 47, 53, 170, 155, 43, 121, 185, 158, 175, 228, 195, 32, 133, 195, 24, 181},
		},
		RealValue: []byte{},
	}
	subProof1 := types.SubProof{
		Number:        12,
		ProofMetaList: []*types.ProofMeta{&proofMeta1_1, &proofMeta1_2},
	}

	proofMeta2_1 := types.ProofMeta{
		HashValue: []byte{197, 125, 187, 4, 183, 161, 40, 100, 241, 37, 169, 35, 63, 186, 4, 248, 185, 144, 109, 202, 100,
			117, 244, 196, 78, 39, 234, 124, 236, 125, 224, 230},
		PathInnerOp: &types.InnerOp{
			Hash: types.HashOp_SHA256,
			Prefix: []byte{4, 8, 174, 2, 32, 162, 92, 236, 189, 37, 157, 132, 68, 15, 98, 141, 76, 213, 170, 9, 34, 163,
				182, 199, 66, 90, 25, 166, 22, 24, 117, 193, 29, 178, 153, 128, 202, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	proofMeta2_2 := types.ProofMeta{
		HashValue: []byte{78, 21, 150, 248, 17, 8, 179, 115, 94, 143, 97, 66, 79, 60, 255, 179, 14, 224, 207, 221, 61, 114,
			10, 151, 212, 163, 253, 25, 30, 151, 6, 100},
		PathInnerOp: &types.InnerOp{
			Hash: types.HashOp_SHA256,
			Prefix: []byte{4, 6, 174, 2, 32, 40, 105, 157, 199, 222, 39, 59, 93, 140, 188, 175, 80, 140, 154, 111, 241, 51,
				4, 121, 145, 168, 156, 167, 113, 159, 78, 211, 192, 82, 195, 39, 233, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}
	// 创建 SubProof 对象
	subProof2 := types.SubProof{
		Number:        11,
		ProofMetaList: []*types.ProofMeta{&proofMeta2_1, &proofMeta2_2},
	}

	proofMeta3_1 := types.ProofMeta{
		HashValue: []byte{78, 226, 193, 112, 194, 142, 113, 172, 220, 202, 197, 124, 94, 253, 75, 61, 243, 116, 52, 177, 72, 149, 164, 91, 200, 93, 213, 4, 115, 116, 39, 106},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{6, 16, 248, 58, 32, 205, 53, 160, 74, 48, 104, 39, 128, 9, 55, 59, 49, 91, 246, 19, 66, 5, 9, 183, 176, 183, 216, 92, 228, 203, 223, 0, 65, 4, 152, 237, 24, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 对象
	subProof3 := types.SubProof{
		Number:        16,
		ProofMetaList: []*types.ProofMeta{&proofMeta3_1},
	}

	proofMeta4_1 := types.ProofMeta{
		HashValue: []byte{30, 166, 25, 0, 40, 2, 149, 181, 29, 41, 53, 171, 253, 33, 204, 28, 211, 83, 61, 160, 117, 27, 7, 179, 199, 229, 29, 59, 43, 24, 87, 176},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{8, 32, 248, 58, 32, 135, 45, 231, 83, 233, 226, 143, 12, 241, 242, 11, 195, 216, 82, 77, 169, 72, 11, 181, 117, 222, 248, 180, 223, 201, 223, 66, 233, 95, 140, 170, 22, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 对象
	subProof4 := types.SubProof{
		Number:        15,
		ProofMetaList: []*types.ProofMeta{&proofMeta4_1},
	}

	proofMeta5_1 := types.ProofMeta{
		HashValue: []byte{113, 68, 236, 11, 35, 45, 77, 44, 34, 106, 145, 253, 55, 79, 174, 198, 200, 205, 217, 173, 245, 125, 149, 160, 105, 252, 247, 85, 252, 50, 101, 218},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{10, 44, 248, 58, 32, 240, 38, 44, 7, 35, 110, 108, 240, 156, 41, 153, 83, 93, 16, 59, 247, 245, 100, 187, 175, 195, 6, 3, 26, 143, 92, 21, 13, 57, 32, 168, 52, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 8 对象
	subProof5 := types.SubProof{
		Number:        14,
		ProofMetaList: []*types.ProofMeta{&proofMeta5_1},
	}

	// 创建 SubProof 7 的 ProofMeta 对象
	proofMeta6_1 := types.ProofMeta{
		HashValue: []byte{192, 67, 90, 117, 203, 94, 82, 24, 59, 91, 75, 145, 16, 60, 90, 144, 76, 189, 66, 167, 153, 132, 239, 177, 12, 119, 33, 49, 14, 213, 200, 10},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{12, 68, 248, 58, 32},
			Suffix: []byte{32, 208, 3, 171, 125, 43, 200, 156, 99, 107, 196, 124, 213, 1, 214, 95, 4, 248, 159, 4, 213, 82, 134, 47, 159, 221, 204, 88, 186, 131, 254, 222, 168},
		},
	}

	// 创建 SubProof 7 对象
	subProof6 := types.SubProof{
		Number:        13,
		ProofMetaList: []*types.ProofMeta{&proofMeta6_1},
	}

	proofMeta7_1 := types.ProofMeta{
		HashValue: []byte{61, 157, 248, 209, 149, 26, 76, 36, 241, 73, 91, 219, 245, 174, 218, 174, 196, 39, 12, 253, 242, 162, 209, 21, 240, 137, 215, 54, 72, 168, 115, 92},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{14, 134, 1, 248, 58, 32},
			Suffix: []byte{32, 213, 238, 23, 207, 205, 200, 110, 10, 38, 203, 76, 115, 153, 76, 79, 134, 8, 76, 68, 162, 125, 215, 93, 117, 29, 26, 92, 72, 174, 138, 49, 152},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 6 对象
	subProof7 := types.SubProof{
		Number:        19,
		ProofMetaList: []*types.ProofMeta{&proofMeta7_1},
	}

	// 创建 SubProof 5 的 ProofMeta 对象
	proofMeta8_1 := types.ProofMeta{
		HashValue: []byte{14, 250, 238, 88, 166, 58, 145, 120, 37, 207, 192, 131, 107, 224, 23, 116, 145, 69, 42, 53, 128, 134, 39, 75, 157, 186, 224, 245, 40, 198, 255, 16},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{16, 250, 1, 248, 58, 32, 244, 0, 123, 39, 101, 75, 114, 23, 201, 219, 119, 206, 61, 202, 92, 231, 155, 187, 174, 207, 99, 17, 186, 206, 12, 142, 185, 242, 68, 187, 19, 13, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 5 对象
	subProof8 := types.SubProof{
		Number:        20,
		ProofMetaList: []*types.ProofMeta{&proofMeta8_1},
	}

	// 创建 SubProof 4 的 ProofMeta 对象
	proofMeta9_1 := types.ProofMeta{
		HashValue: []byte{66, 22, 76, 206, 34, 181, 217, 170, 222, 142, 241, 33, 217, 43, 162, 197, 255, 84, 36, 250, 191, 43, 161, 194, 246, 215, 248, 240, 95, 57, 116, 212},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{18, 128, 3, 248, 58, 32, 190, 165, 31, 61, 27, 68, 70, 67, 119, 208, 229, 77, 161, 225, 49, 230, 50, 74, 222, 228, 204, 16, 179, 90, 133, 211, 21, 159, 214, 144, 88, 181, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 4 对象
	subProof9 := types.SubProof{
		Number:        10,
		ProofMetaList: []*types.ProofMeta{&proofMeta9_1},
	}

	proofMeta10_1 := types.ProofMeta{
		HashValue: []byte{133, 16, 146, 114, 64, 142, 144, 176, 105, 3, 150, 76, 207, 94, 224, 172, 239, 228, 129, 8, 199, 148, 110, 249, 63, 72, 118, 126, 177, 17, 6, 133},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{20, 222, 4, 248, 58, 32, 95, 31, 132, 95, 42, 140, 197, 210, 246, 200, 5, 202, 64, 173, 147, 245, 55, 120, 226, 243, 14, 120, 9, 77, 74, 239, 26, 161, 240, 66, 34, 187, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 3 对象
	subProof10 := types.SubProof{
		Number:        9,
		ProofMetaList: []*types.ProofMeta{&proofMeta10_1},
	}

	// 创建 SubProof 2 的 ProofMeta 对象
	proofMeta11_1 := types.ProofMeta{
		HashValue: []byte{14, 50, 149, 172, 112, 22, 88, 158, 79, 87, 81, 184, 173, 153, 144, 114, 243, 195, 17, 226, 233, 218, 217, 1, 237, 30, 77, 23, 116, 250, 154, 195},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{22, 250, 9, 248, 58, 32, 112, 202, 20, 141, 141, 92, 243, 64, 184, 73, 204, 156, 100, 239, 145, 152, 213, 157, 155, 230, 229, 133, 102, 7, 34, 204, 158, 90, 77, 155, 101, 253, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 2 对象
	subProof11 := types.SubProof{
		Number:        8,
		ProofMetaList: []*types.ProofMeta{&proofMeta11_1},
	}

	// 创建 SubProof 1 的 ProofMeta 对象
	proofMeta12_1 := types.ProofMeta{
		HashValue: []byte{189, 201, 35, 2, 85, 71, 150, 133, 184, 1, 78, 137, 108, 254, 139, 72, 87, 142, 227, 214, 154, 102, 218, 191, 163, 94, 91, 49, 160, 181, 102, 112},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{24, 252, 22, 248, 58, 32, 6, 59, 230, 169, 127, 97, 21, 131, 207, 1, 109, 21, 106, 114, 167, 114, 49, 224, 60, 195, 224, 48, 126, 25, 99, 165, 40, 23, 231, 123, 178, 28, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	// 创建 SubProof 1 对象
	subProof12 := types.SubProof{
		Number:        7,
		ProofMetaList: []*types.ProofMeta{&proofMeta12_1},
	}

	proofMeta13_1 := types.ProofMeta{
		HashValue: []byte{37, 1, 18, 29, 27, 56, 65, 236, 13, 4, 248, 104, 104, 175, 231, 16, 127, 47, 120, 134, 211, 249, 223, 207, 204, 197, 0, 139, 208, 14, 22, 1},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{26, 192, 32, 248, 58, 32},
			Suffix: []byte{32, 243, 234, 195, 118, 150, 60, 229, 252, 152, 16, 227, 109, 7, 54, 107, 130, 51, 103, 166, 84, 38, 254, 7, 31, 113, 146, 26, 85, 161, 236, 104, 97},
		},
		RealValue: []byte{},
	}

	subProof13 := types.SubProof{
		Number:        6,
		ProofMetaList: []*types.ProofMeta{&proofMeta13_1},
	}

	proofMeta14_1 := types.ProofMeta{
		HashValue: []byte{103, 65, 214, 23, 128, 241, 53, 25, 235, 148, 248, 106, 239, 115, 121, 194, 190, 123, 54, 68, 118, 14, 30, 242, 213, 85, 81, 17, 84, 93, 144, 94},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{28, 162, 70, 248, 58, 32, 99, 135, 49, 94, 85, 115, 128, 219, 120, 157, 148, 111, 64, 52, 112, 207, 208, 219, 152, 42, 203, 74, 162, 105, 21, 72, 96, 222, 115, 211, 131, 97, 32},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	subProof14 := types.SubProof{
		Number:        5,
		ProofMetaList: []*types.ProofMeta{&proofMeta14_1},
	}

	proofMeta15_1 := types.ProofMeta{
		HashValue: []byte{71, 234, 120, 187, 209, 229, 214, 193, 171, 33, 148, 201, 120, 103, 69, 39, 87, 128, 45, 98, 235, 62, 33, 122, 212, 240, 220, 28, 189, 148, 108, 88},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{1},
			Suffix: []byte{44, 216, 181, 7, 0, 149, 5, 70, 24, 10, 217, 121, 19, 90, 135, 8, 194, 234, 32, 152, 255, 246, 173, 227, 27, 126, 64, 235, 93, 207, 124, 5},
		},
		RealValue: []byte{79, 143, 253, 82, 114, 229, 188, 202, 67, 140, 16, 75, 113, 183, 48, 135, 244, 172, 233, 140, 237, 195, 211, 77, 133, 210, 243, 127, 94, 131, 58, 75},
	}

	subProof15 := types.SubProof{
		Number:        4,
		ProofMetaList: []*types.ProofMeta{&proofMeta15_1},
	}

	proofMeta16_1 := types.ProofMeta{
		HashValue: []byte{207, 169, 198, 13, 136, 189, 158, 180, 52, 164, 247, 194, 145, 87, 64, 53, 85, 95, 97, 115, 245, 130, 179, 207, 209, 172, 159, 185, 200, 81, 128, 39},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{1, 249, 196, 146, 97, 14, 105, 159, 202, 191, 145, 140, 248, 64, 60, 103, 3, 44, 203, 36, 82, 237, 114, 204, 175, 3, 71, 96, 213, 151, 130, 112, 29},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	subProof16 := types.SubProof{
		Number:        3,
		ProofMetaList: []*types.ProofMeta{&proofMeta16_1},
	}

	proofMeta17_1 := types.ProofMeta{
		HashValue: []byte{235, 154, 134, 87, 17, 149, 35, 196, 68, 65, 53, 89, 199, 97, 34, 161, 148, 214, 209, 248, 20, 151, 60, 243, 55, 186, 53, 39, 37, 189, 151, 186},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{1, 237, 171, 2, 114, 73, 14, 148, 81, 204, 240, 237, 178, 128, 65, 244, 175, 196, 99, 164, 89, 150, 251, 217, 17, 99, 212, 144, 151, 97, 115, 89, 235},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	subProof17 := types.SubProof{
		Number:        2,
		ProofMetaList: []*types.ProofMeta{&proofMeta17_1},
	}

	proofMeta18_1 := types.ProofMeta{
		HashValue: []byte{182, 220, 29, 53, 192, 95, 102, 33, 99, 73, 226, 65, 72, 87, 144, 61, 143, 218, 232, 154, 44, 202, 74, 117, 158, 211, 139, 67, 137, 93, 46, 45},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{1, 242, 237, 175, 150, 190, 126, 23, 172, 205, 226, 249, 74, 133, 63, 33, 86, 133, 136, 157, 245, 103, 54, 210, 14, 114, 56, 4, 231, 253, 17, 202, 206},
			Suffix: []byte{},
		},
		RealValue: []byte{},
	}

	subProof18 := types.SubProof{
		Number:        1,
		ProofMetaList: []*types.ProofMeta{&proofMeta18_1},
	}

	proofMeta19_1 := types.ProofMeta{
		HashValue: []byte{148, 164, 225, 89, 22, 247, 255, 62, 43, 144, 40, 118, 37, 166, 44, 75, 5, 231, 220, 68, 116, 213, 173, 174, 146, 169, 59, 171, 215, 198, 157, 186},
		PathInnerOp: &types.InnerOp{
			Hash:   types.HashOp_SHA256,
			Prefix: []byte{1},
			Suffix: []byte{2, 87, 213, 178, 16, 124, 45, 205, 47, 83, 210, 127, 33, 78, 11, 113, 211, 93, 221, 40, 193, 34, 180, 34, 107, 36, 225, 25, 122, 179, 25, 166},
		},
		RealValue: []byte{},
	}

	subProof19 := types.SubProof{
		Number:        0,
		ProofMetaList: []*types.ProofMeta{&proofMeta19_1},
	}
	subProofs := []*types.SubProof{&subProof1, &subProof2, &subProof3, &subProof4, &subProof5, &subProof6, &subProof7, &subProof8,
		&subProof9, &subProof10, &subProof11, &subProof12, &subProof13, &subProof14, &subProof15, &subProof16, &subProof17, &subProof18, &subProof19}

	signer := "cosmos1uzfl8hhkjytqzdkfzdnj2f04mhlpvzdjlyy6gq"
	height := clienttypes.Height{
		RevisionNumber: 0,
		RevisionHeight: 121,
	}

	msgAggregatePacket := types.MsgAggregatePacket{
		Packets:           packets,
		PacketsLeafNumber: packets_leaf_number,
		Proof:             subProofs,
		Signer:            signer,
		ProofHeight:       height,
	}

	commits := make([][]byte, len(packets))
	for i, packet := range packets {
		commits[i] = types.MainCommitPacket(packet)
	}

	leafOp1 := ics23.LeafOp{
		Hash:         ics23.HashOp_SHA256,
		PrehashKey:   ics23.HashOp_NO_HASH,
		PrehashValue: ics23.HashOp_SHA256,
		Length:       ics23.LengthOp_VAR_PROTO,
		Prefix:       []byte{0, 2, 174, 2},
	}

	leafOp2 := ics23.LeafOp{
		Hash:         ics23.HashOp_SHA256,
		PrehashKey:   ics23.HashOp_NO_HASH,
		PrehashValue: ics23.HashOp_SHA256,
		Length:       ics23.LengthOp_VAR_PROTO,
		Prefix:       []byte{0, 2, 174, 2},
	}

	leafOp3 := ics23.LeafOp{
		Hash:         ics23.HashOp_SHA256,
		PrehashKey:   ics23.HashOp_NO_HASH,
		PrehashValue: ics23.HashOp_SHA256,
		Length:       ics23.LengthOp_VAR_PROTO,
		Prefix:       []byte{0, 2, 174, 2},
	}

	_ = []ics23.LeafOp{leafOp1, leafOp2, leafOp3}

	keyArr := make([][]byte, len(packets))
	for i := 0; i < len(packets); i++ {
		packet := packets[i]
		merklePath := commitmenttypes.NewMerklePath(host.PacketCommitmentPath(packet.SourcePort, packet.SourceChannel, packet.Sequence))
		merklePath = commitmenttypes.NewMerklePath(append([]string{"ibc"}, merklePath.KeyPath...)...)
		fmt.Println(merklePath)
		keyArr[i], _ = merklePath.GetKey(uint64(len(merklePath.KeyPath) - 1 - 0))
		fmt.Printf("%d\n", keyArr[i])
	}

	root := []byte{79, 0, 204, 97, 39, 212, 96, 216, 162, 91, 190, 85, 214, 3, 177, 75, 67, 159, 144, 105, 240, 20, 72, 16, 196, 190, 78, 206, 23, 74, 44, 113}
	err, verified := aggrelite.MainVerifyAggregateProof(
		msgAggregatePacket.PacketsLeafNumber,
		commits,
		keyArr,
		leafOp1,
		subProofs,
		root)
	if err != nil {
		panic(err)
	}
	if verified {

		fmt.Println("Verified Successfully!")
	} else {
		fmt.Println("Verified Failure!")
	}

	subProofMap := make(map[uint64]*types.SubProof)
	for _, sub := range subProofs {
		subProofMap[sub.Number] = sub
	}

	bts := make([][]byte, len(subProofs))
	for i := 0; i < len(subProofs); i++ {
		bts[i], _ = subProofs[i].Marshal()
		fmt.Println(hex.EncodeToString(bts[i]))
	}

	// unmarshal
	sps := make([]*types.SubProof, len(subProofs))
	for i := 0; i < len(subProofs); i++ {
		sps[i] = &types.SubProof{}
		err := sps[i].Unmarshal(bts[i])
		if err != nil {
			return
		}
		fmt.Printf("Hash Value%d\n", sps[i].ProofMetaList[0].HashValue)
	}
}
