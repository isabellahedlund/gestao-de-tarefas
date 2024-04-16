package main

type User struct {
	UserID string
	Nome   string
	Email  string
	Senha  string
}

type Task struct {
	ProjetoID       int
	Titulo          string
	Descricao       string
	Prioridade      string
	Status          string
	DataDeEntrega   int
	DataDeCriacao   int
	AtualizacaoData int
	ResponsavelID   string
	RelatorID       string
	Category        string
}

type Projeto struct {
	ProjetoID              int
	ProjetoNome            string
	DescricaoProjeto       string
	RelatorID              string
	DataDeCriacaoProjeto   int
	DataAtualizacaoProjeto int
}
