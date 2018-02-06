#include <emscripten.h>
#include <stdint.h>

uint32_t EMSCRIPTEN_KEEPALIVE calculateNext(
    uint32_t previousValue,
    uint32_t factor,
    uint32_t requiredMultiple)
{
    do
    {
        previousValue = ((uint64_t)previousValue * factor) % 2147483647;
    }
    while (requiredMultiple != 1 && previousValue % requiredMultiple != 0);

    return previousValue;
}