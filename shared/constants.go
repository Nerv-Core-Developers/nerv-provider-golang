package shared

const (
	NotFoundPage = "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIiBkaXI9Imx0ciI+Cgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9IlVURi04Ij4KICAgIDx0aXRsZT40MDQgTk9UIEZPVU5EPC90aXRsZT4KICAgIDxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyI+CiAgICAgICAgYm9keSB7CiAgICAgICAgICAgIG1hcmdpbjogMDsKICAgICAgICB9CgogICAgICAgIC5vdXRlci1jb250YWluZXIgewogICAgICAgICAgICBwb3NpdGlvbjogYWJzb2x1dGU7CiAgICAgICAgICAgIGRpc3BsYXk6IHRhYmxlOwogICAgICAgICAgICB3aWR0aDogMTAwJTsKICAgICAgICAgICAgaGVpZ2h0OiAxMDAlOwogICAgICAgICAgICBiYWNrZ3JvdW5kOiAjZmZmOwogICAgICAgIH0KCiAgICAgICAgLmlubmVyLWNvbnRhaW5lciB7CiAgICAgICAgICAgIGRpc3BsYXk6IHRhYmxlLWNlbGw7CiAgICAgICAgICAgIHZlcnRpY2FsLWFsaWduOiBtaWRkbGU7CiAgICAgICAgICAgIHRleHQtYWxpZ246IGNlbnRlcjsKICAgICAgICB9CgogICAgICAgIC5jZW50ZXJlZC1jb250ZW50IHsKICAgICAgICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrOwogICAgICAgICAgICB0ZXh0LWFsaWduOiBjZW50ZXI7CiAgICAgICAgICAgIHBhZGRpbmc6IDIwcHg7CiAgICAgICAgICAgIGNvbG9yOiAjYmJiOwogICAgICAgICAgICBmb250LWZhbWlseTogJ1NlZ29lIFVJJywgVGFob21hLCBHZW5ldmEsIFZlcmRhbmEsIHNhbnMtc2VyaWYKICAgICAgICB9CiAgICA8L3N0eWxlPgo8L2hlYWQ+CgoKPGJvZHk+CiAgICA8ZGl2IGNsYXNzPSJvdXRlci1jb250YWluZXIiPgogICAgICAgIDxkaXYgY2xhc3M9ImlubmVyLWNvbnRhaW5lciI+CiAgICAgICAgICAgIDxkaXYgY2xhc3M9ImNlbnRlcmVkLWNvbnRlbnQiPgogICAgICAgICAgICAgICAgPGgxPjQwNDwvaDE+CiAgICAgICAgICAgICAgICA8cD5BcHAgbm90IGV4aXN0IG9yIG5vdCBhdmFpbGlhYmxlLiBQbGVhc2UgdHJ5IGFnYWluIGluIGEgZmV3IHNlY29uZC4uLjwvcD4KICAgICAgICAgICAgPC9kaXY+CiAgICAgICAgPC9kaXY+CiAgICA8L2Rpdj4KPC9ib2R5PgoKPC9odG1sPg=="
)

type DatabaseType string

// var (
// 	IDseparator [2]byte = [2]byte{10, 88}
// )

const (
	FeatureStorage    = "storage"
	FeatureFunctionRT = "funtion"
	FeatureServiceRT  = "service"
	FeatureRegistry   = "registry"
)

const (
	DatabasePersistent DatabaseType = "persist"
	DatabaseCache      DatabaseType = "cache"
)