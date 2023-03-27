package dict

import (
	"fmt"
	"os"
)

func EfficiencyCheck(d Dict) error {
	file, err := os.Create("comp.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString("n,simple,bin,segmented\n")
	for i, item := range d {
		_, simpleCount, _ := BruteS(d, item.Key)
		_, binCount, _ := BinS(d, item.Key)
		_, segmentedCount, _ := SegmS(d, item.Key)
		fmt.Fprintf(file, "%d,%d,%d,%d\n", i+1, simpleCount, binCount, segmentedCount)
	}
	_, simpleCount, _ := BruteS(d, "bmstu")
	_, binCount, _ := BinS(d, "bmstu")
	_, segmentedCount, _ := SegmS(d, "bmstu")
	fmt.Fprintf(file, "%d,%d,%d,%d\n", len(d)+1, simpleCount, binCount, segmentedCount)
	return err
}
