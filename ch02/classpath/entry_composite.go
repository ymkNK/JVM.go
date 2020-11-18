package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (c CompositeEntry) String() string {
	stringArray := make([]string, len(c))
	for i, entry := range c {
		stringArray[i] = entry.String()
	}
	return strings.Join(stringArray, pathListSeparator)
}

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
