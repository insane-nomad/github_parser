package files

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error // Место для хранения первой ошибки
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Пропускает запись, если раньше была ошибка
	}
	_, sw.err = fmt.Fprintln(sw.w, s) // Записывает строку и затем хранить любую ошибку
}

func SaveFile(name, data string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln(data)
	return sw.err // Возвращает ошибку в случае ее возникновения
}

func GetFileFromURL(url string) (text string) {
	var bytes []byte
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Не получилось загрузить страницу")
		fmt.Println(err)
		fmt.Println(resp)
	}
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения данных")
	}
	text = string(bytes)

	return text
}

func SaveTxt(name, data string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(data + "\n"); err != nil {
		return err
	}
	return err
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
