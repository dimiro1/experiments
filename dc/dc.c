/***************************************************************************
 *   Pequena implementação da calculadora "dc" do GNU/Linux                *
 *   Não é uma implementação para calculos de alta precisão                *
 *   Copyright (C) 2009 by Claudemiro Alves Feitosa Neto                   *
 *   dimiro1@gmail.com                                                     *
 *   Modified: <2009-05-21 20:58:52 BRT>                                   *
 *                                                                         *
 *   This program is free software: you can redistribute it and/or modify  *
 *   it under the terms of the GNU General Public License as published by  *
 *   the Free Software Foundation, either version 3 of the License, or     *
 *   (at your option) any later version.                                   *
 *                                                                         *
 *   This program is distributed in the hope that it will be useful,       *
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of        *
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the         *
 *   GNU General Public License for more details.                          *
 *                                                                         *
 *   You should have received a copy of the GNU General Public License     *
 *   along with this program. If not, see <http://www.gnu.org/licences>    *
 *                                                                         *
 ***************************************************************************/

#include <stdio.h>		/* rotinas de entrada e saida */
#include <stdlib.h>		/* rotinas de alocação de memoria, ... */
#include <ctype.h>		/* rotinas usadas no analisador lexico */
#include <string.h>		/* rotinas usadas no analisador lexico */
#include <math.h>		/* rotinas para calculos mais complexos */

#define STACK_SIZE 128		/* tamanho maximo da pilha, uma melhor implementação da pilha seria usar listas ligadas */
#define TOKEN_SIZE 128		/* comprimento de um numero, poderia usar alocação dinamica */

FILE *dc_input;			/* entrada de dados, em main é atribuido o valor da entrada padrao, que é o teclado */

int dc_stack_top;		/* topo da pilha */
double dc_stack[STACK_SIZE];	/* um vetor que representa a pilha */

void dc_start_everything ();	/* inicia diversas variáveis do sistema */

/* coloca numero na pilha e devolve o topo da pilha */
int dc_stack_push (double num);

/* devolve o elemento da pilha */
double dc_stack_pop ();

/* devolve o topo da pilha */
double dc_stack_getop ();

/* imprime todos os elementos da pilha */
void dc_print_stack ();

/* imprime todos os elementos da pilha, removendo-os da pilha */
void dc_dangerous_print_stack ();

/* imprime o topo da pilha */
void dc_stack_print_top ();

/* limpa a pilha */
void dc_clean_stack ();

/* verifica se a pilha esta vazia */
int dc_stack_is_empty ();
/* verifica se a pilha esta cheia */
int dc_stack_is_full ();

/* ciclo de leitura e avaliação das expressões */
void dc_read_eval_cycle ();

/* reporta um erro e sai do programa */
void dc_fatal (char *message);
void dc_error (char *message);

/* funções para a análise léxica */
void dc_next ();
void dc_save_and_next ();
void dc_clean_tokenval ();

/* variáveis para a análise lexica */
int dc_look;			/* usado para nagevar letra por letra */
char dc_tokenval[TOKEN_SIZE];	/* usado para pegar os valores dos numeros */

/* inicio do programa */
int
main (int argc, char **argv)
{
  dc_input = stdin;		/* faz com que a entrada seja a entrada padrao, ou seja o teclado */
  dc_start_everything ();
  dc_read_eval_cycle ();	/* inicia a calculadora */
  return EXIT_SUCCESS;		/* programa terminou com sucesso */
}

void
dc_start_everything ()
{
  dc_clean_stack ();		/* inicia a pilha */
}

