package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type State int

const (
	SEED_TO_SOIL State = iota
	SOIL_TO_FERT
	FERT_TO_WATER
	WATER_TO_LIGHT
	LIGHT_TO_TEMP
	TEMP_TO_HUMID
	HUMID_TO_LOCATION
)

type transformMap struct {
	dest  int
	src   int
	Range int
}

type mappingSet struct {
	seedToSoil      []transformMap
	soilToFert      []transformMap
	fertToWater     []transformMap
	waterToLight    []transformMap
	lightToTemp     []transformMap
	tempToHumid     []transformMap
	humidToLocation []transformMap
}

type Garden struct {
	Seed    []int
	Mapping mappingSet
}

func partOne(file *os.File) int {
	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}

	sc := bufio.NewScanner(file)
	var garden Garden
	var mapParseState State

	for sc.Scan() {
		txt := sc.Text()
		if txt == "" {
			continue
		}

		parseInputFile(&mapParseState, &garden, txt)
	}

	if sc.Err() != nil {
		log.Fatalln(err.Error())
	}

	destinationSet := []int{}
	// find each destination land corresponding to seed
	for _, v := range garden.Seed {
		seed := v

		destinationSet = append(destinationSet, findLocation(seed, garden.Mapping))
	}

	return slices.Min(destinationSet)
}

func findLocation(seed int, maps mappingSet) int {
	soil := lookup(seed, maps.seedToSoil)
	fertilizer := lookup(soil, maps.soilToFert)
	water := lookup(fertilizer, maps.fertToWater)
	light := lookup(water, maps.waterToLight)
	temperature := lookup(light, maps.lightToTemp)
	humid := lookup(temperature, maps.tempToHumid)
	location := lookup(humid, maps.humidToLocation)

	return location
}

func lookup(seed int, maps []transformMap) int {
	for _, mp := range maps {
		if seed > mp.src && seed <= mp.src+mp.Range {
			return seed - mp.src + mp.dest
		}
	}
	return seed
}

func parseInputFile(mapParseState *State, garden *Garden, txt string) {
	switch {

	case strings.Contains(txt, "seeds: "):
		seedStr, found := strings.CutPrefix(txt, "seeds: ")
		if !found {
			log.Fatalf("prefix not found in string \"%s\"", txt)
		}

		seeds := strings.Fields(seedStr)
		for _, seed := range seeds {
			n, err := strconv.Atoi(seed)
			if err != nil {
				log.Fatalln(err.Error())
			}
			garden.Seed = append(garden.Seed, n)
		}

	case strings.Contains(txt, "seed-to-soil map:"):
		*mapParseState = SEED_TO_SOIL

	case strings.Contains(txt, "soil-to-fertilizer map:"):
		*mapParseState = SOIL_TO_FERT

	case strings.Contains(txt, "fertilizer-to-water map:"):
		*mapParseState = FERT_TO_WATER

	case strings.Contains(txt, "water-to-light map:"):
		*mapParseState = WATER_TO_LIGHT

	case strings.Contains(txt, "light-to-temperature map:"):
		*mapParseState = LIGHT_TO_TEMP

	case strings.Contains(txt, "temperature-to-humidity map:"):
		*mapParseState = TEMP_TO_HUMID

	case strings.Contains(txt, "humidity-to-location map:"):
		*mapParseState = HUMID_TO_LOCATION

	default:

		parseMap(*mapParseState, garden, txt)
	}
}

func parseMap(state State, garden *Garden, strMap string) {
	parts := strings.Fields(strMap)

	dest, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalln(err.Error())
	}

	src, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	rangeLen, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalln(err.Error())
	}

	m := transformMap{
		dest:  dest,
		src:   src,
		Range: rangeLen,
	}

	switch state {
	case SEED_TO_SOIL:
		garden.Mapping.seedToSoil = append(garden.Mapping.seedToSoil, m)

	case SOIL_TO_FERT:
		garden.Mapping.soilToFert = append(garden.Mapping.soilToFert, m)

	case FERT_TO_WATER:
		garden.Mapping.fertToWater = append(garden.Mapping.fertToWater, m)

	case WATER_TO_LIGHT:
		garden.Mapping.waterToLight = append(garden.Mapping.waterToLight, m)

	case LIGHT_TO_TEMP:
		garden.Mapping.lightToTemp = append(garden.Mapping.lightToTemp, m)

	case TEMP_TO_HUMID:
		garden.Mapping.tempToHumid = append(garden.Mapping.tempToHumid, m)

	case HUMID_TO_LOCATION:
		garden.Mapping.humidToLocation = append(garden.Mapping.humidToLocation, m)
	}
}
