# Yet Another Pokedex

Mais uma aplicação CLI que utiliza a PokéAPI para mostrar dados de locais e Pokémons.

## 🛠️ Tecnologias

- **Go** - v1.26.3
- **PokéAPI** - v2

## 📁 Estrutura do Projeto

```
pokedex/
├── cmd/
│   └── pokedexcli/
│       └── main.go              
├── internal/
│   ├── repl/
│   │   ├── repl.go              
│   │   ├── commands.go         
│   │   └── repl_test.go         
│   ├── pokeapi/
│   │   ├── api.go              
│   │   ├── types.go            
│   │   ├── utils.go            
│   │   └── pokedex.go          
│   └── pokecache/
│       ├── cache.go             
│       └── cache_test.go        
├── go.mod                                         
└── README.md                  
```

## 🚀 Como Usar

### Pré-requisitos

- Go 1.26.3 ou superior
- Conexão com a internet (Duuh...) 

### Compilação

**Linux/macOS**
```bash
go build -o pokedexcli ./cmd/pokedexcli
```

**Windows**
```powershell
go build -o pokedexcli.exe ./cmd/pokedexcli
```

### Executando

**Linux/macOS**
```bash
./pokedexcli
```

**Windows**
```powershell
.\pokedexcli.exe
```

**Rodar diretamente**
```bash
go run ./cmd/pokedexcli/main.go
```

## 📖 Comandos Disponíveis

### help
Exibe a lista de comandos disponíveis.
```
Pokedex > help
```

### map
Exibe as próximas 20 áreas de localização.
```
Pokedex > map
```

### mapb
Exibe as 20 áreas de localização anteriores.
```
Pokedex > mapb
```

### explore `<location-area>`
Lista todos os Pokémon disponíveis em uma área específica.
```
Pokedex > explore canalave-city
```

### catch `<pokemon-name>`
Tenta capturar um Pokémon.
```
Pokedex > catch pikachu
```

### inspect `<pokemon-name>`
Exibe informações sobre um Pokémon capturado (nome, altura, peso, stats, tipos).
```
Pokedex > inspect pikachu
```

### exit
Encerra a execução.
```
Pokedex > exit
```

## 🔗 Recursos

- [PokéAPI Documentation](https://pokeapi.co/)
- [Go Documentation](https://golang.org/doc/)
>>>>>>> 24e82d2 (feat: adicionar estrutura inicial do projeto)
