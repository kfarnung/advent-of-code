package com.github.kfarnung.adventofcode.aoc2023.day05;

import java.util.ArrayList;
import java.util.List;

public class RangeMap {
    public class SplitResult {
        private final Range mappedRange;
        private final List<Range> unmappedRanges;

        public SplitResult(Range mappedRange, List<Range> unmappedRanges) {
            this.mappedRange = mappedRange;
            this.unmappedRanges = unmappedRanges;
        }

        public Range getMappedRange() {
            return mappedRange;
        }

        public List<Range> getUnmappedRanges() {
            return unmappedRanges;
        }
    }

    private final long destination;
    private final long source;
    private final long count;

    public RangeMap(long destination, long source, long count) {
        this.destination = destination;
        this.source = source;
        this.count = count;
    }

    public boolean contains(long value) {
        return value >= source && value < source + count;
    }

    public long map(long value) {
        return destination + (value - source);
    }

    public SplitResult split(Range range) {
        long oldStart = range.getStart();
        long oldEnd = range.getEnd();
        long newStart = Math.max(source, oldStart);
        long newEnd = Math.min(source + count, oldEnd);

        if (newStart <= newEnd) {
            Range mappedRange = new Range(map(newStart), map(newEnd));

            List<Range> unmappedRanges = new ArrayList<>();
            if (oldStart < newStart) {
                unmappedRanges.add(new Range(oldStart, newStart - 1));
            }

            if (oldEnd > newEnd) {
                unmappedRanges.add(new Range(newEnd + 1, oldEnd));
            }

            return new SplitResult(mappedRange, unmappedRanges);
        }

        return new SplitResult(null, List.of(range));
    }
}
