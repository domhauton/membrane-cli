package commands

func PrintStatus(verbose bool, help bool) {
	if verbose {
		print("Getting Membrane Status.")
	}
	if help {
		print("Status Placeholder")
	} else {
		print("Membraned Status: FOOBAR")
	}

}
