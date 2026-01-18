package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type PlugBoard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func (pb *PlugBoard) Init(alphabet string) {
	pb.alphabet = alphabet
	pb.forwardMap = make(map[string]string)
	pb.backwardMap = make(map[string]string)
}

func (pb *PlugBoard) mapping(mapAlphabet string) {
	// 元のアルファベットを保存
	originalAlphabet := pb.alphabet
	for i := 0; i < len(mapAlphabet); i++ {
		pb.forwardMap[string(originalAlphabet[i])] = string(mapAlphabet[i])
		pb.backwardMap[string(mapAlphabet[i])] = string(originalAlphabet[i])
	}
	// pb.alphabetをマッピング先のアルファベットに更新
	pb.alphabet = mapAlphabet
}

func (pb PlugBoard) forward(index int) int {
	if index < 0 || index >= len(alphabet) {
		return index
	}
	char := string(alphabet[index])
	mapped, exists := pb.forwardMap[char]
	if !exists || len(mapped) == 0 {
		return index
	}
	// mappedはpb.alphabetの文字で、それをalphabetでのインデックスに変換
	result := strings.Index(alphabet, mapped)
	if result == -1 {
		return index
	}
	return result
}

func (pb PlugBoard) backward(index int) int {
	if index < 0 || index >= len(alphabet) {
		return index
	}
	char := string(alphabet[index])
	// backwardMapで元のアルファベットの文字を取得
	mapped, exists := pb.backwardMap[char]
	if !exists || len(mapped) == 0 {
		return index
	}
	// mappedは元のアルファベットの文字なので、そのインデックスを返す
	result := strings.Index(alphabet, mapped)
	if result == -1 {
		return index
	}
	return result
}

type Rotor struct {
	plugBoard        PlugBoard
	originalAlphabet string
	offset           int
	rotations        int
}

func (r *Rotor) Init(wiringAlphabet string, offset int) {
	r.plugBoard = PlugBoard{}
	r.plugBoard.Init(wiringAlphabet)
	// 元のアルファベットを保存
	r.originalAlphabet = wiringAlphabet
	// 元のアルファベット（定数）をランダムアルファベット（wiringAlphabet）にマッピング
	for i := 0; i < len(wiringAlphabet); i++ {
		r.plugBoard.forwardMap[string(alphabet[i])] = string(wiringAlphabet[i])
		r.plugBoard.backwardMap[string(wiringAlphabet[i])] = string(alphabet[i])
	}
	r.offset = offset
	r.rotations = 0
}

func (r *Rotor) rotate(offset *int) int {
	if offset != nil {
		r.offset = *offset
	}

	// 1ステップ回転（左にシフト）
	r.plugBoard.alphabet = r.plugBoard.alphabet[1:] + r.plugBoard.alphabet[:1]
	// 回転に伴ってforwardMapとbackwardMapを更新
	r.updateMaps()
	r.rotations++
	return r.rotations
}

func (r *Rotor) updateMaps() {
	// 回転後の配線に合わせてマッピングを更新
	for i := 0; i < len(r.plugBoard.alphabet); i++ {
		r.plugBoard.forwardMap[string(alphabet[i])] = string(r.plugBoard.alphabet[i])
		r.plugBoard.backwardMap[string(r.plugBoard.alphabet[i])] = string(alphabet[i])
	}
}

func (r *Rotor) reset() {
	r.rotations = 0
	r.plugBoard.alphabet = r.originalAlphabet
	// リセット時にマッピングも更新
	r.updateMaps()
}

type Reflector struct {
	mapping map[string]string
}

func (r *Reflector) Init(mapAlphabet string) {
	r.mapping = make(map[string]string)
	for i := 0; i < len(mapAlphabet); i++ {
		r.mapping[string(alphabet[i])] = string(mapAlphabet[i])
	}

	for key, value := range r.mapping {
		if mappedValue := r.mapping[value]; mappedValue != key {
			panic("Invalid reflector mapping")
		}
	}
}

func (r Reflector) reflect(index int) int {
	reflected := r.mapping[string(alphabet[index])]
	return strings.Index(alphabet, reflected)
}

type EnigmaMachine struct {
	plugBoard PlugBoard
	rotors    []Rotor
	reflector Reflector
}

