package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	backend "tokenManager/http"
)

func main() {

	// load environmental variables
	var environmentalVariables backend.EnvConfig
	if err := backend.LoadEnv(&environmentalVariables); err != nil {
		log.Fatalln("Error loading environmental variables: " + err.Error()) // exit
	}

	app := backend.NewApp(&environmentalVariables)
	
	port, err := findFreePort(app.Config.Build == "dev")
	if err != nil {
		log.Fatalln("Error finding port address: " + err.Error())
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: app.Routes(),
		// ErrorLog:    slog.NewLogLogger(jsonLogger, slog.LevelError),
		// IdleTimeout: time.Minute,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("server is running on http://localhost:%v\n", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panic("Server error: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	<-done
	close(done)

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Panic("Graceful server shutdown Failed: " + err.Error())
	}

	log.Println("SERVER STOPPED GRACEFULLY")
}

func findFreePort(buildMode bool) (int, error) {
	
	if buildMode {
		return 8080, nil
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// Type assert the Addr to a *net.TCPAddr to extract the port number
	addr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return 0, errors.New("could not extract the address number")
	}

	return addr.Port, nil
}

type ExtensionType int

const (
	Uninitialized ExtensionType = iota
	TransferFeeConfig
	TransferFeeAmount
	MintCloseAuthority
	ConfidentialTransferMint
	ConfidentialTransferAccount
	DefaultAccountState
	ImmutableOwner
	MemoTransfer
	NonTransferable
	InterestBearingConfig
	CpiGuard
	PermanentDelegate
	NonTransferableAccount
	TransferHook
	TransferHookAccount
	ConfidentialTransferFee       // Not implemented yet
	ConfidentialTransferFeeAmount // Not implemented yet
	MetadataPointer
	TokenMetadata
	GroupPointer
	TokenGroup
	GroupMemberPointer
	TokenGroupMember
)
