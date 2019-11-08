#include <stdio.h>
#include "crypto.h"

int test_main() {
    char *key = "-----BEGIN RSA PRIVATE KEY-----\n"
                "MIIJKQIBAAKCAgEAxlVkL8ePTNkijdxef2onC4ccmnvHx7d4Vym0IuIUqZ2QFNFU\n"
                "Yw/IMT31UFIqegeW5TGs9a43cEBLNNgCyTEKsouPW5ZjSo4c8CeONRdIW9GsRQk0\n"
                "fK4XzHVarm9pjPMx2Ln4VynIsWj1P3qy6S9NNIow9z0akBvgzr+pRF1HostGz/RU\n"
                "snjxg91o6kGe0gEhAuyzQDoumCS+LnRbHeNj6AsGa773dLvQZmFE9955SfMEy8m6\n"
                "UFsv9/oWiKhLdHvKIM+HTXenZ2owI6k/hrcmU03Ruap1nQAwYDpjdRY8AN/3hiAk\n"
                "EoO3gU/sx+s3MXPe/iIrZAhKGMXPtVDb2ID1f6u3b8K4tL03SSg1yaimfy77i1vG\n"
                "TpHznUBQDwWH5Tz8hXzy6kRAD2AGsvwvbi/QzIJGhtMw1KVLXm9uWTUf6e8vrwZ1\n"
                "gSnqzrnT3tBqTWoMHkLJ7oLnj46TqKeh9LPrwBsbsws1bDR4WIsXdpwzgkoTWc7L\n"
                "UYyYr4b8ACWKV6QBwKtDvATV36bnx5P51zkATLmmr+zKGtyMn0ldlpZH9YwKpW4K\n"
                "NfcqbjknbB+bQMNollHUxF/JuVgw2LGbB4sz2byYYZAoFXqNq31kFngi5nPcHi6j\n"
                "8DwgoQ001W8ksim45wRvuqJmUKWRyjJheFCswfS/ge6FVTIUOYEybucjih0CAwEA\n"
                "AQKCAgEAlvEkB2z0NYNHVfmx/Xx3bMGOVlAAEpIiRwvZKXcwTIo6vm01sRKwxDEo\n"
                "QtHVu/uMrq+ot224iXiVBAmlzSLCxnGKUCTbOkF/6pHgG621hxPC7ON9i6ofOJ7T\n"
                "vc/S38+yTwPx6bxGHicIByDFisxSELtfWrqpPaXJ6O9azknDnDTilp+X2iBLhpT6\n"
                "JNZ+Hct4KTalkSr2jRnhl488TTnirhW99EBpKfFKQLCkgZRScKJAMyw63K8ZibtO\n"
                "bQDQND2F7oSir6VxxTW+n1VOoKNAysN96rS9QBiFuKaXTOP4FJ5fTjel3GVcQlDm\n"
                "npv37G4H9xdgOIhKhCH/2zlHp7U1oHX+Qsa+GUq/c3uk1Cuw3bnt9ZvtivOfIeYQ\n"
                "I7hoz7mPbU0ijpxFKQRnZT/GmZaYWai6EDpapY9ff7X1r59tVvGMSA3NgmZuz+6t\n"
                "Kbg0uOZzJhPzD//Vny7NRqr99/f1VSbJ6YvvBLyWl9OeuQ6IPGzvOi8XaICkqIpT\n"
                "tiLvKLYu55noxhAzDrjAvOEKYA07THfS67LFZrj8qZE6lF5umnBl1XwAOH9lQhy/\n"
                "8fTDgcu5altQMKVfqr9fh/b7GIcXNBgLtfk1Pu7XGLfj3dpYWvZW5xBO7hFAbauk\n"
                "jdfUdFl1lLL3qm/QTr6p32RleGTYN0+Tc2FO3EHmTbDJoFzHPe0CggEBAPDJoJ/4\n"
                "Ya+NBrFYfNJ7ndxrjsUdUJE3Ee6l2Nv9JOOFshzmwU/l+Ze0yCcK0etnzEHms+32\n"
                "BwBePCmc1zYQc1M+7BA1Tqv+sG8Ho3f3nks/blF4V9YSY6MxH3B+ry1W2wHX9Irf\n"
                "tQEY6Dfa96xrV0VB0r9qty1f1YOlfEWJvfDwc14pcgnT2U/zIFju8xGBPDgDakiF\n"
                "5NAUlpNS0YwxILnM09TXU2oUN7F8njrftiMyiWOOIQF/UTs+f3jSy8fEPlXGa9xs\n"
                "eFMlyHrr2U4exvJyO5zmciLOhrGGopwkzdEc+BINcXiPjJiND7yrW+939VJit0Ej\n"
                "/9yA463T6iUSceMCggEBANLdImhPK510+LU+qsF/i7tcrZ0GF/E8TDj9/hc0nWzY\n"
                "HDWbRji8LFkG8mMORshHJRaCR4ROStWC7dxy2/M9ewTF8S5H3+XgoS/hvZlx/ulG\n"
                "umhWpPxEZT5cmatRD1iYzN9ADqSWu3fIrQ70RBw9PQNPronhPS244lAEocjeOqBH\n"
                "bTYwUC+ckDvZEThzAt60Idqiscr194LDgFZ4yNJ08E50a5P8TANekKYIEhuKoT0u\n"
                "YHc02pO8LlW7jQXUjJSvKyUKmnRlJ3EcNAx6/ce9mrXtG9vMhNuOqcirrllitvLQ\n"
                "aJZVYquNTggJYiQuKK+mcfZanu7hrUh2w5dvFK9i0/8CggEAfMG7O6dR1cdYBGM4\n"
                "qUXrUN1Zp7+8ksDZxbCgX7sVdd07n8XfuyoI3BWK7s+oXDP3nN2PtGeY0RQCT/03\n"
                "dIepeSRM40j7bhoUCDMI+4uMtKg03Hlh6US140P8aij5UqCB8L6Xsaye9+aTyvzk\n"
                "/qzPFs84Bn2gUx4oXoFLliv8Ae5TmCIZOAZPviDWTb3gqt0u+kaqttDI8Rb5vXNX\n"
                "py99KUd7Kfg2++tlv8w1n4Nxt2Lj1HU7nK7+w5dqLIvrkaGYOpEIKbj5zvrwmN/C\n"
                "Q7umkM+nG3A7CtW+7BQ6BHT9Pq+nyJK2jCS0UAYmdTbD95tLvFfxYwrn8rPFQ7dc\n"
                "xcB8yQKCAQBD9ng3jITvPBtJN4iL01MzMVzXxnYDD781g0/ZJOE0ircU5BYPBT95\n"
                "9k47dQeFV8Dxb04jq6RdCtUlf3O7A27aC/5/PzU//1WUfDrC8UYK4/wC0yJcGKNV\n"
                "JT12RSsgECfAMQJHNDn6EpkMv9gQDgDTR2RnFkzEptlylvuaJV5Z+IuPsqS1o82t\n"
                "LHprak5bf02GDXgmhX6gC+kaddWsV3p4nvdpfCD32QvgJ6vGarkrYf4/ja6BfV6l\n"
                "zUxXu7kP1yGdz7wWld/PihqQhzeyoD70MhcPkeykY2f/wK3yK2nx+xAqnBywVFv5\n"
                "JSUXqjT84DXNBEpDjkNunrDN50SQftb7AoIBAQDtXI9wqY3hAjVRg6bkcMsxJthz\n"
                "NRNsMnR7AHI5Cej+exU7etYIIj0d38DJryhnk1SQVvz2fI2u9ixC19A0b/5RiVAM\n"
                "/iUOA7NI9d1+Hok3s09Y3a9K+PPotXSKMQ2fFTmHjwEPimvT0d7CL5hrHc4q3jNw\n"
                "uMy/gh7qpZnQ6vmWo508oiN8+AoQUU14D6AK+vZ3vdxkmX+AehywgZgT0yw66Hfc\n"
                "fArL0LBi2hGgUuwaIWgMINcbUUhSiCtj6O0+2zJXdBk1oV+kqs0yN7YjYEq77gPC\n"
                "rchLJIcwUeB/hk1VnY2alsQQthOPW4JZkrqSSg2H9OUCMw5Nw1mt0RCl0HdZ\n"
                "-----END RSA PRIVATE KEY-----";
    char *rlt;

    rlt = NewCipherDecodeJNI(key, 1573205547,
                             "IraOAjhxuuQYuTfYQsbQ90aetg/AKYO82Ro4JTe17QZkAwRoO79cob0z2RtK1/YtoNyMFzoxgTIegLflimIQmjAnP9we7vytdS3tC50Y7TlbuquUvpZsWpGTJ+dFhY2f2817zTz9aG+diHLho23UZXrEjHmeiw3+a/Zc0GSfiC4a2Y0mSiG7TEBAuY5qaQ+xC5DI15B184qEwckbmIkEG7m0uI7arcnBrBzufkqrUM6nXongVfYigauKksx4feM0N1vCtVo+WjDAZLRaXlQuPgsXCObZsI+NX7b15FxzFD/YfQKlpFUmzbRknMO1VGkH+jymkMXZE307XKMHjHD8WRJeTR6a2mSL71BMFOkDtliaLQHFIsexofNZoxrlBE4GWmwzUubdYBWcoHEiAx6G6vajzCUlRdI3D6fhbzdPV1usDXtY9+xCivqxv6x10HB7O9gL5Z9hgA8OXzXG/+YuJAivJ3DU+5urhL2J77j30P7gh9sv1HTUbsTdFHpKvkVvGFAhTgTzaU74fbPPwQA047b/weRd1aDJucg8DvN19Ftju1XnlmcFQPKg4qiX3aKhIu7iVFF5fNlQL1KOQ74W8Q/ynSByzjbSuxOhWRUpEGy/h1MNaI9bJlXMMptLHWLjTfXWjZYF8pSmD1rBUKR66gPkqd5TpulV2aqaikbJc8Q=");

    printf("%s", rlt);

}