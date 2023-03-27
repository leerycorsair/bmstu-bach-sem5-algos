package dict

type Item struct {
	Key   string
	Value int
}

type Dict []Item

type Segment struct {
	Key   byte
	Value Dict
}

type SegmentedDict []Segment
