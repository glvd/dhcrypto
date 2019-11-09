#include <stdio.h>
#include "crypto.h"

int main() {
    char *rlt;
    rlt = CipherDecodeJNI(
            "IraOAjhxuuQYuTfYQsbQ90aetg/AKYO82Ro4JTe17QZkAwRoO79cob0z2RtK1/YtoNyMFzoxgTIegLflimIQmjAnP9we7vytdS3tC50Y7TlbuquUvpZsWpGTJ+dFhY2f2817zTz9aG+diHLho23UZXrEjHmeiw3+a/Zc0GSfiC4a2Y0mSiG7TEBAuY5qaQ+xC5DI15B184qEwckbmIkEG7m0uI7arcnBrBzufkqrUM6nXongVfYigauKksx4feM0N1vCtVo+WjDAZLRaXlQuPgsXCObZsI+NX7b15FxzFD/YfQKlpFUmzbRknMO1VGkH+jymkMXZE307XKMHjHD8WRJeTR6a2mSL71BMFOkDtliaLQHFIsexofNZoxrlBE4GWmwzUubdYBWcoHEiAx6G6vajzCUlRdI3D6fhbzdPV1usDXtY9+xCivqxv6x10HB7O9gL5Z9hgA8OXzXG/+YuJAivJ3DU+5urhL2J77j30P7gh9sv1HTUbsTdFHpKvkVvGFAhTgTzaU74fbPPwQA047b/weRd1aDJucg8DvN19Ftju1XnlmcFQPKg4qiX3aKhIu7iVFF5fNlQL1KOQ74W8Q/ynSByzjbSuxOhWRUpEGy/h1MNaI9bJlXMMptLHWLjTfXWjZYF8pSmD1rBUKR66gPkqd5TpulV2aqaikbJc8Q=",
            1573205546);
    printf("%s", rlt);
}