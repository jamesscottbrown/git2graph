package main

import (
	"testing"
)

func validateColumns(t *testing.T, expectedColumns []int, data []map[string]interface{}) {
	for idx, row := range data {
		if row["column"] != expectedColumns[idx] {
			t.Fail()
			t.Logf("Id: %s, Expected column: %d, Actual column: %d", row["id"], expectedColumns[idx], row["column"])
		}
	}
}

func validatePaths(t *testing.T, expectedPaths []map[string]Path, data []map[string]interface{}) {
	for nodeIdx, node := range data {
		for _, parentId := range node["parents"].([]string) {
			if len(expectedPaths)-1 < nodeIdx {
				continue
			}
			if len(node["parentsPaths"].(map[string]Path)[parentId].Path) != len(expectedPaths[nodeIdx][parentId].Path) {
				t.Fail()
				t.Logf("Id: %s, Expected nb paths: %d, Actual nb paths: %d", node["id"], len(expectedPaths[nodeIdx][parentId].Path), len(node["parentsPaths"].(map[string]Path)[parentId].Path))
				t.Logf("Id: %s, Expected: %v, Actual: %v", node["id"], expectedPaths[nodeIdx][parentId], node["parentsPaths"].(map[string]Path)[parentId])
				return
			}
			for pathIdx, pathNode := range node["parentsPaths"].(map[string]Path)[parentId].Path {
				if pathNode != expectedPaths[nodeIdx][parentId].Path[pathIdx] {
					t.Fail()
					t.Logf("Id: %s, Expected path: %d, Actual path: %d", node["id"], expectedPaths[nodeIdx][parentId].Path[pathIdx], pathNode)
					t.Logf("Id: %s, Expected: %v, Actual: %v", node["id"], expectedPaths[nodeIdx][parentId].Path, node["parentsPaths"].(map[string]Path)[parentId].Path)
				}
			}
		}
	}
}

func validateColors(t *testing.T, expectedPaths []map[string]Path, data []map[string]interface{}) {
	for nodeIdx, node := range data {
		for _, parentId := range node["parents"].([]string) {
			if len(expectedPaths)-1 < nodeIdx {
				continue
			}
			if expectedPaths[nodeIdx][parentId].Color != node["parentsPaths"].(map[string]Path)[parentId].Color {
				t.Logf("Id: %s, Expected: %v, Actual: %v", node["id"], expectedPaths[nodeIdx][parentId].Color, node["parentsPaths"].(map[string]Path)[parentId].Color)
				t.Fail()
			}
		}
	}
}

var customColors []string = []string{"color1", "color2", "color3", "color4", "color5", "color6"}

