package file

import "os"

// RedirectNull redirect files to the operating system's “null device”.
// The returned function restores this redirection.
func RedirectNull(files ...**os.File) func() {
	tmp := make([]*os.File, 0, cap(files))

	for _, v := range files {
		tmp = append(tmp, *v)
	}

	null, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}

	for i := range files {
		*files[i] = null
	}

	return func() {
		for i := range files {
			*files[i] = tmp[i]
		}
	}
}
