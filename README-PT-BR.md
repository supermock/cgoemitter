# **CGOEmitter** (golang package)

<p align="center">
  <img src="cgoemitter.png" alt="Logo" width="360">
</p>
<p align="center">
<img src="https://img.shields.io/badge/coverage-62.7%25-yellow.svg" alt="Coverage">
<img src="https://img.shields.io/badge/golang-package-blue.svg" alt="Golang Package">
<img src="https://img.shields.io/badge/license-MIT-orange.svg" alt="License">
</p>

## Para que serve?
O pacote CGOEmitter serve para obter retornos assíncronos de resposta do código em **C**, facilitando o processo de devolução dos dados retornados para o programa em **GO**.

## Como utilizar?

### 1. Crie o arquivo de cabeçalho do C
É necessário incluir o arquivo **cgoemitter.h** para obter os métodos necessários.

Arquivo: **(``github.com/user/packagename/x/x.h``)**
```c
#ifndef X_H_
#define X_H_

#include <stdlib.h>
#include <stddef.h>
#include <string.h>
#include "../../cgoemitter.h"

void say_message();

#endif
```

### 2. Crie o arquivo C com sua lógica
Arquivo: **(``github.com/user/packagename/x/x.c``)**

```c
#include "x.h"

void check_err_cgoemitter_args_halloc_arg(void* value) {
  if (value == NULL) puts("Failed on cgoemitter_args_halloc_arg()");
}

void check_err_cgoemitter_args_add_arg(int code) {
  if (code == EXIT_FAILURE) puts("Failed on cgoemitter_args_add_arg()");
}

void say_message() {
  char* message = "Parameter sent from C language";

  cgoemitter_args_t cgoemitter_args = cgoemitter_new_args(1);
  
  void* message_arg = cgoemitter_args_halloc_arg(&message, (strlen(message)+1) * sizeof(char));
  check_err_cgoemitter_args_halloc_arg(message_arg);
  check_err_cgoemitter_args_add_arg(cgoemitter_args_add_arg(&cgoemitter_args, &message_arg));

  emit("message", &cgoemitter_args);
}
```

### 3. Crie o pacote que irá chamar esse método
É extremamente importante adicionar as flags abaixo, para que o seu código seja compilado.

```md
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
```

Arquivo: **(``github.com/user/packagename/x/x.go``)**

```go
package x

// The LDFLAGS lines below are needed to prevent linker errors
// since not all packages are present while building intermediate
// packages. The darwin build tag is used as a proxy for clang
// versus gcc because there doesn't seem to be a better way
// to detect this.

/*
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#include "x.h"
*/
import "C"

//SayMessage | Execute C function
func SayMessage() {
  C.say_message()
}
```

### 4. Utilizando o CGOEmitter para receber os dados retornados

Arquivo: **(``github.com/user/packagename/main.go``)**

```go
package main

import (
  "fmt"
  "github.com/supermock/cgoemitter"
  "github.com/user/packagename/x"
)

func main() {
  cgoemitter.On("message", cgoemitter.NewListener(func(args cgoemitter.Arguments) {
    fmt.Printf("Receveid message: %s\n", args.String(0))
  }))

  x.SayMessage()
}
```

### 5. Basta executar seu programa e ver a mágica acontecer

```sh
$ go run main.go
```

### Projeto com exemplo de uso do pacote: https://github.com/supermock/cgoemitter-demo

## Métodos suportados:
- **On()** => Adiciona um novo ouvinte ao evento.
- **Off()** => Remove um ouvinte existente no evento.
- **NewListener()** => Cria um novo ouvinte.
- **GetListeners()** => Retorna todos os ouvintes de um evento.
- **``parser``/CStructToGoStruct()** => Transporta os dados de uma estrutura recebida do C para uma estrutura no GO.

## Contribuições
Basta baixar o código realizar sua alteração e enviar um pull request explicando a finalidade se é um bug ou uma melhoria e etc... Após isto será analizado para ser aprovado.

## Licença 
MIT