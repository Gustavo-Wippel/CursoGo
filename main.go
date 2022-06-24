package main

import ( //"fmt"
	//"github.com/proway/exemplo/calculadora"

	"log"
	"sync"
)

// var (
// 	ErrInvalidData = errors.New("invalid data")
// )

var (
	wg = sync.WaitGroup{}
)

func main() {
	// var c calculadora.Calculadora = calculadora.NewCalculadora()

	// c.Soma(50, 2)
	// //var getNumber int = c.GetNumero()
	// c.SetNumero(10)
	// fmt.Print(c.GetNumero())

	// var u models.User = models.User{
	// 	ID:    1,
	// 	Email: "usuario@gmail.com",
	// 	Senha: "senha123",
	// }

	// c, err := json.Marshal(u)
	// if err != nil {
	// 	logrus.Error("Ocorreu um erro ao executar o JSON Marshal: %v", err)
	// 	return
	// }

	// f, erro := os.Create("user.json")
	// if erro != nil {
	// 	logrus.Error("Houve um erro ao criar o arquivo user.json: %v", err)
	// 	return
	// }

	// bytesEscritos, err := f.Write(c)
	// if erro != nil {
	// 	logrus.Error("Houve um erro ao inserir conteudo  no arquivo user.json: %v", err)
	// 	return
	// }
	// logrus.Infof("Foram escritos %d", bytesEscritos)

	// f.Close()

	// fr, err := os.Open("user.json")
	// if erro != nil {
	// 	logrus.Error("Erro ao abrir o arquivo: %v", err)
	// 	return
	// }

	// r := make([]byte, bytesEscritos)
	// bytesLidos, err := fr.Read(r)
	// if erro != nil {
	// 	logrus.Error("Houve um erro ao inserir conteudo  no arquivo user.json: %v", err)
	// 	return
	// }

	// logrus.Infof("Foram lidos %d bytes", bytesLidos)
	// logrus.Infof("%s", string(r))
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		log.Printf("%#v", err)
	// 	}
	// }()
	// panic(ErrInvalidData)

	// fmt.Print("Aaaaaaaaaaaa")
	// c := make(chan string)
	// go func() {

	// 	time.Sleep(time.Second)
	// 	c <- "teste"
	// }()

	// select {
	// case s := <-c:
	// 	log.Println(s)
	// case <-time.After(3 * time.Second):
	// 	panic("timeout")
	// }

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			log.Println("%#v\n teste", id)
			wg.Done()

		}(i)
	}
	wg.Wait()
}
