package hedera

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate24WordMnemonic(t *testing.T) {
	mnemonic, err := GenerateMnemonic24()
	assert.NoError(t, err)

	assert.Equal(t, 24, len(mnemonic.Words()))
}

func TestGenerate12WordMnemonic(t *testing.T) {
	mnemonic, err := GenerateMnemonic24()
	assert.NoError(t, err)

	assert.Equal(t, 24, len(mnemonic.Words()))
}

func TestMnemonicFromString(t *testing.T) {
	mnemonic, err := MnemonicFromString(testMnemonic)
	assert.NoError(t, err)

	assert.Equal(t, testMnemonic, mnemonic.String())
	assert.Equal(t, 24, len(mnemonic.Words()))
}

func TestNew24MnemonicFromGeneratedMnemonic(t *testing.T) {
	generatedMnemonic, err := GenerateMnemonic24()
	assert.NoError(t, err)

	mnemonicFromSlice, err := NewMnemonic(generatedMnemonic.Words())
	assert.NoError(t, err)
	assert.Equal(t, generatedMnemonic.words, mnemonicFromSlice.words)

	mnemonicFromString, err := MnemonicFromString(generatedMnemonic.String())
	assert.NoError(t, err)
	assert.Equal(t, generatedMnemonic, mnemonicFromString)

	gKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	slKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	stKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	assert.Equal(t, gKey.keyData, slKey.keyData)
	assert.Equal(t, gKey.keyData, stKey.keyData)
}

func TestNew12MnemonicFromGeneratedMnemonic(t *testing.T) {
	generatedMnemonic, err := GenerateMnemonic12()
	assert.NoError(t, err)

	mnemonicFromSlice, err := NewMnemonic(generatedMnemonic.Words())
	assert.NoError(t, err)
	assert.Equal(t, generatedMnemonic.words, mnemonicFromSlice.words)

	mnemonicFromString, err := MnemonicFromString(generatedMnemonic.String())
	assert.NoError(t, err)
	assert.Equal(t, generatedMnemonic, mnemonicFromString)

	gKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	slKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	stKey, err := generatedMnemonic.ToPrivateKey(passphrase)
	assert.NoError(t, err)

	assert.Equal(t, gKey.keyData, slKey.keyData)
	assert.Equal(t, gKey.keyData, stKey.keyData)
}

func TestLegacyMnemonic(t *testing.T) {
	legacyString := "jolly,kidnap,tom,lawn,drunk,chick,optic,lust,mutter,mole,bride,galley,dense,member,sage,neural,widow,decide,curb,aboard,margin,manure"

	mnemonicLegacy, err := NewMnemonic(strings.Split(legacyString, ","))
	assert.NoError(t, err)

	legacyWithSpaces := strings.Join(strings.Split(legacyString, ","), " ")

	mnemonicFromString, err := MnemonicFromString(legacyWithSpaces)
	assert.NoError(t, err)
	assert.Equal(t, mnemonicLegacy, mnemonicFromString)

	gKey, err := mnemonicLegacy.ToLegacyPrivateKey()
	assert.NoError(t, err)

	slKey, err := mnemonicLegacy.ToLegacyPrivateKey()
	assert.NoError(t, err)

	stKey, err := mnemonicLegacy.ToLegacyPrivateKey()
	assert.NoError(t, err)

	assert.Equal(t, gKey.keyData, slKey.keyData)
	assert.Equal(t, gKey.keyData, stKey.keyData)
	assert.Equal(t, gKey.String(), "302e020100300506032b657004220420882a565ad8cb45643892b5366c1ee1c1ef4a730c5ce821a219ff49b6bf173ddf")
}

func TestMnemonicBreaksWithBadLength(t *testing.T) {
	// note this mnemonic is probably invalid and is only used to test breakage based on length
	shortMnemonic := "inmate flip alley wear offer often piece magnet surge toddler submit right business"

	_, err := MnemonicFromString(shortMnemonic)
	assert.Error(t, err)

	_, err = NewMnemonic(strings.Split(shortMnemonic, " "))
	assert.Error(t, err)
}
