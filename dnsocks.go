package main

func main() {
	/* Initialize */
	InitLogger()
	Welcome()
	InitConfig()

	/* Start DNS */
	RunDNS()
}
