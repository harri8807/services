package main

import (
	"errors"
	"flag"
	"io"
	"os"
)

var (
	SEED     = flag.String("seed", "", "seed used to generate skycoin addresses")
	WAIT     = flag.Int("wait", 30, "time in seconds to wait between transactions")
	N        = flag.Int("n", 10, "number of addresses to generate from the seed")
	LOG_FILE = flag.String("log_file", "", "filepath where logs will be saved")
	LOG_TXS  = flag.Bool("log_txs", true, "show transaction logs")
	LOG_SUM  = flag.Bool("log_sum", true, "show test summary")
	CLEANUP  = flag.Bool("cleanup", false, "on test completion, send all coins to first address generated by seed")
)

func initFlags() error {
	flag.Parse()

	if len(*SEED) == 0 {
		return errors.New("seed parameter must be set to generate addresses")
	}

	if *N <= 0 {
		return errors.New("at least one address must be generated")
	}

	if *LOG_FILE != "" {
		file, err := os.OpenFile(*LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		LogWriter = io.MultiWriter(file, os.Stdout)
	} else {
		LogWriter = os.Stdout
	}

	return nil
}
