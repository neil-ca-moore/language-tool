package strings

import "golang.org/x/text/unicode/norm"

const accented = "ÂÃÄÀÁÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõöøùúûüýþÿ"

func AccentedNFC() string {
	return norm.NFC.String(accented)
}

func AccentedNFD() string {
	return norm.NFC.String(accented)
}

func AccentedNFKC() string {
	return norm.NFKC.String(accented)
}

func AccentedNFKD() string {
	return norm.NFKD.String(accented)
}
