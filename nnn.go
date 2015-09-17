package main

import (
	"flag"
	"fmt"
)

// These variables are initialized in compilation time by go build.
var (
	Name    string
	Version string
)

//
// var (
// 	base   = flag.String("base", "", "Base image to use when building container image")
// 	output = flag.String("output", "mycontainer", "Container name to use when exporting container rootfs")
// 	aci    = flag.Bool("aci", true, "Whether to export an AppC image as result")
// 	oci    = flag.Bool("oci", false, "Whether to export an OpenContainer image as result")
// 	docker = flag.Bool("docker", false, "Whether to export a Docker image as result")
// )
//
// type rsaKeyPair struct {
// 	privateKey []byte
// 	publicKey  []byte
// }
//
// func genRSAKeyPair() (*rsaKeyPair, error) {
// 	// Private key generation
// 	priv, err := rsa.GenerateKey(rand.Reader, 2014)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = priv.Validate()
// 	if err != nil {
// 		return nil, fmt.Errorf("Validation failed: %s", err)
// 	}
//
// 	privDer := x509.MarshalPKCS1PrivateKey(priv)
// 	privBlk := pem.Block{
// 		Type:    "RSA PRIVATE KEY",
// 		Headers: nil,
// 		Bytes:   privDer,
// 	}
// 	privPem := pem.EncodeToMemory(&privBlk)
//
// 	// Public Key generation
// 	pub := priv.PublicKey
// 	pubDer, err := x509.MarshalPKIXPublicKey(&pub)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to get der format for public key: %s", err)
// 	}
//
// 	pubBlk := pem.Block{
// 		Type:    "PUBLIC KEY",
// 		Headers: nil,
// 		Bytes:   pubDer,
// 	}
// 	pubPem := pem.EncodeToMemory(&pubBlk)
//
// 	keyPair := new(rsaKeyPair)
// 	keyPair.publicKey = pubPem
// 	keyPair.privateKey = privPem
//
// 	return keyPair, nil
// }
//
// func writeSSHKeys() {
// 	fmt.Print("Generating RSA Key pair for sshing into VM... ")
// 	keys, err := genRSAKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	keysPath := filepath.Join("."+Name, "keys")
// 	if err := os.MkdirAll(keysPath, 0640); err != nil {
// 		panic(err)
// 	}
//
// 	if err := ioutil.WriteFile(filepath.Join(keysPath, "id_rsa"), keys.privateKey, 0600); err != nil {
// 		panic(err)
// 	}
//
// 	if err := ioutil.WriteFile(filepath.Join(keysPath, "id_rsa.pub"), keys.publicKey, 0600); err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Printf("Done. Keys were placed at %s\n", keysPath)
// }

func main() {
	flag.Parse()

	// if len(os.Args) < 2 {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }

	var p XHyveParams
	p.VCPUs = 1
	p.Memory = "1024"
	p.ACPI = new(bool)
	*p.ACPI = true
	p.PCIDevs = []string{"0:0,hostbridge", "2:0,virtio-net", "31,lpc"}
	p.LPCDevs = "com1,stdio"
	p.BootParams = `kexec,imgs/alpha.801.0.0.coreos_production_pxe.vmlinuz,imgs/stable.766.3.0.coreos_production_pxe_image.cpio.gz,"earlyprintk=serial console=ttyS0 coreos.autologin cloud-config=https://raw.githubusercontent.com/coreos/coreos-xhyve/master/cloud-init/docker-only.txt"`

	if err := RunXHyve(p); err != nil {
		panic(err)
	}

	fmt.Println("-----------------> Boom!")
	//writeSSHKeys()
	//downloadHostOS()
	//bootHostOS()
	//genImage()
}
