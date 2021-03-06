package initialize

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/scmn-dev/secman-v1/pkg/pc"
	"github.com/scmn-dev/secman-v1/pkg/pio"
	"golang.org/x/crypto/nacl/box"
)

const (
	saltLen     = 32
	configFound = "A secman config file was already found."
)

// Init will initialize a new password vault in the home directory.
func Init() {
	var needsDir bool
	var hasConfig bool
	var hasVault bool

	if dirExists, err := pio.PassDirExists(); err == nil {
		if !dirExists {
			needsDir = true
		} else {
			if _, err := pio.PassConfigExists(); err == nil {
				hasConfig = true
			}
			if _, err := pio.SitesVaultExists(); err == nil {
				hasVault = true
			}
		}
	}

	passDir, err := pio.GetPassDir()
	if err != nil {
		log.Fatalf("Could not get pass dir: %s", err.Error())
	}

	sitesFile, err := pio.GetSitesFile()
	if err != nil {
		log.Fatalf("Could not get sites dir: %s", err.Error())
	}

	configFile, err := pio.GetConfigPath()
	if err != nil {
		log.Fatalf("Could not get pass config: %s", err.Error())
	}

	// Prompt for the password immediately. The reason for doing this is
	// because if the user quits before the vault is fully initialized
	// (probably during password prompt since it's blocking), they will
	// be able to run init again a second time.
	pass, err := pio.PromptPass("Please enter a strong master password")
	if err != nil {
		log.Fatalf("Could not read password: %s", err.Error())
	}

	if needsDir {
		err = os.Mkdir(passDir, 0700)
		if err != nil {
			log.Fatalf("Could not create secman vault: %s", err.Error())
		} else {
			fmt.Printf("Created directory to store passwords: %s\n", passDir)
		}
	}

	if fileDirExists, err := pio.PassFileDirExists(); err == nil {
		if !fileDirExists {
			encryptedFileDir, err := pio.GetEncryptedFilesDir()
			if err != nil {
				log.Fatalf("Could not get encrypted files dir: %s", err)
			}
			err = os.Mkdir(encryptedFileDir, 0700)
			if err != nil {
				log.Fatalf("Could not create encrypted file dir: %s", err)
			}
		}
	}

	// Don't just go around deleting things for users or prompting them
	// to delete things. Make them do this manaully. Maybe this saves 1
	// person an afternoon.
	if hasConfig {
		log.Fatalf(configFound)
	}

	// Create file with secure permission.  os.Create() leaves file world-readable.
	config, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Could not create secman config: %s", err.Error())
	}

	config.Close()

	// Handle creation and initialization of the site vault.
	if !hasVault {
		// Create file, with secure permissions.
		sf, err := os.OpenFile(sitesFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("Could not create pass sites vault: %s", err.Error())
		}
		// Initialize an empty SiteFile
		siteFileContents := []byte("[]")
		_, err = sf.Write(siteFileContents)
		if err != nil {
			log.Fatalf("Could not save site file: %s", err.Error())
		}
		sf.Close()
	}

	// Generate a master password salt.
	var keySalt [32]byte
	_, err = rand.Read(keySalt[:])
	if err != nil {
		log.Fatalf("Could not generate random salt: %s", err.Error())
	}

	// Create a new salt for encrypting public key.
	var hmacSalt [32]byte
	_, err = rand.Read(hmacSalt[:])

	if err != nil {
		log.Fatalf("Could not generate random salt: %s", err.Error())
	}

	// kdf the master password.
	passKey, err := pc.Scrypt([]byte(pass), keySalt[:])

	if err != nil {
		log.Fatalf("Could not generate master key from pass: %s", err.Error())
	}

	pub, priv, err := box.GenerateKey(rand.Reader)

	if err != nil {
		log.Fatalf("Could not generate master key pair: %s", err.Error())
	}

	// Encrypt master private key with master password key.
	sealedMasterPrivKey, err := pc.Seal(&passKey, priv[:])

	if err != nil {
		log.Fatalf("Could not encrypt master key: %s", err.Error())
	}

	passConfig := pio.ConfigFile{
		MasterKeyPrivSealed: sealedMasterPrivKey,
		MasterPubKey:        *pub,
		MasterPassKeySalt:   keySalt,
	}

	if err = passConfig.SaveFile(); err != nil {
		log.Fatalf("Could not write to config file: %s", err.Error())
	}

	fmt.Println("Password Vault successfully initialized")
}