// 1
// |
// 2
// |
// 3
func Test1(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 0, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{0, 1, 0}}, "color1"},
		},
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 1, 0}, Point{0, 2, 0}}, "color1"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// | 2
// |/
// 3
func Test2(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
		},
		map[string]Path{
			"3": Path{"3", []Point{Point{1, 1, 0}, Point{1, 2, 1}, Point{0, 2, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// |/
// 3
func Test3(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"3": Path{"3", []Point{Point{1, 1, 0}, Point{1, 2, 1}, Point{0, 2, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// 3 |
// |\|
// | |\
// | | 4
// | |/
// |/
// 5
func Test4(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 2, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 1, 0}, Point{1, 4, 1}, Point{0, 4, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{0, 4, 0}}, "color1"},
			"4": Path{"4", []Point{Point{0, 2, 0}, Point{2, 2, 2}, Point{2, 3, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{2, 3, 0}, Point{2, 4, 1}, Point{0, 4, 0}}, "color3"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// | 2
// | | 3
// | |/
// |/
// 4
// | 5
// |/
// 6
func Test5(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 2, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 0, 0}, Point{0, 3, 0}}, "color1"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 3, 1}, Point{0, 3, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{2, 2, 0}, Point{2, 3, 1}, Point{0, 3, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 3, 0}, Point{0, 5, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 4, 0}, Point{1, 5, 1}, Point{0, 5, 0}}, "color4"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// |/|
// 3 |
// | 4
// |/
// 5
func Test6(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"3": Path{"3", []Point{Point{1, 1, 0}, Point{0, 1, 3}, Point{0, 2, 0}}, "color1"},
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 3, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 3, 0}, Point{1, 4, 1}, Point{0, 4, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// 3 |\
// | 4 |
// | |/
// |/|
// 5 |
// |/
// 6
func Test7(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"4", "5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 1, 0, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 3, 0}}, "color2"},
			"5": Path{"5", []Point{Point{1, 1, 0}, Point{2, 1, 2}, Point{2, 4, 1}, Point{0, 4, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 3, 0}, Point{1, 5, 1}, Point{0, 5, 0}}, "color2"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 4, 0}, Point{0, 5, 0}}, "color1"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// 3 |
// |\|
// | |\
// |/ /
// 4 |
// | 5
// |/
// 6
func Test8(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"4", "5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 3, 1}, Point{0, 3, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 2, 0}, Point{0, 3, 0}}, "color1"},
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{2, 2, 2}, Point{2, 3, 1}, Point{1, 3, 0}, Point{1, 4, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 3, 0}, Point{0, 5, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 4, 0}, Point{1, 5, 1}, Point{0, 5, 0}}, "color3"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// 3 |
// |\|
// | |\
// 4 | |
// |\| |
// | |\|
// | | |\
// |/ / /
// 5 | |
// | | 6
// | 7 |
// | |/
// |/
// 8
func Test9(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"4", "7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5", "6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "8", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 0, 0, 2, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 1, 0}, Point{1, 4, 1}, Point{0, 4, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 2, 0}, Point{0, 3, 0}}, "color1"},
			"7": Path{"7", []Point{Point{0, 2, 0}, Point{2, 2, 2}, Point{2, 4, 1}, Point{1, 4, 0}, Point{1, 6, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 3, 0}, Point{0, 4, 0}}, "color1"},
			"6": Path{"6", []Point{Point{0, 3, 0}, Point{3, 3, 2}, Point{3, 4, 1}, Point{2, 4, 0}, Point{2, 5, 0}}, "color4"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{0, 4, 0}, Point{0, 7, 0}}, "color1"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{2, 5, 0}, Point{2, 7, 1}, Point{0, 7, 0}}, "color4"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{1, 6, 0}, Point{1, 7, 1}, Point{0, 7, 0}}, "color3"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// | |\
// | | 3
// 4 | |
// |\| |
// | |\|
// | | |\
// | |/ /
// | 5 |
// | |\|
// | | 6
// | |/
// | 7
// |/
// 8
func Test10(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"4", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"5", "3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"8", "6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"7", "6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "8", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 2, 0, 1, 2, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 0, 0}, Point{0, 3, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 1, 0}, Point{1, 4, 0}}, "color2"},
			"3": Path{"3", []Point{Point{1, 1, 0}, Point{2, 1, 2}, Point{2, 2, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{2, 2, 0}, Point{2, 4, 1}, Point{1, 4, 0}}, "color3"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{0, 3, 0}, Point{0, 7, 0}}, "color1"},
			"6": Path{"6", []Point{Point{0, 3, 0}, Point{3, 3, 2}, Point{3, 4, 1}, Point{2, 4, 0}, Point{2, 5, 0}}, "color4"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 4, 0}, Point{2, 4, 2}, Point{2, 5, 0}}, "color4"},
			"7": Path{"7", []Point{Point{1, 4, 0}, Point{1, 6, 0}}, "color2"},
		},
		map[string]Path{
			"7": Path{"7", []Point{Point{2, 5, 0}, Point{2, 6, 1}, Point{1, 6, 0}}, "color4"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{1, 6, 0}, Point{1, 7, 1}, Point{0, 7, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |-2
// 3 |
// |-4
// 5 |
// |/
// 6
func Test11(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5", "6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 1, 0, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
		},
		map[string]Path{
			"3": Path{"3", []Point{Point{1, 1, 0}, Point{0, 1, 3}, Point{0, 2, 0}}, "color1"},
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 3, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 3, 0}, Point{0, 3, 3}, Point{0, 4, 0}}, "color1"},
			"6": Path{"6", []Point{Point{1, 3, 0}, Point{1, 5, 1}, Point{0, 5, 0}}, "color2"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 4, 0}, Point{0, 5, 0}}, "color1"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |-2
// 3 |
// |\|
// | |\
// | | 4
// | |/
// |/|
// 5 |
// | 6
// |/
// 7
func Test12(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3", "6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 2, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// |\
// | 2
// | |
// | 3
// 4 |\
// |\| |
// | 5 |
// | |\|
// | 6 |\
// 7 | | |
// | | |/
// | |/|
// |/| |
// 8 | |
// | | 9
// | |/
// |/
// 10
func Test13(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"4", "2"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5", "9"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"7", "5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6", "8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"10"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "8", "parents": []string{"10"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "9", "parents": []string{"10"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "10", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 1, 0, 1, 1, 0, 0, 2, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 0, 0}, Point{0, 3, 0}}, "color1"},
			"2": Path{"2", []Point{Point{0, 0, 0}, Point{1, 0, 2}, Point{1, 1, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// | 2
// 3 |\
// |\| |
// | |\|
// | | 4
// | |/
// | 5
// 6 |\
// |\| |
// | |\|
// | | 7
// | |/
// |/
// 8
func Test14(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"5", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"6", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"8", "7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"8", "7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "8", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 2, 1, 0, 2, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{1, 1, 0}, Point{1, 4, 0}}, "color2"},
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{2, 1, 2}, Point{2, 3, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 2, 0}, Point{0, 5, 0}}, "color1"},
			"4": Path{"4", []Point{Point{0, 2, 0}, Point{2, 2, 2}, Point{2, 3, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{2, 3, 0}, Point{2, 4, 1}, Point{1, 4, 0}}, "color3"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{1, 4, 0}, Point{1, 7, 1}, Point{0, 7, 0}}, "color2"},
			"7": Path{"7", []Point{Point{1, 4, 0}, Point{2, 4, 2}, Point{2, 6, 0}}, "color4"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// | 2
// 3 |\
// |\| |
// | |\|
// | | 4
// 5 | |
// |\| |
// | |\|
// | | |\
// | |/ /
// | 6 |
// | |\|
// | | 7
// | |/
// |/
// 8
func Test15(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"3"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"6", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5", "4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"8", "7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"8", "7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{"8"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "8", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 0, 2, 0, 1, 2, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"3": Path{"3", []Point{Point{0, 0, 0}, Point{0, 2, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 1, 0}, Point{1, 5, 0}}, "color2"},
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{2, 1, 2}, Point{2, 3, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 2, 0}, Point{0, 4, 0}}, "color1"},
			"4": Path{"4", []Point{Point{0, 2, 0}, Point{2, 2, 2}, Point{2, 3, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{2, 3, 0}, Point{2, 5, 1}, Point{1, 5, 0}}, "color3"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{0, 4, 0}, Point{0, 7, 0}}, "color1"},
			"7": Path{"7", []Point{Point{0, 4, 0}, Point{3, 4, 2}, Point{3, 5, 1}, Point{2, 5, 0}, Point{2, 6, 0}}, "color4"},
		},
		map[string]Path{
			"8": Path{"8", []Point{Point{1, 5, 0}, Point{1, 7, 1}, Point{0, 7, 0}}, "color2"},
			"7": Path{"7", []Point{Point{1, 5, 0}, Point{2, 5, 2}, Point{2, 6, 0}}, "color4"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

// 1
// | 2
// | | 3
// | |/ 4
// |/| /
// 5 |/
// | 6
// |/
// 7
func Test16(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{"7"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "7", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 2, 3, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 0, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 1, 0}, Point{1, 5, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{2, 2, 0}, Point{2, 4, 1}, Point{0, 4, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{3, 3, 0}, Point{3, 4, 1}, Point{2, 4, 0}, Point{2, 5, 1}, Point{1, 5, 0}}, "color4"},
		},
		map[string]Path{
			"7": Path{"7", []Point{Point{0, 4, 0}, Point{0, 6, 0}}, "color1"},
		},
		map[string]Path{
			"7": Path{"7", []Point{Point{1, 5, 0}, Point{1, 6, 1}, Point{0, 6, 0}}, "color2"},
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

func Test17(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "0", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 2, 3, 0, 0, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 0, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 4, 1}, Point{0, 4, 0}}, "color2"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{2, 2, 0}, Point{2, 4, 1}, Point{0, 4, 0}}, "color3"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{3, 3, 0}, Point{3, 4, 1}, Point{1, 4, 0}, Point{1, 6, 1}, Point{0, 6, 0}}, "color4"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{0, 4, 0}, Point{0, 5, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 5, 0}, Point{0, 6, 0}}, "color1"},
		},
		map[string]Path{},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}

func Test18(t *testing.T) {
	// Initial input
	inputNodes := make([]map[string]interface{}, 0)
	inputNodes = append(inputNodes, map[string]interface{}{"id": "0", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "1", "parents": []string{"4"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "2", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "3", "parents": []string{"5"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "4", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "5", "parents": []string{"6"}})
	inputNodes = append(inputNodes, map[string]interface{}{"id": "6", "parents": []string{}})

	out, _ := buildTree(inputNodes, customColors)

	// Expected output
	expectedColumns := []int{0, 1, 2, 3, 0, 1, 0}

	expectedPaths := []map[string]Path{
		map[string]Path{
			"4": Path{"4", []Point{Point{0, 0, 0}, Point{0, 4, 0}}, "color1"},
		},
		map[string]Path{
			"4": Path{"4", []Point{Point{1, 1, 0}, Point{1, 4, 1}, Point{0, 4, 0}}, "color2"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{2, 2, 0}, Point{2, 4, 1}, Point{1, 4, 0}, Point{1, 5, 0}}, "color3"},
		},
		map[string]Path{
			"5": Path{"5", []Point{Point{3, 3, 0}, Point{3, 4, 1}, Point{2, 4, 0}, Point{2, 5, 1}, Point{1, 5, 0}}, "color4"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{0, 4, 0}, Point{0, 6, 0}}, "color1"},
		},
		map[string]Path{
			"6": Path{"6", []Point{Point{1, 5, 0}, Point{1, 6, 1}, Point{0, 6, 0}}, "color3"},
		},
		map[string]Path{
		},
	}

	// Validation
	validateColumns(t, expectedColumns, out)
	validatePaths(t, expectedPaths, out)
	validateColors(t, expectedPaths, out)
}
