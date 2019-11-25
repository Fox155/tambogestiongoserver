package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tgs/internal/gestores"
	"tgs/internal/models"

	"github.com/gin-gonic/gin"
)

var (
	keyFile = flag.String("key", "./id_rsa", "Path to RSA private key")
)

// ProduccionesController Controlador de Producciones
type ProduccionesController struct {
	Gestor gestores.GestorProducciones
}

// Alta Permite dar de alta una Produccion
func (controlador *ProduccionesController) Alta(contexto *gin.Context) {
	produccion := models.Producciones{}
	if err := contexto.BindJSON(&produccion); err != nil {
		contexto.JSON(http.StatusBadRequest, err)
		contexto.Error(err)
		return
	}
	fmt.Println("Produccion:\t", produccion)
	if err := produccion.Validacion(); err != nil {
		contexto.JSON(http.StatusBadRequest, err.Error())
		contexto.Error(err)
		return
	}

	privada, errP := fileToPrivateKey()
	if errP != nil {
		contexto.JSON(http.StatusBadRequest, errP.Error())
		contexto.Error(errP)
		return
	}
	fmt.Println("Privada:\t", privada)

	bytes := produccion.Mensaje
	fmt.Println("Bytes:\t", bytes)

	desresultado, errD := decryptWithPrivateKey(bytes, privada)
	if errD != nil {
		contexto.JSON(http.StatusBadRequest, errD.Error())
		contexto.Error(errD)
		return
	}
	fmt.Println("Desresultado:\t", desresultado)

	fmt.Println("Desresultado Strin:\t", string(desresultado))

	contexto.JSON(http.StatusOK, produccion)
	return
}

// EstoyVivo Responde con el Mensaje "EstoyVivo"
func (controlador *ProduccionesController) EstoyVivo(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, "Estoy Vivo")
	return
}

// encryptWithPublicKey encrypts data with public key
func encryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Panicln(err)
	}
	return ciphertext
}

// fileToPrivateKey bytes to private key
func fileToPrivateKey() (*rsa.PrivateKey, error) {
	flag.Parse()
	// Read the private key
	priv, errR := ioutil.ReadFile(*keyFile)
	if errR != nil {
		return nil, errR
	}

	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// decryptWithPrivateKey decrypts data with private key
func decryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha1.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
