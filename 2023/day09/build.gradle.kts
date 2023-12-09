plugins {
    id("com.github.kfarnung.adventofcode.aoc2023.java-application-conventions")
}

dependencies {
    implementation(project(":utilities"))
}

application {
    mainClass.set("com.github.kfarnung.adventofcode.aoc2023.day09.App")
}
