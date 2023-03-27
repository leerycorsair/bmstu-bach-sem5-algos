package dict

import (
	"bufio"
	"os"
	"regexp"
	"sort"
)

func (d Dict) getPos(key string) (int, bool) {
	for i, item := range d {
		if item.Key == key {
			return i, true
		}
	}
	return 0, false
}

func DictInit(d *Dict, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		return err
	}

	*d = nil

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := reg.ReplaceAllString(scanner.Text(), "")
		if i, exists := d.getPos(word); exists {
			(*d)[i].Value += 1
		} else if word != "" {
			(*d) = append((*d), Item{word, 1})
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Slice(*d, func(i, j int) bool { return (*d)[i].Key < (*d)[j].Key })

	return err
}

func BruteS(d Dict, searchword string) (int, int, bool) {
	var count int
	for _, item := range d {
		count++
		if item.Key == searchword {
			return item.Value, count, true
		}
	}
	return 0, count, false
}

func BinS(d Dict, searchword string) (int, int, bool) {
	var count int
	var l int
	var r int = len(d) - 1

	for r > l {
		m := (l + r) / 2

		if d[m].Key < searchword {
			count++
			l = m + 1
		} else if d[m].Key > searchword {
			count++
			r = m - 1
		} else {
			return d[m].Value, count, true
		}
	}

	if len(d) != 0 && d[l].Key == searchword {
		count++
		return d[l].Value, count, true
	} else {
		return 0, count, false
	}
}

func SegmS(d Dict, searchword string) (int, int, bool) {
	sd := DictSegmentation(d)
	var found bool
	var segment Segment
	var count int
	letter := searchword[0]
	for _, s := range sd {
		count++
		if s.Key == letter {
			segment = s
			found = true
			break
		}
	}
	if !found {
		return 0, count, false
	}
	res, countBin, found := BinS(segment.Value, searchword)
	return res, countBin, found
}

func DictSegmentation(d Dict) SegmentedDict {
	var sd SegmentedDict
	for _, item := range d {
		letter := item.Key[0]
		var found bool
		for i := 0; i < len(sd); i++ {
			if sd[i].Key == letter {
				sd[i].Value = append(sd[i].Value, item)
				found = true
				break
			}
		}

		if !found {
			var segment Segment
			segment.Key = letter
			segment.Value = append(segment.Value, item)
			sd = append(sd, segment)
		}
	}
	sort.Slice(sd, func(i, j int) bool { return len(sd[i].Value) > len(sd[j].Value) })
	return sd
}