func (e *EnigmaMachine) Init(plugBoard PlugBoard, rotors []Rotor, reflector Reflector) {
	e.plugBoard = plugBoard
	e.rotors = rotors
	e.reflector = reflector
}

func (e EnigmaMachine) encrypt(plainText string) string {
	// 暗号化前にロータをリセット
	for i := range e.rotors {
		e.rotors[i].reset()
	}
	encryptedText := ""
	for _, char := range plainText {
		encryptedText += e.goThrough(char)
	}
	return encryptedText
}

func (e EnigmaMachine) decrypt(cipherText string) string {
	// 復号化前にロータをリセット
	for i := range e.rotors {
		e.rotors[i].reset()
	}
	decryptedText := ""
	for _, char := range cipherText {
		decryptedText += e.goThrough(char)
	}
	return decryptedText
}

func (e EnigmaMachine) goThrough(char rune) string {
	char = unicode.ToUpper(char)
	if !unicode.IsLetter(char) {
		return string(char)
	}

	// エニグマでは、文字を暗号化する前にロータを回転させる
	// 最右のロータが常に回転
	// 各ロータがちょうど26回転（1周）したら、その左のロータも回転させる
	i := len(e.rotors) - 1
	for i >= 0 {
		// 最右のロータは常に回転
		rotations := e.rotors[i].rotate(nil)

		// 回転後、26回転（1周）していなければ、これ以上左のロータは回転しない
		if rotations%len(alphabet) != 0 {
			break
		}
		// 1周したので、左のロータも回転させる
		i--
	}

	index := strings.Index(alphabet, string(char))
	index = e.plugBoard.forward(index)

	for _, rotor := range e.rotors {
		index = rotor.plugBoard.forward(index)
	}

	index = e.reflector.reflect(index)

	for i := len(e.rotors) - 1; i >= 0; i-- {
		index = e.rotors[i].plugBoard.backward(index)
	}

	index = e.plugBoard.backward(index)
	returnString := string(alphabet[index])

	return returnString
}

func getRandomAlphabet() string {
	alphabetRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Shuffle(len(alphabetRunes), func(i, j int) {
		alphabetRunes[i], alphabetRunes[j] = alphabetRunes[j], alphabetRunes[i]
	})
	return string(alphabetRunes)
}

func main() {
	plugBoard := PlugBoard{}
	plugBoard.Init(alphabet)
	plugBoard.mapping(getRandomAlphabet())

	r1 := Rotor{}
	r1.Init(getRandomAlphabet(), 3)
	r2 := Rotor{}
	r2.Init(getRandomAlphabet(), 2)
	r3 := Rotor{}
	r3.Init(getRandomAlphabet(), 1)

	rotors := []Rotor{r1, r2, r3}

	var alphabetArray []string
	var indexes []int
	for i, char := range alphabet {
		indexes = append(indexes, i)
		alphabetArray = append(alphabetArray, string(char))
	}

	for i := 0; i < int(len(indexes)/2); i++ {
		randomIndexX := rand.Intn(len(indexes))
		x := indexes[randomIndexX]
		indexes = append(indexes[0:randomIndexX], indexes[randomIndexX+1:]...)
		// randomIndexYは削除後のindexesサイズで計算
		if len(indexes) == 0 {
			break
		}
		randomIndexY := rand.Intn(len(indexes))
		y := indexes[randomIndexY]
		indexes = append(indexes[0:randomIndexY], indexes[randomIndexY+1:]...)
		alphabetArray[x], alphabetArray[y] = alphabetArray[y], alphabetArray[x]
	}

	reflector := Reflector{}
	reflector.Init(strings.Join(alphabetArray, ""))

	machine := EnigmaMachine{}
	machine.Init(plugBoard, rotors, reflector)

	plainText := "ATTACK SILICON VALLEY"
	encryptedText := machine.encrypt(plainText)
	decryptedText := machine.decrypt(encryptedText)

	fmt.Println("Plain Text:", plainText)
	fmt.Println("Encrypted Text:", encryptedText)
	fmt.Println("Decrypted Text:", decryptedText)

	// reflectedIndex := reflector.reflect(strings.Index(alphabet, "C"))
}
