package com.github.kfarnung.adventofcode.aoc2023.utilities;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.List;

public class InputUtils {
    public static List<String> readLinesFromResources(Object myClass, String resourceName) throws IOException {
        try (InputStream stream = myClass.getClass().getClassLoader().getResourceAsStream(resourceName);
                BufferedReader reader = new BufferedReader(new InputStreamReader(stream))) {
            return reader.lines().toList();
        }
    }

    public static List<String> readLinesFromFile(String fileName) throws IOException {
        try (FileReader file = new FileReader(fileName);
                BufferedReader reader = new BufferedReader(file)) {
            return reader.lines().toList();
        }
    }
}
