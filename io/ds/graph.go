package ds

import (
	"fmt"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

/***** GNode : Graph Node *****/

const keySep = "-"

type Graph struct {
	AdList map[string]map[string]int
	AdMat  map[string]int // key = v1-v2
	VMap   map[string]int
}

func NewGraph() *Graph {
	g := &Graph{
		AdList: make(map[string]map[string]int), // list is a map to easily check duplicates while adding
		AdMat:  make(map[string]int),
		VMap:   make(map[string]int),
	}
	return g
}

func (g *Graph) VList() []string {
	vlist := make([]string, 0)
	for v, _ := range g.VMap {
		vlist = append(vlist, v)
	}
	return vlist
}

func (g *Graph) AddVertex(vid string, vval int, adjWeightMap map[string]int) {
	g.VMap[vid] = vval
	g.addVtoAdList(vid, adjWeightMap)
	g.addVtoAdMat(vid, adjWeightMap)
}

func (g *Graph) FindVertexesByValue(v int) []string {
	nodeIds := make([]string, 0)
	for vid, val := range g.VMap {
		if val == v {
			nodeIds = append(nodeIds, vid)
		}
	}
	return nodeIds
}

func (g *Graph) addVtoAdList(v string, adjWeightMap map[string]int) {
	currAdj, ok := g.AdList[v]
	if !ok {
		currAdj = make(map[string]int)
	}
	for adj, wei := range adjWeightMap {
		currAdj[adj] = wei
	}
	g.AdList[v] = currAdj
}

func (g *Graph) addVtoAdMat(v string, adjWeightMap map[string]int) {
	key := ""
	for adj, wei := range adjWeightMap {
		key = fmt.Sprintf("%v%v%v", v, keySep, adj)
		g.AdMat[key] = wei
	}
}

func (g *Graph) PrintAdList() {
	for v, adj := range g.AdList {
		fmt.Printf("%v (%q) : %v => %v\n", v, g.VMap[v], len(adj), g.MapToStr(adj))
	}
}

func (g *Graph) PrintAdMat() {

	matSize := len(g.VMap) + 1
	mat := make([][]string, matSize)

	vToIdxMap := make(map[string]int)
	for i := 0; i < matSize; i++ {
		mat[i] = make([]string, matSize)
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = "0"
		}
	}

	i := 0
	for v, _ := range g.VMap {
		vToIdxMap[v] = i
		mat[0][i+1] = v
		mat[i+1][0] = v
		i++
	}

	var tokens []string
	for key, wei := range g.AdMat {
		tokens = strings.Split(key, keySep)
		mat[vToIdxMap[tokens[0]]+1][vToIdxMap[tokens[1]]+1] = fmt.Sprintf("%v", wei)
	}
	stringz.PPrint2D(mat)
}

func (g *Graph) MapToStr(m map[string]int) string {
	s := ""
	for k, _ := range m {
		s += fmt.Sprintf("%v(%q) ", k, g.VMap[k])
	}
	return s
}
