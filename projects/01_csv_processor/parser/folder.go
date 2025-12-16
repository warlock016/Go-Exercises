package parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

type Header struct {
	File string
	Cols []string
	Set  map[string]struct{}
}
type Cluster struct {
	Rep     Header
	Members []Header
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), " ")
	return s
}
func makeHeader(file string, cols []string) Header {
	set := make(map[string]struct{}, len(cols))
	out := make([]string, 0, len(cols))

	for _, c := range cols {
		c = normalize(c)
		out = append(out, c)
		set[c] = struct{}{}
	}

	return Header{File: file, Cols: out, Set: set}
}
func jaccard(a, b map[string]struct{}) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 1
	}
	inter := 0
	if len(a) > len(b) {
		a, b = b, a
	}
	for k := range a {
		if _, ok := b[k]; ok {
			inter++
		}
	}
	union := len(a) + len(b) - inter
	return float64(inter) / float64(union)
}
func clusterHeaders(headers []Header, threshold float64) []Cluster {
	clusters := []Cluster{}
	for _, h := range headers {
		bestIdx := -1
		bestSim := float64(0)
		for i := range clusters {
			sim := jaccard(h.Set, clusters[i].Rep.Set)
			if sim > bestSim {
				bestSim = sim
				bestIdx = i
			}
		}

		if bestIdx >= 0 && bestSim >= threshold {
			clusters[bestIdx].Members = append(clusters[bestIdx].Members, h)
		} else {
			clusters = append(clusters, Cluster{Rep: h, Members: []Header{h}})
		}
	}
	return clusters
}
func setDiff(a, b map[string]struct{}) (onlyA, onlyB []string) {
	for k := range a {
		if _, ok := b[k]; !ok {
			onlyA = append(onlyA, k)
		}
	}
	for k := range b {
		if _, ok := a[k]; !ok {
			onlyB = append(onlyB, k)
		}
	}
	sort.Strings(onlyA)
	sort.Strings(onlyB)
	return onlyA, onlyB
}

func testMain() {

	dirPath := "./testdata/"
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("failed to open %s: %v", dirPath, err)
	}
	if len(dirPath) == 0 {
		log.Fatalf("empty: %s", dirPath)
	}
	sort.Slice(dir, func(i, j int) bool { return dir[i].Name() < dir[j].Name() })

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get $PWD: %v", err)
	}

	headers := make([]Header, 0, len(dir))
	ref := Header{
		File: "",
		Cols: make([]string, 0),
		Set:  make(map[string]struct{}),
	}
	for i, file := range dir {
		filepath := path.Join(pwd, dirPath, file.Name())
		f, err := os.Open(filepath)
		if err != nil {
			log.Fatalf("failed to open: %s: %v", file.Name(), err)
		}
		defer f.Close()

		csv := csv.NewReader(f)
		header, err := csv.Read()
		if err != nil {
			log.Fatalf("%s failed to parse csv record: %v", filepath, err)
		}
		if len(header) == 0 {
			log.Fatalf("returned empty header")
		}

		currentHeader := makeHeader(filepath, header)
		headers = append(headers, currentHeader)

		if i == 0 {
			ref.File = filepath
			ref.Cols = header
			for _, val := range header {
				ref.Set[val] = struct{}{}
			}
		} else {
			onlyRef, onlyFile := setDiff(ref.Set, currentHeader.Set)
			if len(onlyFile) != 0 && len(onlyRef) != 0 {
				continue
			}

			if len(onlyRef) > 0 {
				fmt.Printf("missing vs ref: %v", onlyRef)
			}
			if len(onlyFile) > 0 {
				fmt.Printf("extra vs ref: %v", onlyFile)
			}
		}
	}

	cluster := clusterHeaders(headers, 0.9)

	fmt.Println("Number of clusters:", len(cluster))
	fmt.Println("first cluster:", cluster[0].Rep.Set)
}
