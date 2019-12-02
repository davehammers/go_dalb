/*
Copyright (c) 2019 Dave Hammers
*/
package cors

import (
	"io/ioutil"
)

const (
	DEBUG_PRIVATE_KEY = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCpLfLTdhUg0+Ja
/QPqi61hvIWk+/Q7sJCp+lH2PypYY8dlE1vRsAAClP8RtGuNGR6QDxdwFaULZb28
pWKVf4zahClo9QKBIfS0VJDpZAW0snkqWUHQWVyGd3vRk/5JN3ZdPhNqo4M2npI/
1n7Ag0iCnbzUC8u77rYeH+I7u2l09EoNSHJKYsWKDhxgN/cp4+Ai51lW3gTohK9z
fpqbieaojMEbupTAZASg2dVEVrOFSNnpzGMqisQ9SWqKuqqrrUVaI0tW2+Lvz+VI
kynnNANjX0Za2b/ELXNWuhwE/zN/3iAYiZHbFAkCboxWQFaitGg1xWkZzUvyD+54
hdGbGwuFAgMBAAECggEAbEs0V/YLWkMnbvTs79NPPfPuXIjHcvn38EaC9qzCT2g3
d9Tnfpc5um1jxRKHkf9VWAPBDgdc6anLxZjcPTQzlqDo3P2RB9YTjOdhB3T2Tg+8
jcYq4dKB7rVHNgWUzYtKIi+dQDLAyFLC7UhBRLwEy09rUxTl6jvIqgngyDmULL80
TRMyt4slXrY6qCh4GUkCIQ8Ppn6BGX6gpQB1M3tYSa+ktEXLvEpdGxMP5zN+6dIH
zp8KoxwT9PdgLeGvGVsnHErWuF4ckQMuCXOkoGuNsHOJ+yDCgOhG9jO7mKqL/gcG
IXkCDbfQVej+8ipuG3adv+YE2y9shIjOrwIglS6JdQKBgQDc1VW4AyN0kRgeb0Uq
QojXQuC3audRds0gaCX3DkVMTj8SLkrJbMuy7o3A6HFqP81BA2B3tEmZqHwl0XLb
PKZFAHBOKFyYL3SFvsleZc2XTFyFIu3M5xtCajNagUzFlhlvriO3xwvgYAFiB8mS
TPEC0yNLsGifvYqCfatrM9/7fwKBgQDEHtoqu+/8djn4T/3sZkWThS9fpZDhJp+C
LszhplhoECLE5xASlt91PwKrFelQJZeb/wC2REpJ+LUlGxolfriYbj1x/2/XWuUC
9iyXdUR68WgZnoWueGn8jPJweQLJiutVmnisp2q0k08Rvleea/WP7YMarTUAvDrC
FEVF6ZOK+wKBgEtVh6F3iJ/aY4T2ZnztJ6tviCNqF7FlusJkZRcryh4mz0NPgXnH
YJIQ6VC2uwII8+dK1JzhZv5BSODJ28nIndwKM0WmZTgRcEmz9RP73K/Rf/p3GPJ+
oID+o7grRdpdwx2jJVIrV/TaK36as3vyPYG+L1tBud9MlLTBVDoE/1LXAoGBAJpS
9LW/4V414fPlbhhBeepVWSvYaqLg849LKGk/rj1kxRdQAzO9iOUHyh+6RDeO/TgB
dxv1rMd8b35dTzvF9Zfs12kG6Yj6u06TTA96dYKQx6uxM9xQYIYcmwGqF2lkaT6Z
KPAjZTHm4MdDkkrVBQxWh13MgfKsK0hhDwmufS1VAoGAeZDFxrLIJzHnfi4jZLvJ
MmZ5/Fd3BuFCofkvBa+5knbVLnfd3j0WgOCnlqsmkkySs4hPqNqSxwgCcTEgKbs2
sck11qgOkgCgj3/1Eq2MJwEGYa4ZMTboS9oYFcTetkzjZY1hzHVyddc0P7+33QBI
xPZJuZL0jEg1QQjKee7hSdg=
-----END PRIVATE KEY-----
`

	DEBUG_CERT = `-----BEGIN CERTIFICATE-----
MIIDTTCCAjWgAwIBAgIJAIiec1gDB+DaMA0GCSqGSIb3DQEBCwUAMD0xCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybWlhMRkwFwYDVQQHDBBFeHRyZW1lIE5l
dHdvcmtzMB4XDTE4MTEzMDE0MTcxNFoXDTM4MTEyNTE0MTcxNFowPTELMAkGA1UE
BhMCVVMxEzARBgNVBAgMCkNhbGlmb3JtaWExGTAXBgNVBAcMEEV4dHJlbWUgTmV0
d29ya3MwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCpLfLTdhUg0+Ja
/QPqi61hvIWk+/Q7sJCp+lH2PypYY8dlE1vRsAAClP8RtGuNGR6QDxdwFaULZb28
pWKVf4zahClo9QKBIfS0VJDpZAW0snkqWUHQWVyGd3vRk/5JN3ZdPhNqo4M2npI/
1n7Ag0iCnbzUC8u77rYeH+I7u2l09EoNSHJKYsWKDhxgN/cp4+Ai51lW3gTohK9z
fpqbieaojMEbupTAZASg2dVEVrOFSNnpzGMqisQ9SWqKuqqrrUVaI0tW2+Lvz+VI
kynnNANjX0Za2b/ELXNWuhwE/zN/3iAYiZHbFAkCboxWQFaitGg1xWkZzUvyD+54
hdGbGwuFAgMBAAGjUDBOMB0GA1UdDgQWBBQRTRgF9iZEtAR9mmHTZMMoXJFoGDAf
BgNVHSMEGDAWgBQRTRgF9iZEtAR9mmHTZMMoXJFoGDAMBgNVHRMEBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQBNXeHEH9SkzkJab6yyQl5UEejckVwi8hjXG4VkG5Rb
mnNTaffd3KNw5HkTiO/twJslkIJ8JHfGB17D7wnSSxxcEkSSGUtQJJ0xaOcBaOc2
0C0Gw/++5y8qDsASiHn+wHiHlf5M1yCdjEMbhvi2UrzlUrjH8JSEEPMiBNwLmQAy
HcVpBwP1LP3Pk9+TmqW0rL8N9H6uYc23bpI2A1HyNP5vlpXpzaaPiDTxAefU7NCZ
dvbeZn9feJHvsO/91iTvAvTo1MSt+t5isFW1vhpHccl7DfdEy7GJJp/UT3Wd5tBg
SH0Z6+2uTrqXZbdTR5mGineNt1551gVkX+oB6X46eO9E
-----END CERTIFICATE-----
`
)

func CertKeys() (pubkey string, privatekey string, err error) {
	privateKeyFile, err := ioutil.TempFile("", "")
	if err == nil {
		privateKeyFile.Write([]byte(DEBUG_PRIVATE_KEY))
		privatekey = privateKeyFile.Name()
	} else {
		return
	}

	certFile, err := ioutil.TempFile("", "")
	if err == nil {
		certFile.Write([]byte(DEBUG_CERT))
		pubkey = certFile.Name()
	}
	return
}
