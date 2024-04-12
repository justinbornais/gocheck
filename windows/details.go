package windows

func GetPSPath(path string) (string, error) {
	return GetProperty(path, "PSPath")
}

func GetPSParentPath(path string) (string, error) {
	return GetProperty(path, "PSParentPath")
}

func GetPSChildName(path string) (string, error) {
	return GetProperty(path, "PSChildName")
}

func GetPSDrive(path string) (string, error) {
	return GetProperty(path, "PSDrive")
}

func GetPSProvider(path string) (string, error) {
	return GetProperty(path, "PSProvider")
}

func GetMode(path string) (string, error) {
	return GetProperty(path, "Mode")
}

func GetBaseName(path string) (string, error) {
	return GetProperty(path, "BaseName")
}

func GetTarget(path string) (string, error) {
	return GetProperty(path, "Target")
}
