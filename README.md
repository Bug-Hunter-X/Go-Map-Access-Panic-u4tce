# Go Map Access Panic

This repository demonstrates a common error in Go: panicking when accessing a map element using a non-existent key.  The `bug.go` file contains the erroneous code, while `bugSolution.go` provides the corrected version.

**Problem:**

The code in `bug.go` attempts to access a map element using index 0, which is incorrect. Go maps use strings as keys, not integers.  If the map has not been initialized, accessing it with any key will panic.

**Solution:**

The `bugSolution.go` file demonstrates the correct way to handle this. It checks if the key exists before accessing the element.  Alternatively, the map can be initialized first to avoid this error entirely.

This example highlights the importance of proper error handling and map initialization in Go to prevent unexpected program crashes.