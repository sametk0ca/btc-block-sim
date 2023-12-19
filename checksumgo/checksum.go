package main

import (
    "crypto/sha256"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "bytes"
)

func main() {
    // Terminalden parametreleri al
    inputFile, input2File, nStr := os.Args[1], os.Args[2], os.Args[3]

    // n parametresini integer'a dönüştür
    n, err := strconv.Atoi(nStr)
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }

    // Input.txt dosyasını checksumla
    inputChecksum, err := checksum(inputFile, n)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Input2.txt dosyasını checksumla
    input2Checksum, err := checksum(input2File, n)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Input.txt checksum'unu string formatına dönüştür
    inputChecksumString := ""
    for _, byteValue := range inputChecksum {
        inputChecksumString += fmt.Sprintf("%08b", byteValue)
    }

    // Input2.txt checksum'unu string formatına dönüştür
    input2ChecksumString := ""
    for _, byteValue := range input2Checksum {
        input2ChecksumString += fmt.Sprintf("%08b", byteValue)
    }

    // Dosyaların checksum'unu karşılaştır
    if bytes.Equal(inputChecksum[:n], input2Checksum[:n]) {
        fmt.Println("İki dosya aynı.")
        fmt.Println("Checksum:", inputChecksumString)
    } else {
        fmt.Println("İki dosya farklı.")
        fmt.Println(inputFile, "checksum:", inputChecksumString)
        fmt.Println(input2File, "checksum:", input2ChecksumString)
    }
}

// Bir dosyanın checksum'unu hesaplar
func checksum(fileName string, n int) ([]byte, error) {
    // Dosyayı oku
    fileData, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    // SHA-256 hash fonksiyonunu kullan
    hash := sha256.New()
    hash.Write(fileData)

    // n bytelık checksum'u döndür
    return hash.Sum(nil)[:n], nil
}
