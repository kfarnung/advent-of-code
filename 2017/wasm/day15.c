#include <emscripten.h>
#include <stdbool.h>
#include <stdint.h>

typedef struct GeneratorData
{
    uint32_t currentValue;
    uint32_t factor;
    uint32_t requiredMultiple;
} GeneratorData;

void calculateNext(GeneratorData* data)
{
    do
    {
        data->currentValue = ((uint64_t)data->currentValue * data->factor) % 2147483647;
    }
    while (data->requiredMultiple != 1 && data->currentValue % data->requiredMultiple != 0);
}

bool areBitsEqual(uint32_t val1, uint32_t val2, uint32_t numBits) {
    uint32_t mask = (1 << numBits) - 1;
    return (val1 & mask) == (val2 & mask);
}

uint32_t EMSCRIPTEN_KEEPALIVE countMatches(
    uint32_t startingValueA,
    uint32_t startingValueB,
    uint32_t numIterations,
    uint32_t requiredMultipleA,
    uint32_t requiredMultipleB)
{
    GeneratorData generatorA = { startingValueA, 16807, requiredMultipleA };
    GeneratorData generatorB = { startingValueB, 48271, requiredMultipleB };

    uint32_t count = 0;

    for (uint32_t i = 0; i < numIterations; i++) {
        calculateNext(&generatorA);
        calculateNext(&generatorB);

        if (areBitsEqual(generatorA.currentValue, generatorB.currentValue, 16)) {
            count++;
        }
    }

    return count;
}