/* analise lexica e avaliacao de expressoes, bem simples */
void
dc_read_eval_cycle ()
{
  /* usada para para divisões, modulo ... */
  double aux;
  /* enquanto diferente de fim de arquivo,
     no windows EOF é o mesmo que CTRL-Z, no GNU/Linux CTRL-D */
  while ((dc_look = getc (dc_input)) != EOF)
    {
      dc_clean_tokenval ();	/* limpa o valor atual, do numero */

      if (dc_look == ' ' || dc_look == '\t' || dc_look == '\n')	/* elimina espaços e tabs */
        ;
      /* testa se é um numero */
      else if (isdigit (dc_look))
        {
          dc_save_and_next ();
          /* continua ate encontrar algo diferente de um numero */
          while (isdigit (dc_look))
            dc_save_and_next ();
          if (dc_look == '.')
            {
              dc_save_and_next ();
              while (isdigit (dc_look))
                dc_save_and_next ();
            }
          ungetc (dc_look, dc_input);	/* precisa devolver, a parte que não é numero */
          dc_stack_push (atof (dc_tokenval));	/* coloca o numero na pilha */
        }
      /* operações da calculadora */
      else if (dc_look == '+')
        dc_stack_push (dc_stack_pop () + dc_stack_pop ());
      else if (dc_look == '-')
        {
          aux = dc_stack_pop ();
          dc_stack_push (dc_stack_pop () + aux);
        }
      else if (dc_look == '*')
        dc_stack_push (dc_stack_pop () * dc_stack_pop ());
      else if (dc_look == '/')
        {
          aux = dc_stack_pop ();
          if (aux == 0)
            dc_fatal ("divisão por zero.");
          dc_stack_push (dc_stack_pop () / aux);
        }
      else if (dc_look == '^')	/* exponenciação */
        dc_stack_push (pow (dc_stack_pop (), dc_stack_pop ()));
      else if (dc_look == 'v')	/* raiz quadrada, cuidado ao usar essa função pois ela é uma funçao unária,
                                   ela opera sobre o topo da pilha */
        dc_stack_push (sqrt (dc_stack_pop ()));
      else if (dc_look == '%')	/* resto da divisão */
        {
          aux = dc_stack_pop ();
          if (aux == 0)
            dc_fatal ("divisão por zero.");
          dc_stack_push ((int) dc_stack_pop () % (int) aux);	/* a operação de modulo so aceita inteiros */
        }
      /* comandos */
      else if (dc_look == 'p')	/* imprime o topo da pilha */
        dc_stack_print_top ();
      else if (dc_look == 'f')	/* mostra todos os valores da pilha, sem os alterar */
        dc_print_stack ();
      else if (dc_look == 'n')	/* mostra todos os valores da pilha, eliminando-os */
        dc_dangerous_print_stack ();
      else if (dc_look == 'c')	/* limpa a pilha */
        dc_clean_stack ();
      else if (dc_look == 'd')	/* duplica o topo da pilha */
        dc_stack_push (dc_stack[dc_stack_top]);
      else if (dc_look == 'q')	/* sai normalmente do programa */
        return;
      else
        dc_error ("valor não numerico.");
    }
}

/* funções do analisador lexico */
void
dc_next ()
{
  dc_look = getc (dc_input);
}

void
dc_save_and_next ()
{
  char s[10];
  sprintf (s, "%c", dc_look);
  strcat (dc_tokenval, s);

  dc_next ();
}

void
dc_clean_tokenval ()
{
  strcpy (dc_tokenval, "");
}

/* fim funções do analisador lexico */

/* funções que operam sobre a pilha */
int
dc_stack_push (double num)
{
  int position;
  if (dc_stack_is_full ())
    dc_error ("pilha cheia.");
  position = dc_stack_top;
  dc_stack[++dc_stack_top] = num;
  return position;
}

double
dc_stack_pop ()
{
  if (dc_stack_is_empty ())
    dc_error ("pilha vazia.");
  return dc_stack[dc_stack_top--];
}

double
dc_stack_getop ()
{
  if (dc_stack_is_empty ())
    dc_error ("pilha vazia.");
  else
    return dc_stack[dc_stack_top];
}

void
dc_stack_print_top ()
{
  if (dc_stack_is_empty ())
    dc_error ("pilha vazia.");
  else
    printf ("%g\n", dc_stack_getop ());
}

void
dc_print_stack ()
{
  int i;
  if (dc_stack_is_empty ())
    dc_error ("pilha vazia.");
  else
    for (i = dc_stack_top; i >= 0; i--)
      printf ("%g\n", dc_stack[i]);
}

void
dc_dangerous_print_stack ()
{
  while (!dc_stack_is_empty ())	/* enquanto a pilha não for vazia */
    printf ("%g\n", dc_stack_pop ());
}

int
dc_stack_is_empty ()
{
  return dc_stack_top == -1;	/* caso seja -1 a pilha esta vazia */
}

int
dc_stack_is_full ()
{
  return dc_stack_top == (STACK_SIZE - 1);	/* esta cheia? */
}

void
dc_clean_stack ()
{
  dc_stack_top = -1;		/* simples assim, para limpar a pilha a unica coisa que faço é retornar dc_stack_top para -1 */
}

/* fim funções que operam sobre a pilha */

/* funções para reportar erros */
void
dc_error (char *message)
{
  fprintf (stderr, "erro: %s\n", message);	/* envia erro para a saida de erro padrão, geralmente é o monitor */
}

void
dc_fatal (char *message)
{
  fprintf (stderr, "fatal: %s\n", message);
  exit (EXIT_FAILURE);
}

/* fim funções para reportar erros */