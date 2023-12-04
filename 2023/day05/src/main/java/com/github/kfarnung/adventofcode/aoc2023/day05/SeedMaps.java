package com.github.kfarnung.adventofcode.aoc2023.day05;

import java.util.ArrayList;
import java.util.List;

public class SeedMaps {
    private final List<Long> seeds;
    private final List<RangeMap> seedToSoilMap;
    private final List<RangeMap> soilToFertilizerMap;
    private final List<RangeMap> fertilizerToWaterMap;
    private final List<RangeMap> waterToLightMap;
    private final List<RangeMap> lightToTemperatureMap;
    private final List<RangeMap> temperatureToHumidityMap;
    private final List<RangeMap> humidityToLocationMap;

    public SeedMaps(List<String> lines) {
        seeds = new ArrayList<>();
        seedToSoilMap = new ArrayList<>();
        soilToFertilizerMap = new ArrayList<>();
        fertilizerToWaterMap = new ArrayList<>();
        waterToLightMap = new ArrayList<>();
        lightToTemperatureMap = new ArrayList<>();
        temperatureToHumidityMap = new ArrayList<>();
        humidityToLocationMap = new ArrayList<>();

        parse(lines);
    }

    public List<Long> getSeeds() {
        return seeds;
    }

    public long getLocationForSeed(long seed) {
        long soil = map(seedToSoilMap, seed);
        long fertilizer = map(soilToFertilizerMap, soil);
        long water = map(fertilizerToWaterMap, fertilizer);
        long light = map(waterToLightMap, water);
        long temperature = map(lightToTemperatureMap, light);
        long humidity = map(temperatureToHumidityMap, temperature);
        return map(humidityToLocationMap, humidity);
    }

    public List<Range> getLocationRangesForSeed(Range seedRange) {
        List<Range> seedRanges = List.of(seedRange);
        List<Range> soilRanges = convertRanges(seedToSoilMap, seedRanges);
        List<Range> fertilizerRanges = convertRanges(soilToFertilizerMap, soilRanges);
        List<Range> waterRanges = convertRanges(fertilizerToWaterMap, fertilizerRanges);
        List<Range> lightRanges = convertRanges(waterToLightMap, waterRanges);
        List<Range> temperatureRanges = convertRanges(lightToTemperatureMap, lightRanges);
        List<Range> humidityRanges = convertRanges(temperatureToHumidityMap, temperatureRanges);
        return convertRanges(humidityToLocationMap, humidityRanges);
    }

    private List<Range> convertRanges(List<RangeMap> rangeMaps, List<Range> ranges) {
        List<Range> unmappedRanges = new ArrayList<>(ranges);
        List<Range> mappedRanges = new ArrayList<>();

        for (RangeMap rangeMap : rangeMaps) {
            List<Range> currentRanges = new ArrayList<>(unmappedRanges);
            unmappedRanges.clear();

            for (Range currentRange : currentRanges) {
                RangeMap.SplitResult splitResult = rangeMap.split(currentRange);
                if (splitResult.getMappedRange() != null) {
                    mappedRanges.add(splitResult.getMappedRange());
                }
                unmappedRanges.addAll(splitResult.getUnmappedRanges());
            }
        }
        
        mappedRanges.addAll(unmappedRanges);
        return mappedRanges;
    }

    private void parse(List<String> lines) {
        int lineIndex = 0;
        while (lineIndex < lines.size()) {
            String line = lines.get(lineIndex);

            if (line.startsWith("seeds:")) {
                parseSeeds(line);
            } else if (line.startsWith("seed-to-soil map:")) {
                parseMap(lines, lineIndex, seedToSoilMap);
            } else if (line.startsWith("soil-to-fertilizer map:")) {
                parseMap(lines, lineIndex, soilToFertilizerMap);
            } else if (line.startsWith("fertilizer-to-water map:")) {
                parseMap(lines, lineIndex, fertilizerToWaterMap);
            } else if (line.startsWith("water-to-light map:")) {
                parseMap(lines, lineIndex, waterToLightMap);
            } else if (line.startsWith("light-to-temperature map:")) {
                parseMap(lines, lineIndex, lightToTemperatureMap);
            } else if (line.startsWith("temperature-to-humidity map:")) {
                parseMap(lines, lineIndex, temperatureToHumidityMap);
            } else if (line.startsWith("humidity-to-location map:")) {
                parseMap(lines, lineIndex, humidityToLocationMap);
            }

            lineIndex++;
        }
    }

    private void parseSeeds(String line) {
        String[] parts = line.split(" ");
        for (int i = 1; i < parts.length; i++) {
            seeds.add(Long.parseLong(parts[i]));
        }
    }

    private void parseMap(List<String> lines, int lineIndex, List<RangeMap> map) {
        lineIndex++;
        while (lineIndex < lines.size()) {
            String line = lines.get(lineIndex);
            if (line.isEmpty()) {
                break;
            }

            String[] parts = line.split(" ");
            map.add(new RangeMap(Long.parseLong(parts[0]), Long.parseLong(parts[1]), Long.parseLong(parts[2])));
            lineIndex++;
        }
    }

    private long map(List<RangeMap> map, long value) {
        for (RangeMap rangeMap : map) {
            if (rangeMap.contains(value)) {
                return rangeMap.map(value);
            }
        }
        return value;
    }
}
