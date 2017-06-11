package repository

import (
    "bufio"
    "encoding/csv"
    "os"
    "io"
    
	. "rbacapi/types"
)

func loadKRA(fileName string) map[Keyset]map[Role]bool {
    return toKRA(loadFromCSV(fileName))
}

func toKRA(orig map[string]map[string]bool) map[Keyset]map[Role]bool {
	result := make(map[Keyset]map[Role]bool)
	for key1, _ := range orig {
		if result[Keyset(key1)] == nil {
			result[Keyset(key1)] = make(map[Role]bool)
		}
		for key2, value := range orig[key1] {
			result[Keyset(key1)][Role(key2)] = value
		}
	}
    return result
}

func loadPKA(fileName string) map[Permission]map[Keyset]bool {
    return toPKA(loadFromCSV(fileName))
}

func toPKA(orig map[string]map[string]bool) map[Permission]map[Keyset]bool {
	result := make(map[Permission]map[Keyset]bool)
	for key1, _ := range orig {
		if result[Permission(key1)] == nil {
			result[Permission(key1)] = make(map[Keyset]bool)
		}
		for key2, value := range orig[key1] {
			result[Permission(key1)][Keyset(key2)] = value
		}
	}
    return result
}

func loadPRA(fileName string) map[Permission]map[Role]bool {
    return toPRA(loadFromCSV(fileName))
}

func toPRA(orig map[string]map[string]bool) map[Permission]map[Role]bool {
	result := make(map[Permission]map[Role]bool)
	for key1, _ := range orig {
		if result[Permission(key1)] == nil {
			result[Permission(key1)] = make(map[Role]bool)
		}
		for key2, value := range orig[key1] {
			result[Permission(key1)][Role(key2)] = value
		}
	}
    return result
}

func loadRRA(fileName string) map[Role]map[Role]bool {
    return toRRA(loadFromCSV(fileName))
}

func toRRA(orig map[string]map[string]bool) map[Role]map[Role]bool {
	result := make(map[Role]map[Role]bool)
	for key1, _ := range orig {
		if result[Role(key1)] == nil {
			result[Role(key1)] = make(map[Role]bool)
		}
		for key2, value := range orig[key1] {
			result[Role(key1)][Role(key2)] = value
		}
	}
    return result
}

func loadURA(fileName string) map[User]map[Role]bool {
    return toURA(loadFromCSV(fileName))
}

func toURA(orig map[string]map[string]bool) map[User]map[Role]bool {
	result := make(map[User]map[Role]bool)
	for key1, _ := range orig {
		if result[User(key1)] == nil {
			result[User(key1)] = make(map[Role]bool)
		}
		for key2, value := range orig[key1] {
			result[User(key1)][Role(key2)] = value
		}
	}
    return result
}

func loadFromCSV(fileName string) map[string]map[string]bool {
	result := make(map[string]map[string]bool)
	
	f, _ := os.Open(fileName)
	
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		
		if result[record[0]] == nil {
			result[record[0]] = make(map[string]bool)
		}
		
		result[record[0]][record[1]] = true
	}
	
	return result
}